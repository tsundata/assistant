package collection

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/internal/pkg/app"
	"math"
	"testing"
	"time"
)

func TestNewBinSet(t *testing.T) {
	rdb, err := CreateRedisClient(app.Message)
	if err != nil {
		t.Fatal(err)
	}
	s := NewBinSet(rdb, fmt.Sprintf("bin_set:%s", time.Now().String()))

	s.SetTo(1, 1)
	b := s.Test(1)
	require.True(t, b)

	s.SetTo(2, 0)
	b2 := s.Test(2)
	require.False(t, b2)

	s.SetTo(1, 0)
	b3 := s.Test(1)
	require.False(t, b3)

	s.SetTo(math.MaxInt8, 1)
	b4 := s.Test(math.MaxInt8)
	require.True(t, b4)
}
