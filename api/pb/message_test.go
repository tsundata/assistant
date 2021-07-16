package pb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessage_IsMessageOfAction(t *testing.T) {
	m := Message{Text: `#!action
test`}
	require.True(t, m.IsMessageOfAction())

	m2 := Message{Text: ""}
	require.False(t, m2.IsMessageOfAction())
}

func TestMessage_RemoveActionFlag(t *testing.T) {
	m := Message{Text: `#!action
test`}
	require.Equal(t, "test", m.RemoveActionFlag())
}
