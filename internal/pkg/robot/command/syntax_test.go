package command

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSyntax(t *testing.T) {
	s := NewSyntax([]rune("subs open [string] [number] [any] [bool]"))
	token, err := s.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "subs", token.Value)

	token, err = s.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "open", token.Value)

	token, err = s.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "string", token.Value)

	token, err = s.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "number", token.Value)

	token, err = s.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "any", token.Value)

	token, err = s.GetNextToken()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "bool", token.Value)
}

func TestCheck(t *testing.T) {
	define := "subs open [string] [number] [any] [bool]"
	tests := []struct {
		define string
		input  string
		want   bool
	}{
		{define, "subs open abc 123 demo true", true},
		{define, "subs open abc no_num demo true", false},
		{define, "subs open abc 123 demo t", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Syntax Check #%d", i), func(t *testing.T) {
			a, err := ParseCommand(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			c, err := SyntaxCheck(tt.define, a)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, tt.want, c)
		})
	}
}
