package collection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderedDict(t *testing.T) {
	od := NewOrderedDict()

	od.Set("a", 1)
	od.Set("b", 2)

	num := 0
	for range od.Iterate() {
		num++
	}
	require.Equal(t, 2, num)
	require.Equal(t, 1, od.Get("a"))
	require.Equal(t, 2, od.Get("b"))

	od.Remove("a")
	od.Remove("b")

	num = 0
	for range od.Iterate() {
		num++
	}
	require.Equal(t, 0, num)
}
