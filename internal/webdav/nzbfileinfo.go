package webdav

import (
	"io/fs"
	"os"
	"sync"
	"time"

	"github.com/javi11/usenet-drive/internal/domain"
	"github.com/javi11/usenet-drive/internal/utils"
)

type nzbFileInfoWithMetadata struct {
	nzbFile  os.FileInfo
	name     string
	metadata domain.Metadata
}

func NewFileInfoWithMetadata(name string) (fs.FileInfo, error) {
	var nzbFileInfo os.FileInfo
	var metadata domain.Metadata
	var err error

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		file, err := os.OpenFile(name, os.O_RDONLY, 0)
		if err != nil {
			return
		}

		nzbFile, err := parseNzbFile(file)
		if err != nil {
			return
		}

		metadata, err = domain.LoadFromNzb(nzbFile)
		if err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nzbFileInfo, err = os.Stat(name)
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	fileName := nzbFileInfo.Name()

	return &nzbFileInfoWithMetadata{
		nzbFile:  nzbFileInfo,
		metadata: metadata,
		name:     utils.ReplaceFileExtension(fileName, metadata.FileExtension),
	}, nil
}

func (fi *nzbFileInfoWithMetadata) Size() int64 {
	// We need the original file size to display it.
	return fi.metadata.FileSize
}

func (fi *nzbFileInfoWithMetadata) ModTime() time.Time {
	// We need the original file mod time in order to allow comparing when replace a file. Files will never be modified.
	return fi.metadata.ModTime
}

func (fi *nzbFileInfoWithMetadata) IsDir() bool {
	return false
}

func (fi *nzbFileInfoWithMetadata) Sys() any {
	return nil
}

func (fi *nzbFileInfoWithMetadata) Name() string {
	return fi.name
}

func (fi *nzbFileInfoWithMetadata) Mode() fs.FileMode {
	return fi.nzbFile.Mode()
}
