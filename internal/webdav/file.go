package webdav

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
)

type customFile struct {
	*os.File
	folderName string
	mutex      *sync.RWMutex
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
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.File.Read(b)
}

func (f *customFile) ReadAt(b []byte, off int64) (int, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.File.ReadAt(b, off)
}

func (f *customFile) Readdir(n int) ([]os.FileInfo, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	infos, err := f.File.Readdir(n)
	if err != nil {
		return nil, err
	}

	var merr multierror.Group

	for i, info := range infos {
		if isNzbFile(info.Name()) {
			info := info
			i := i
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

	finalInfo := make([]os.FileInfo, 0)
	for i, info := range infos {
		if !isMetadataFile(info.Name()) {
			finalInfo = append(finalInfo, infos[i])
		}
	}

	return finalInfo, nil
}

func (f *customFile) Readdirnames(n int) ([]string, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.File.Readdirnames(n)
}

func (f *customFile) Seek(offset int64, whence int) (int64, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
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
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.File.Stat()
}

func (f *customFile) Sync() error {
	return f.File.Sync()
}

func (f *customFile) Truncate(size int64) error {
	if isNzbFile(f.Name()) {
		return os.ErrPermission
	}
	return f.File.Truncate(size)
}

func (f *customFile) Write(b []byte) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isNzbFile(f.Name()) {
		return 0, os.ErrPermission
	}
	return f.File.Write(b)
}

func (f *customFile) WriteAt(b []byte, off int64) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isNzbFile(f.Name()) {
		return 0, os.ErrPermission
	}
	return f.File.WriteAt(b, off)
}

func (f *customFile) WriteString(s string) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	if isNzbFile(f.Name()) {
		return 0, os.ErrPermission
	}
	return f.File.WriteString(s)
}
