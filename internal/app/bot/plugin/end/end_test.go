package end

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnd(t *testing.T) {
	p := End{}
	input := "test"
	out, err := p.Run(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}
