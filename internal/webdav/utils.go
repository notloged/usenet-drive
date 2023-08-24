package webdav

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/exp/constraints"
)

func slashClean(name string) string {
	if name == "" || name[0] != '/' {
		name = "/" + name
	}
	return path.Clean(name)
}

func isNzbFile(name string) bool {
	return strings.HasSuffix(name, ".nzb")
}

func replaceFileExtension(name string, extension string) string {
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext) + extension
}

func getOriginalNzb(name string) *string {
	originalName := replaceFileExtension(name, ".nzb")
	_, err := os.Stat(originalName)
	if os.IsNotExist(err) {
		return nil
	}

	return &originalName
}

func isMetadataFile(name string) bool {
	return strings.HasSuffix(name, ".metadata.json")
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
