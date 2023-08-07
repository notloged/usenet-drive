package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Metadata struct {
	FileName      string `json:"file_name"`
	FileExtension string `json:"file_extension"`
	FileSize      int64  `json:"file_size"`
	FileBlocks    int64  `json:"file_blocks"`
	FileIOBlock   int64  `json:"file_io_block"`
}

func LoadMetadata(nzbFilePath string) (Metadata, error) {
	var m Metadata

	data, err := os.ReadFile(fmt.Sprintf("%s.metadata.json", strings.TrimSuffix(nzbFilePath, ".nzb")))
	if err != nil {
		return Metadata{}, err
	}

	err = json.Unmarshal(data, &m)
	if err != nil {
		return Metadata{}, err
	}

	return m, nil
}
