package webdav

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/chrisfarms/nntp"
	"github.com/chrisfarms/nzb"
)

var extStrip = regexp.MustCompile(`\.nzb$`)
var existErr = errors.New("file exists")

type stream struct {
	cp UsenetConnectionPool
}

func NewStream(cp *UsenetConnectionPool) *stream {
	return &stream{
		cp: cp,
	}
}

func (s *stream) GetNzbStream(nzbFile *nzb.Nzb) (error, string) {
	// For the moment we only support single file nzb's

	file := nzbFile.Files[0]

	return s.getStream(file)
}

// download a single file contained in an nzb.
func (s *stream) getStream(nzbFile *nzb.NzbFile) error {
	for _, f := range nzbFile.Segments {
		c, err := s.cp.GetConnection()
		if err != nil {
			return err
		}
		go decodeMsg(c, file, nzbfile.Groups, f.MsgId)
	}
	return nil
}

// decodes an nntp message and writes it to a section of the file.
func decodeMsg(c *nntp.Conn, f *file, groups []string, MsgId string) {
	var err error
	defer f.Done()
	err = findGroup(c, groups)
	if err != nil {
		putBroken(c)
		fmt.Fprintln(os.Stderr, "nntp error:", err)
		return
	}
	rc, err := c.GetMessage(MsgId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "nntp error getting", MsgId, ":", err)
		if _, ok := err.(*textproto.Error); ok {
			putConn(c)
		} else {
			putBroken(c)
		}
		return
	}
	putConn(c)

	yread, err := yenc.NewPart(bytes.NewReader(rc))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	wr := f.WriterAt(yread.Begin)
	_, err = io.Copy(wr, yread)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func findGroup(c *nntp.Conn, groups []string) error {
	var err error
	for _, g := range groups {
		err = c.SwitchGroup(g)
		if err == nil {
			return nil
		}
	}
	return err
}

// waitgroup that keeps track if there are any files being downloaded.
var filewg sync.WaitGroup

type file struct {
	name      string
	path      string
	file      *os.File
	partsLeft int
	mu        sync.Mutex
}

func newFile(dirname string, nzbfile *nzb.File) (*file, error) {
	filename := nzbfile.Subject.Filename()
	if filename == "" {
		return nil, errors.New("bad subject")
	}

	path := filepath.Join(dirname, filename)
	if _, err := os.Stat(path); err == nil {
		return nil, existErr
	}

	temppath := path + ".gonztemp"
	f, err := os.Create(temppath)
	if err != nil {
		return nil, err
	}

	ret := &file{
		name:      filename,
		path:      path,
		partsLeft: len(nzbfile.Segments),
		file:      f,
	}
	filewg.Add(1)
	return ret, nil
}

func (f *file) WriterAt(offset int64) io.Writer {
	return &fileWriter{
		f:      f.file,
		offset: offset,
	}
}

func (f *file) Done() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.partsLeft--
	if f.partsLeft != 0 {
		return
	}
	fmt.Printf("Done downloading file %q\n", f.name)
	os.Rename(f.file.Name(), f.path)
	f.file.Close()
	filewg.Done()

}

// filewriter allows for multiple goroutines to write concurrently to
// non-overlapping sections of a file
type fileWriter struct {
	f      *os.File
	offset int64
}

func (f *fileWriter) Write(b []byte) (int, error) {
	n, err := f.f.WriteAt(b, f.offset)
	f.offset += int64(n)
	return n, err
}
