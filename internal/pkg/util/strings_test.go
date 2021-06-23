package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsUrl(t *testing.T) {
	require.True(t, IsUrl("https://github.com/tsundata/assistant"))
}

func TestIsIPv4(t *testing.T) {
	require.True(t, IsIPv4("127.0.0.1"))
	require.False(t, IsIPv4("172.888.2.1"))
}

func TestGeneratePassword(t *testing.T) {
	pwd := GeneratePassword(32, "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar")
	require.Len(t, pwd, 32)

	pwd2 := GeneratePassword(32, "")
	require.Len(t, pwd2, 0)
}

func TestGenerateUUID(t *testing.T) {
	uuid, err := GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, uuid, 36)
}

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(32, "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar")
	}
}

func TestExtractUUID(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		expect string
	}{
		{
			"case1",
			"/page/b58ca090-06cf-4593-812a-9992a5bec526",
			"b58ca090-06cf-4593-812a-9992a5bec526",
		},
		{
			"case2",
			"/page/b58ca090-06cf-4593-812a-9992a5bec526/create",
			"b58ca090-06cf-4593-812a-9992a5bec526",
		},
		{
			"case3",
			"/page/b58ca090-06cf-4593-812a-9992a5",
			"",
		},
		{
			"case4",
			"b58ca090-06cf-4593-812a-9992a5bec526",
			"b58ca090-06cf-4593-812a-9992a5bec526",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expect, ExtractUUID(tt.path))
		})
	}
}

func TestDataMasking(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		expect string
	}{
		{
			"case1",
			"b58ca090-06cf-4593-812a-9992a5bec526",
			"b58******526",
		},
		{
			"case2",
			"b58",
			"b******8",
		},
		{
			"case2",
			"c",
			"c******c",
		},
		{
			"case2",
			"",
			"",
		},
		{
			"case2",
			"1234",
			"123******234",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expect, DataMasking(tt.data))
		})
	}
}
