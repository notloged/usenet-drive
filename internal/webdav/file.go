package webdav

import (
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-multierror"
)

type customFile struct {
	*os.File
	folderName string
}

func (f *customFile) Chdir() error {
	return f.File.Chdir()
}

func (f *customFile) Chmod(mode os.FileMode) error {
	return f.File.Chmod(mode)
}

func (f *customFile) Chown(uid, gid int) error {
	return f.File.Chown(uid, gid)
}

func (f *customFile) Close() error {
	return f.File.Close()
}

func (f *customFile) Fd() uintptr {
	return f.File.Fd()
}

func (f *customFile) Name() string {
	return f.File.Name()
}

func (f *customFile) Read(b []byte) (int, error) {
	return f.File.Read(b)
}

func (f *customFile) ReadAt(b []byte, off int64) (int, error) {
	return f.File.ReadAt(b, off)
}

func (f *customFile) Readdir(n int) ([]os.FileInfo, error) {
	infos, err := f.File.Readdir(n)
	if err != nil {
		return nil, err
	}

	var merr multierror.Group

	for i, info := range infos {
		if isNzbFile(info.Name()) {
			info := info
			merr.Go(func() error {
				infos[i], err = NewFileInfoWithMetadata(filepath.Join(f.folderName, info.Name()))
				if err != nil {
					return err
				}

				return nil
			})
		}
	}

	if err := merr.Wait(); err != nil {
		return nil, err
	}

	return infos, nil
}

func (f *customFile) Readdirnames(n int) ([]string, error) {
	return f.File.Readdirnames(n)
}

func (f *customFile) Seek(offset int64, whence int) (int64, error) {
	return f.File.Seek(offset, whence)
}

func (f *customFile) SetDeadline(t time.Time) error {
	return f.File.SetDeadline(t)
}

func (f *customFile) SetReadDeadline(t time.Time) error {
	return f.File.SetReadDeadline(t)
}

func (f *customFile) SetWriteDeadline(t time.Time) error {
	return f.File.SetWriteDeadline(t)
}

func (f *customFile) Stat() (os.FileInfo, error) {
	return f.File.Stat()
}

func (f *customFile) Sync() error {
	return f.File.Sync()
}

func (f *customFile) Truncate(size int64) error {
	return f.File.Truncate(size)
}

func (f *customFile) Write(b []byte) (int, error) {
	return f.File.Write(b)
}

func (f *customFile) WriteAt(b []byte, off int64) (int, error) {
	return f.File.WriteAt(b, off)
}

func (f *customFile) WriteString(s string) (int, error) {
	return f.File.WriteString(s)
}
