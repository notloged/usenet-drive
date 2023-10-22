package filereader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/javi11/usenet-drive/internal/usenet"
	"github.com/javi11/usenet-drive/internal/usenet/connectionpool"
	"github.com/javi11/usenet-drive/internal/usenet/corruptednzbsmanager"
	"github.com/javi11/usenet-drive/internal/usenet/nzbloader"
	"github.com/javi11/usenet-drive/pkg/osfs"
)

type file struct {
	path      string
	buffer    Buffer
	innerFile osfs.File
	fsMutex   sync.RWMutex
	log       *slog.Logger
	metadata  usenet.Metadata
	nzbLoader nzbloader.NzbLoader
	onClose   func() error
	cNzb      corruptednzbsmanager.CorruptedNzbsManager
	fs        osfs.FileSystem
}

func openFile(
	ctx context.Context,
	path string,
	flag int,
	perm os.FileMode,
	cp connectionpool.UsenetConnectionPool,
	log *slog.Logger,
	onClose func() error,
	nzbLoader nzbloader.NzbLoader,
	cNzb corruptednzbsmanager.CorruptedNzbsManager,
	fs osfs.FileSystem,
	dc downloadConfig,
) (bool, *file, error) {
	if !isNzbFile(path) {
		originalFile := getOriginalNzb(fs, path)
		if originalFile != nil {
			// If the file is a masked call the original nzb file
			path = filepath.Join(filepath.Dir(path), originalFile.Name())
		} else {
			return false, nil, nil
		}
	}

	f, err := fs.OpenFile(path, flag, perm)
	if err != nil {
		return true, nil, err
	}

	n, err := nzbLoader.LoadFromFileReader(f)
	if err != nil {
		log.ErrorContext(ctx, fmt.Sprintf("Error getting loading nzb %s", path), "err", err)
		return true, nil, os.ErrNotExist
	}

	buffer, err := NewBuffer(
		&n.Nzb.Files[0],
		int(n.Metadata.FileSize),
		int(n.Metadata.ChunkSize),
		dc,
		cp,
		log,
	)
	if err != nil {
		return true, nil, err
	}

	return true, &file{
		innerFile: f,
		buffer:    buffer,
		metadata:  n.Metadata,
		path:      usenet.ReplaceFileExtension(path, n.Metadata.FileExtension),
		log:       log,
		nzbLoader: nzbLoader,
		onClose:   onClose,
		cNzb:      cNzb,
		fs:        fs,
	}, nil
}

func (f *file) Chdir() error {
	return f.innerFile.Chdir()
}

func (f *file) Chmod(mode os.FileMode) error {
	return f.innerFile.Chmod(mode)
}

func (f *file) Chown(uid, gid int) error {
	return f.innerFile.Chown(uid, gid)
}

func (f *file) Close() error {
	if err := f.buffer.Close(); err != nil {
		return err
	}

	if f.onClose != nil {
		if err := f.onClose(); err != nil {
			return err
		}
	}

	return f.innerFile.Close()
}

func (f *file) Fd() uintptr {
	return f.innerFile.Fd()
}

func (f *file) Name() string {
	return f.path
}

func (f *file) Read(b []byte) (int, error) {
	f.fsMutex.RLock()
	defer f.fsMutex.RUnlock()

	n, err := f.buffer.Read(b)
	if err != nil {
		if errors.Is(err, ErrCorruptedNzb) {
			f.log.Error("Marking file as corrupted:", "error", err, "fileName", f.path)
			err := f.cNzb.Add(context.Background(), f.path, err.Error())
			if err != nil {
				f.log.Error("Error adding corrupted nzb to the database:", "error", err)
			}

			return n, io.ErrUnexpectedEOF
		}

		return n, err
	}

	return n, nil
}

func (f *file) ReadAt(b []byte, off int64) (int, error) {
	f.fsMutex.RLock()
	defer f.fsMutex.RUnlock()

	n, err := f.buffer.ReadAt(b, off)
	if err != nil {
		if errors.Is(err, ErrCorruptedNzb) {
			f.log.Error("Marking file as corrupted:", "error", err, "fileName", f.path)
			err := f.cNzb.Add(context.Background(), f.path, err.Error())
			if err != nil {
				f.log.Error("Error adding corrupted nzb to the database:", "error", err)
			}

			return n, io.ErrUnexpectedEOF
		}

		return n, err
	}

	return n, nil
}

func (f *file) Readdir(n int) ([]os.FileInfo, error) {
	// remote files will never be a dir
	return []os.FileInfo{}, os.ErrPermission
}

func (f *file) Readdirnames(n int) ([]string, error) {
	return f.innerFile.Readdirnames(n)
}

func (f *file) Seek(offset int64, whence int) (n int64, err error) {
	f.fsMutex.RLock()
	n, err = f.buffer.Seek(offset, whence)
	f.fsMutex.RUnlock()
	return
}

func (f *file) SetDeadline(t time.Time) error {
	return f.innerFile.SetDeadline(t)
}

func (f *file) SetReadDeadline(t time.Time) error {
	return f.innerFile.SetReadDeadline(t)
}

func (f *file) SetWriteDeadline(t time.Time) error {
	return os.ErrPermission
}

func (f *file) Stat() (os.FileInfo, error) {
	f.fsMutex.RLock()
	defer f.fsMutex.RUnlock()

	return NeFileInfoWithMetadata(
		f.innerFile.Name(),
		f.metadata,
		f.fs,
	)
}

func (f *file) Sync() error {
	return f.innerFile.Sync()
}

func (f *file) Truncate(size int64) error {
	return os.ErrPermission
}

func (f *file) Write(b []byte) (int, error) {
	return 0, os.ErrPermission
}

func (f *file) WriteAt(b []byte, off int64) (int, error) {
	return 0, os.ErrPermission
}

func (f *file) WriteString(s string) (int, error) {
	return 0, os.ErrPermission
}