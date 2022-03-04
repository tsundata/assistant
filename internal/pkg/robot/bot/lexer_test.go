package bot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLexer(t *testing.T) {
	l := NewLexer([]rune("test @user @bot #tag1 #tag2"))
	token, err := l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, StringToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, ObjectToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, ObjectToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, TagToken, token.Type)
	token, err = l.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, TagToken, token.Type)
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

	c, err = ParseText("subs @open #abc")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, c, 3)

	require.Equal(t, "subs", c[0].Value)
	require.Equal(t, "open", c[1].Value)
	require.Equal(t, "abc", c[2].Value)
}
