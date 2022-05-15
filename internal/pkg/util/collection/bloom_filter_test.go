package collection

import (
	"crypto/rand"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/enum"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	rdb, err := CreateRedisClient(enum.Message)
	if err != nil {
		t.Fatal(err)
	}
	randId, _ := rand.Read(nil)
	f := NewBloomFilter(rdb, fmt.Sprintf("bloom_filter:%d", randId), 10000, 10)

	f.Add("abc")
	b := f.Lookup("abc")
	require.True(t, b)

	f.Add("abc123456")
	b2 := f.Lookup("abc1234567")
	require.False(t, b2)

	f.Add("")
	b3 := f.Lookup("")
	require.True(t, b3)

	k := "abc123456abc123456qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm"
	f.Add(k)
	b4 := f.Lookup(k)
	require.True(t, b4)

	f.Add("abc123456")
	b5 := f.Lookup("abc1234567")
	require.False(t, b5)
}
