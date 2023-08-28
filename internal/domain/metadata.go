package domain

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/chrisfarms/nzb"
)

type Metadata struct {
	FileName      string `json:"file_name"`
	FileExtension string `json:"file_extension"`
	FileSize      int64  `json:"file_size"`
	ChunkSize     int64  `json:"chunk_size"`
}

func LoadFromNzb(nzbFile *nzb.Nzb) (Metadata, error) {
	fileSize, err := strconv.ParseInt(nzbFile.Meta["file_size"], 10, 64)
	if err != nil {
		return Metadata{}, err
	}

	// Chunk size is present in the file subject string
	// segment size is not the real size of a segment. Segment size = chunk size + yenc overhead
	chunkSize, err := getChunkSizeFromSubject(nzbFile.Files[0].Subject)
	if err != nil {
		return Metadata{}, err
	}

	return Metadata{
		FileName:      nzbFile.Meta["file_name"],
		FileExtension: nzbFile.Meta["file_extension"],
		FileSize:      fileSize,
		ChunkSize:     chunkSize,
	}, nil
}

func getChunkSizeFromSubject(s string) (int64, error) {
	re := regexp.MustCompile(`size=(\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return 0, fmt.Errorf("no size found in string")
	}
	size, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}
