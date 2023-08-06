package webdav

import (
	"os"
	"path"
	"sync"
	"time"

	"github.com/javi11/usenet-drive/internal/domain"
)

type NzbFile struct {
	name string
	*os.File
}

func NewNzbFile(name string, flag int, perm os.FileMode) (*NzbFile, error) {
	var metadata *domain.NZB
	var err error
	var file *os.File

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		metadata, err = domain.LoadNZBFileMetadata(name)
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

	originalName := metadata.Head.GetMetaByType(domain.FileName)
	extension := path.Ext(originalName)

	return &NzbFile{
		File: file,
		name: replaceFileExtension(name, extension),
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
	return f.File.Read(b)
}

func (f *NzbFile) ReadAt(b []byte, off int64) (int, error) {
	return f.File.ReadAt(b, off)
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
	return f.File.Seek(offset, whence)
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
