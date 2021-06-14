package util

import (
	"crypto/sha1"
	"fmt"
)

func SHA1(url string) string {
	s := sha1.New()
	s.Write(StringToByte(url))
	bs := s.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
