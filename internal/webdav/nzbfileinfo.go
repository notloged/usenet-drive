package webdav

import (
	"io/fs"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/javi11/usenet-drive/internal/domain"
)

type nzbFileInfoWithMetadata struct {
	nzbFile os.FileInfo
	size    int64
	name    string
}

func NewFileInfoWithMetadata(name string) (fs.FileInfo, error) {
	var nzbFileInfo os.FileInfo
	var metadata *domain.NZB
	var err error

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		metadata, err = domain.LoadNZBFileMetadata(name)
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

	sizeStr := metadata.Head.GetMetaByType(domain.FileSize)
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return nil, err
	}

	originalName := metadata.Head.GetMetaByType(domain.FileName)
	extension := path.Ext(originalName)

	return &nzbFileInfoWithMetadata{
		nzbFile: nzbFileInfo,
		size:    size,
		name:    replaceFileExtension(fileName, extension),
	}, nil
}

func (fi *nzbFileInfoWithMetadata) Size() int64 {
	return fi.size
}

func (fi *nzbFileInfoWithMetadata) ModTime() time.Time {
	return fi.nzbFile.ModTime()
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
