package utils

import (
	"testing"
)

func TestSliceDiff(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"c", "d", "a"}
	diff := SliceDiff(s1, s2)
	if len(diff) != 1 && diff[0] != "b" {
		t.Fatal("error: slice diff")
	}

	var s3 []string
	diff = SliceDiff(s1, s3)
	if len(diff) != 3 {
		t.Fatal("error: slice diff")
	}
}
