package rulebot

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRuleBot(t *testing.T) {
	b := New(nil)
	var opts []Option
	b.SetOptions(opts...)
	rb := b.Process(context.Background(), "test")
	require.Equal(t, "", rb.Name())
	require.Len(t, rb.MessageProviderOut(), 0)

	rb2 := b.Process(context.Background(), "help")
	require.Len(t, rb2.MessageProviderOut(), 1)
}

func TestRegisterRuleset(t *testing.T) {
	RegisterRuleset(nil)
}
