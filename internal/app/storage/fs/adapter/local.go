package adapter

import (
	"os"
)

type Local struct {
	Dir    string
	Domain string
}

func (l *Local) AbsolutePath(path string) string {
	return l.Dir + path
}

func (l *Local) FullPath(path string) string {
	return l.Domain + path
}

func (l *Local) Exists(path string) (bool, error) {
	panic("implement me")
}

func (l *Local) Get(path string, lock bool) ([]byte, error) {
	panic("implement me")
}

func (l *Local) Hash(path string) (string, error) {
	panic("implement me")
}

func (l *Local) Put(path string, contents []byte, _ bool) error {
	f, err := os.Create(l.AbsolutePath(path))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = f.Write(contents)
	if err != nil {
		return err
	}
	return nil
}

func (l *Local) Replace(path string, content []byte) error {
	panic("implement me")
}

func (l *Local) Append(path string, data []byte) error {
	panic("implement me")
}

func (l *Local) Delete(path string) error {
	panic("implement me")
}

func (l *Local) Move(path, target string) error {
	panic("implement me")
}

func (l *Local) Copy(path, target string) error {
	panic("implement me")
}

func (l *Local) Size(path string) (int, error) {
	panic("implement me")
}

func (l *Local) IsDir(path string) (bool, error) {
	panic("implement me")
}

func (l *Local) IsFile(path string) (bool, error) {
	panic("implement me")
}

func (l *Local) Glob(pattern string, flags int) ([]os.File, error) {
	panic("implement me")
}

func (l *Local) Files(dir string, hidden bool) ([]os.File, error) {
	panic("implement me")
}

func (l *Local) Dirs(dir string) ([]os.File, error) {
	panic("implement me")
}

func (l *Local) MakeDir(path string) error {
	panic("implement me")
}

func (l *Local) MoveDir(from, to string, overwrite bool) error {
	panic("implement me")
}

func (l *Local) CopyDir(dir, target string) error {
	panic("implement me")
}

func (l *Local) DeleteDir(dir string) error {
	panic("implement me")
}

func (l *Local) CleanDir(dir string) error {
	panic("implement me")
}
