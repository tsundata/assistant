package adapter

import (
	"os"
)

type Dropbox struct{}

func (d *Dropbox) Exists(path string) (bool, error) {
	panic("implement me")
}

func (d *Dropbox) Get(path string, lock bool) ([]byte, error) {
	panic("implement me")
}

func (d *Dropbox) Hash(path string) (string, error) {
	panic("implement me")
}

func (d *Dropbox) Put(path string, contents []byte, lock bool) error {
	panic("implement me")
}

func (d *Dropbox) Replace(path string, content []byte) error {
	panic("implement me")
}

func (d *Dropbox) Append(path string, data []byte) error {
	panic("implement me")
}

func (d *Dropbox) Delete(path string) error {
	panic("implement me")
}

func (d *Dropbox) Move(path, target string) error {
	panic("implement me")
}

func (d *Dropbox) Copy(path, target string) error {
	panic("implement me")
}

func (d *Dropbox) Size(path string) (int, error) {
	panic("implement me")
}

func (d *Dropbox) IsDir(path string) (bool, error) {
	panic("implement me")
}

func (d *Dropbox) IsFile(path string) (bool, error) {
	panic("implement me")
}

func (d *Dropbox) Glob(pattern string, flags int) ([]os.File, error) {
	panic("implement me")
}

func (d *Dropbox) Files(dir string, hidden bool) ([]os.File, error) {
	panic("implement me")
}

func (d *Dropbox) Dirs(dir string) ([]os.File, error) {
	panic("implement me")
}

func (d *Dropbox) MakeDir(path string) error {
	panic("implement me")
}

func (d *Dropbox) MoveDir(from, to string, overwrite bool) error {
	panic("implement me")
}

func (d *Dropbox) CopyDir(dir, target string) error {
	panic("implement me")
}

func (d *Dropbox) DeleteDir(dir string) error {
	panic("implement me")
}

func (d *Dropbox) CleanDir(dir string) error {
	panic("implement me")
}
