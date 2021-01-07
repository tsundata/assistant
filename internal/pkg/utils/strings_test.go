package utils

import "testing"

func TestIsUrl(t *testing.T) {
	if !IsUrl("https://github.com/tsundata/assistant") {
		t.Fatal("error: don't match url")
	}
}

func TestIsIPv4(t *testing.T) {
	if !IsIPv4("127.0.0.1") {
		t.Fatal("error: don't match ip")
	}
	if IsIPv4("172.888.2.1") {
		t.Fatal("error: match ip")
	}
}

func TestGeneratePassword(t *testing.T) {
	pwd := GeneratePassword(32, "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar")
	if len(pwd) != 32 {
		t.Fatal("error: generate password")
	}
}

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(32, "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar")
	}
}
