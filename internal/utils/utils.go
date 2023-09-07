package utils

import (
	"path/filepath"
	"strings"
)

func ReplaceFileExtension(name string, extension string) string {
	ext := filepath.Ext(name)
	return strings.TrimSuffix(name, ext) + extension
}

func TruncateFileName(name string, extension string, length int) string {
	if len(name) <= length {
		return name
	}

	name = strings.TrimSuffix(name, extension)

	if len(name) <= length {
		return name + extension
	}

	return name[:length] + extension
}

func HasAllowedExtension(path string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
