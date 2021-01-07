package utils

import (
	"bytes"
	"testing"
)

func TestByteToString(t *testing.T) {
	if ByteToString([]byte("Test")) != "Test" {
		t.Error("error ByteToString")
	}
}

func TestStringToByte(t *testing.T) {
	if b := StringToByte("Test"); !bytes.Equal(b, []byte("Test")) {
		t.Error("error StringToByte")
	}
}
