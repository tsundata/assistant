package fs

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/app/storage/fs/adapter"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"os"
)

type Filesystem interface {
	// Exists Determine if a file or directory exists.
	Exists(path string) (bool, error)
	// Get the contents of a file.
	Get(path string, lock bool) ([]byte, error)
	// Hash Get the MD5 hash of the file at the given path.
	Hash(path string) (string, error)
	// Put Write the contents of a file
	Put(path string, contents []byte, lock bool) error
	// Replace write the contents of a file, replacing it atomically if it already exists.
	Replace(path string, content []byte) error
	// Append to a file
	Append(path string, data []byte) error
	// Delete the file at a given path.
	Delete(path string) error
	// Move a file to a new location.
	Move(path, target string) error
	// Copy a file to a new location.
	Copy(path, target string) error
	// Size Get the file size of a given file.
	Size(path string) (int, error)
	// IsDir Determine if the given path is a directory.
	IsDir(path string) (bool, error)
	// IsFile Determine if the given path is a file.
	IsFile(path string) (bool, error)
	// Glob Find path names matching a given pattern.
	Glob(pattern string, flags int) ([]os.File, error)
	// Files Get an array of all files in a directory.
	Files(dir string, hidden bool) ([]os.File, error)
	// Dirs Get all of the directories within a given directory.
	Dirs(dir string) ([]os.File, error)
	// MakeDir Create a directory.
	MakeDir(path string) error
	// MoveDir Move a directory
	MoveDir(from, to string, overwrite bool) error
	// CopyDir Copy a directory from one location to another
	CopyDir(dir, target string) error
	// DeleteDir Recursively delete a directory
	DeleteDir(dir string) error
	// CleanDir Empty the specified directory of all files and folders
	CleanDir(dir string) error
	// FullPath Get full path
	FullPath(path string) string
	// AbsolutePath Get absolute path
	AbsolutePath(path string) string
}

func FS(config *config.AppConfig) (Filesystem, error) {
	switch config.Storage.Adapter {
	case dropbox.ID:
		return &adapter.Dropbox{}, nil
	case "local":
		return &adapter.Local{
			Dir:    config.Storage.Dir,
			Domain: config.Gateway.Url,
		}, nil
	default:
		return nil, errors.New("not filesystem adapter")
	}
}
