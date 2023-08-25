package webdav

import (
	"os"

	"github.com/chrisfarms/nzb"
	"github.com/hraban/lrucache"
)

var nzbParserCache = lrucache.New(100)

func parseNzbFile(file *os.File) (nzbFile *nzb.Nzb, err error) {
	// Cache nzb file parser since it should never change and is expensive to parse
	name := file.Name()
	hit, _ := nzbParserCache.Get(name)
	if f, ok := hit.(*nzb.Nzb); ok {
		nzbFile = f
	} else {
		nzbFile, err = nzb.New(file)
		if err != nil {
			return nil, err
		}
		nzbParserCache.Set(name, nzbFile)
	}

	return
}
