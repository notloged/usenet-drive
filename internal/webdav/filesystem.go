package webdav

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/webdav"
)

type nzbFilesystem string

func (fs nzbFilesystem) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	if name = fs.resolve(name); name == "" {
		return os.ErrNotExist
	}
	return os.Mkdir(name, perm)
}

func (fs nzbFilesystem) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	if name = fs.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}

	if isNzbFile(name) {
		// If the file doesn't have the .nzb extension, use a normal os.Stat() call
		return NewNzbFile(name, flag, perm)
	}

	originalName := getOriginalNzb(name)
	if originalName != nil {
		// If the file is a masked call the orginal nzb file
		return NewNzbFile(*originalName, flag, perm)
	}

	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return &customFile{File: f, folderName: name}, nil
}

func (fs nzbFilesystem) RemoveAll(ctx context.Context, name string) error {
	if name = fs.resolve(name); name == "" {
		return os.ErrNotExist
	}
	if name == filepath.Clean(string(fs)) {
		// Prohibit removing the virtual root directory.
		return os.ErrInvalid
	}
	return os.RemoveAll(name)
}

func (fs nzbFilesystem) Rename(ctx context.Context, oldName, newName string) error {
	if oldName = fs.resolve(oldName); oldName == "" {
		return os.ErrNotExist
	}
	if newName = fs.resolve(newName); newName == "" {
		return os.ErrNotExist
	}
	if root := filepath.Clean(string(fs)); root == oldName || root == newName {
		// Prohibit renaming from or to the virtual root directory.
		return os.ErrInvalid
	}
	return os.Rename(oldName, newName)
}

func (fs nzbFilesystem) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	if name = fs.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}

	if isNzbFile(name) {
		// If the file doesn't have the .nzb extension, use a normal os.Stat() call
		return NewFileInfoWithMetadata(name)
	}

	originalName := getOriginalNzb(name)
	if originalName != nil {
		// If the file is a masked call the orginal nzb file
		return NewFileInfoWithMetadata(*originalName)
	}

	// Build a new os.FileInfo with a mix of nzbFileInfo and metadata
	return os.Stat(name)
}

func (fs nzbFilesystem) resolve(name string) string {
	// This implementation is based on Dir.Open's code in the standard net/http package.
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return ""
	}
	dir := string(fs)
	if dir == "" {
		dir = "."
	}
	return filepath.Join(dir, filepath.FromSlash(slashClean(name)))
}
