package utils

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	defer Duration(time.Now(), "TestDuration")
}
