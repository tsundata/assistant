package util

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBoolInt(t *testing.T) {
	require.Equal(t, BoolInt(true), 1)
	require.Equal(t, BoolInt(false), 0)
}

type demo struct {
	A int64   `db:"a"`
	B float64 `db:"b"`
	C string  `db:"c"`
	D bool    `db:"d"`

	E int64 `db:"e"`
	F string `db:"f"`
}

func TestInject(t *testing.T) {
	now := time.Now()
	d := &demo{}
	Inject(d, map[string]interface{}{
		"a": int64(123),
		"b": 456.789,
		"c": "test",
		"d": int64(1),

		"e": float64(159),
		"f": now,
	})

	require.Equal(t, d.A, int64(123))
	require.Equal(t, d.B, 456.789)
	require.Equal(t, d.C, "test")
	require.Equal(t, d.D, true)
	require.Equal(t, d.E, int64(159))
	require.Equal(t, d.F, now.Format("2006-01-02 15:04:05"))
}
