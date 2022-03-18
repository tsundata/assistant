package adapter

import (
	"os"
)

type GoogleDrive struct{}

func (g *GoogleDrive) AbsolutePath(path string) string {
	panic("implement me")
}

func (g *GoogleDrive) FullPath(path string) string {
	panic("implement me")
}

func (g *GoogleDrive) Exists(path string) (bool, error) {
	panic("implement me")
}

func (g *GoogleDrive) Get(path string, lock bool) ([]byte, error) {
	panic("implement me")
}

func (g *GoogleDrive) Hash(path string) (string, error) {
	panic("implement me")
}

func (g *GoogleDrive) Put(path string, contents []byte, lock bool) error {
	panic("implement me")
}

func (g *GoogleDrive) Replace(path string, content []byte) error {
	panic("implement me")
}

func (g *GoogleDrive) Append(path string, data []byte) error {
	panic("implement me")
}

func (g *GoogleDrive) Delete(path string) error {
	panic("implement me")
}

func (g *GoogleDrive) Move(path, target string) error {
	panic("implement me")
}

func (g *GoogleDrive) Copy(path, target string) error {
	panic("implement me")
}

func (g *GoogleDrive) Size(path string) (int, error) {
	panic("implement me")
}

func (g *GoogleDrive) IsDir(path string) (bool, error) {
	panic("implement me")
}

func (g *GoogleDrive) IsFile(path string) (bool, error) {
	panic("implement me")
}

func (g *GoogleDrive) Glob(pattern string, flags int) ([]os.File, error) {
	panic("implement me")
}

func (g *GoogleDrive) Files(dir string, hidden bool) ([]os.File, error) {
	panic("implement me")
}

func (g *GoogleDrive) Dirs(dir string) ([]os.File, error) {
	panic("implement me")
}

func (g *GoogleDrive) MakeDir(path string) error {
	panic("implement me")
}

func (g *GoogleDrive) MoveDir(from, to string, overwrite bool) error {
	panic("implement me")
}

func (g *GoogleDrive) CopyDir(dir, target string) error {
	panic("implement me")
}

func (g *GoogleDrive) DeleteDir(dir string) error {
	panic("implement me")
}

func (g *GoogleDrive) CleanDir(dir string) error {
	panic("implement me")
}
