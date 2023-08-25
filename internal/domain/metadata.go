package domain

import (
	"strconv"

	"github.com/chrisfarms/nzb"
)

type Metadata struct {
	FileName      string `json:"file_name"`
	FileExtension string `json:"file_extension"`
	FileSize      int64  `json:"file_size"`
}

func LoadFromNzb(nzbFile *nzb.Nzb) (Metadata, error) {
	fileSize, err := strconv.ParseInt(nzbFile.Meta["file_size"], 10, 64)
	if err != nil {
		return Metadata{}, err
	}

	return Metadata{
		FileName:      nzbFile.Meta["file_name"],
		FileExtension: nzbFile.Meta["file_extension"],
		FileSize:      fileSize,
	}, nil
}
