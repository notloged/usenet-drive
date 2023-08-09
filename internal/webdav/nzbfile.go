package webdav

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"

	"github.com/chrisfarms/nzb"
	"github.com/chrisfarms/yenc"
	"github.com/javi11/usenet-drive/internal/domain"
	"github.com/javi11/usenet-drive/internal/usenet"
)

type NzbFile struct {
	name string
	size int64
	cp   UsenetConnectionPool
	*os.File
	nzbFile *nzb.Nzb
	pos     int
	mu      sync.Mutex
}

func NewNzbFile(name string, flag int, perm os.FileMode, cp UsenetConnectionPool) (*NzbFile, error) {
	var metadata domain.Metadata
	var err error
	var file *os.File
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		metadata, err = domain.LoadMetadata(name)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		file, err = os.OpenFile(name, flag, perm)
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}
	nzbFile, err := nzb.New(file)
	if err != nil {
		return nil, err
	}

	return &NzbFile{
		File:    file,
		size:    metadata.FileSize,
		name:    replaceFileExtension(name, metadata.FileExtension),
		cp:      cp,
		nzbFile: nzbFile,
	}, nil
}

func (f *NzbFile) Chdir() error {
	return f.File.Chdir()
}

func (f *NzbFile) Chmod(mode os.FileMode) error {
	return f.File.Chmod(mode)
}

func (f *NzbFile) Chown(uid, gid int) error {
	return f.File.Chown(uid, gid)
}

func (f *NzbFile) Close() error {
	return f.File.Close()
}

func (f *NzbFile) Fd() uintptr {
	return f.File.Fd()
}

func (f *NzbFile) Name() string {
	return f.name
}

func (f *NzbFile) Read(b []byte) (int, error) {
	return f.readAt(b, 0)
}

func (f *NzbFile) ReadAt(b []byte, off int64) (int, error) {
	return f.readAt(b, off)
}

func (f *NzbFile) Readdir(n int) ([]os.FileInfo, error) {
	infos, err := f.File.Readdir(n)
	if err != nil {
		return nil, err
	}

	for i, info := range infos {
		if isNzbFile(info.Name()) {
			infos[i], _ = NewFileInfoWithMetadata(info.Name())
		}
	}

	return infos, nil
}

func (f *NzbFile) Readdirnames(n int) ([]string, error) {
	return f.File.Readdirnames(n)
}

func (f *NzbFile) Seek(offset int64, whence int) (int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	npos := f.pos
	switch whence {
	case io.SeekStart:
		npos = int(offset)
	case io.SeekCurrent:
		npos += int(offset)
	case io.SeekEnd:
		npos = int(f.size) + int(offset)
	default:
		npos = -1
	}
	if npos < 0 {
		return 0, os.ErrInvalid
	}
	f.pos = npos
	return int64(f.pos), nil
}

func (f *NzbFile) SetDeadline(t time.Time) error {
	return f.File.SetDeadline(t)
}

func (f *NzbFile) SetReadDeadline(t time.Time) error {
	return f.File.SetReadDeadline(t)
}

func (f *NzbFile) SetWriteDeadline(t time.Time) error {
	return f.File.SetWriteDeadline(t)
}

func (f *NzbFile) Stat() (os.FileInfo, error) {
	if isNzbFile(f.File.Name()) {
		return NewFileInfoWithMetadata(f.File.Name())
	}

	return f.File.Stat()
}

func (f *NzbFile) Sync() error {
	return f.File.Sync()
}

func (f *NzbFile) Truncate(size int64) error {
	return f.File.Truncate(size)
}

func (f *NzbFile) Write(b []byte) (int, error) {
	return f.File.Write(b)
}

func (f *NzbFile) WriteAt(b []byte, off int64) (int, error) {
	return f.File.WriteAt(b, off)
}

func (f *NzbFile) WriteString(s string) (int, error) {
	return f.File.WriteString(s)
}

func (f *NzbFile) readAt(b []byte, off int64) (int, error) {
	file := f.nzbFile.Files[0]
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.pos >= int(f.size) {
		return 0, io.EOF
	}

	n := 0
	cp := math.Round(float64(f.pos) / float64(file.Segments[0].Bytes))
	for i, segment := range file.Segments[int(cp):] {
		if n >= len(b) {
			break
		}
		// Get the connection from the pool
		conn, err := f.cp.GetConnection()
		if err != nil {
			f.cp.CloseConnection(conn)
			fmt.Fprintln(os.Stderr, "nntp error:", err)
			break
		}
		err = usenet.FindGroup(conn, file.Groups)
		if err != nil {
			f.cp.CloseConnection(conn)
			fmt.Fprintln(os.Stderr, "nntp error:", err)
			break
		}

		body, err := conn.Body(fmt.Sprintf("<%v>", segment.Id))
		if err != nil {
			f.cp.CloseConnection(conn)
			fmt.Fprintln(os.Stderr, "nntp error:", err)
			break
		}

		yread, err := yenc.Decode(body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		beginReadAt := ((i * segment.Bytes) - f.pos) - 1
		if beginReadAt < 0 {
			beginReadAt = 0
		}
		beginWriteAt := n - 1
		if beginWriteAt < 0 {
			beginWriteAt = 0
		}
		bc := copy(b[beginWriteAt:], yread.Body[beginReadAt:])
		n += bc
	}
	f.pos += n

	return n, nil
}
