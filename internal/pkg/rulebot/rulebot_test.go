package rulebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRuleBot(t *testing.T) {
	b := New(nil)
	rb := b.Process("test")
	require.Equal(t, "", rb.Name())
	require.Len(t, rb.MessageProviderOut(), 0)

	rb2 := b.Process("help")
	require.Len(t, rb2.MessageProviderOut(), 1)
}
