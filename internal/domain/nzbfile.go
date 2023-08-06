package domain

import (
	"encoding/xml"
	"os"
)

type MetaTypes string

const (
	FileSize MetaTypes = "file_size"
	FileName MetaTypes = "file_name"
	Password MetaTypes = "password"
	Blocks   MetaTypes = "blocks"
	IOBlock  MetaTypes = "io_block"
	NZBName  MetaTypes = "nzb_name"
)

type NZB struct {
	XMLName xml.Name `xml:"nzb"`
	Files   []File   `xml:"file"`
	Head    Head     `xml:"head"`
}

type Head struct {
	Meta []Meta `xml:"meta"`
}

func (n *Head) GetMetaByType(t MetaTypes) string {
	for _, meta := range n.Meta {
		if meta.Type == t {
			return meta.Value
		}
	}

	return ""
}

type Meta struct {
	Type  MetaTypes `xml:"type,attr"`
	Value string    `xml:",chardata"`
}

type File struct {
	Poster   string    `xml:"poster,attr"`
	Date     int64     `xml:"date,attr"`
	Subject  string    `xml:"subject,attr"`
	Groups   []string  `xml:"groups>group"`
	Segments []Segment `xml:"segments>segment"`
}

type Segment struct {
	Bytes  int    `xml:"bytes,attr"`
	Number int    `xml:"number,attr"`
	Value  string `xml:",chardata"`
}

func LoadNZBFileMetadata(filename string) (*NZB, error) {
	// Read the file contents
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the XML into a NZB struct
	var nzb NZB
	err = xml.Unmarshal(contents, &nzb)
	if err != nil {
		return nil, err
	}

	return &nzb, nil
}
