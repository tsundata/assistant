package collection

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/enum"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNewBinSet(t *testing.T) {
	rand.Seed(int64(time.Now().Second()))
	rdb, err := CreateRedisClient(enum.Message)
	if err != nil {
		t.Fatal(err)
	}
	s := NewBinSet(rdb, fmt.Sprintf("bin_set:%d", rand.Int63()))

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
