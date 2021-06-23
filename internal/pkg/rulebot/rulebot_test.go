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
}
