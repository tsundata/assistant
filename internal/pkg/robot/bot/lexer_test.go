package bot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLexer(t *testing.T) {
	l := NewLexer([]rune("test  @user @bot #tag1  #tag2 /version info/ / help  / #123"))
	// string
	token, err := l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, StringToken, token.Type)
	// object
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, ObjectToken, token.Type)
	// object
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, ObjectToken, token.Type)
	// tag
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, TagToken, token.Type)
	// tag
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, TagToken, token.Type)
	// command
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, CommandToken, token.Type)
	// command
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, CommandToken, token.Type)
	// message
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, MessageToken, token.Type)
	// eof
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, EOFToken, token.Type)
}

func TestParseText(t *testing.T) {
	c, err := ParseText("subs list")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c, 2)

	c, err = ParseText("subs @open #abc /help/ #123")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c, 5)

	require.Equal(t, "subs", c[0].Value)
	require.Equal(t, "open", c[1].Value)
	require.Equal(t, "abc", c[2].Value)
	require.Equal(t, "help", c[3].Value)
	require.Equal(t, "123", c[4].Value)
}
