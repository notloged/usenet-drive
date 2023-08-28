package utils

import (
	"path/filepath"
	"strings"
)

func ReplaceFileExtension(name string, extension string) string {
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext) + extension
}
