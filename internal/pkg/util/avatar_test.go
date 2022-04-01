package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvatar(t *testing.T) {
	data, err := Avatar("d")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(data) > 0)
}
