package collection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()
	n1 := ll.Append(1)
	n2 := ll.Append(2)

	num := 0
	for range ll.Iterate() {
		num++
	}
	require.Equal(t, 2, num)

	ll.Remove(n1)
	ll.Remove(n2)

	num = 0
	for range ll.Iterate() {
		num++
	}
	require.Equal(t, 0, num)

}
