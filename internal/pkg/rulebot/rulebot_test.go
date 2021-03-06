package rulebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewContext(t *testing.T) {
	ctx := NewContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	require.Nil(t, ctx.Message())
	require.Nil(t, ctx.Middle())
	require.Nil(t, ctx.Subscribe())
	require.Nil(t, ctx.Workflow())
	require.Nil(t, ctx.Storage())
	require.Nil(t, ctx.Todo())
	require.Nil(t, ctx.User())
	require.Nil(t, ctx.NLP())
	require.Nil(t, ctx.GetConfig())
	require.Nil(t, ctx.GetRedis())
	require.Nil(t, ctx.GetLogger())
}

func TestRuleBot(t *testing.T) {
	b := New(nil)
	var opts []Option
	b.SetOptions(opts...)
	rb := b.Process("test")
	require.Equal(t, "", rb.Name())
	require.Len(t, rb.MessageProviderOut(), 0)

	rb2 := b.Process("help")
	require.Len(t, rb2.MessageProviderOut(), 1)
}

func TestRegisterRuleset(t *testing.T) {
	RegisterRuleset(nil)
}
