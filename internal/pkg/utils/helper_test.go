package utils

import (
	"github.com/tsundata/assistant/internal/pkg/model"
	"testing"
)

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

func TestIsMessageOfScript(t *testing.T) {
	if !IsMessageOfScript(`#!script:js
	var a = 1;
	console.log(a);
`) {
		t.Fatal("error: match script")
	}
}

func TestIsMessageOfAction(t *testing.T) {
	if !IsMessageOfAction(`#!action
	trigger message.id > 0 do something`) {
		t.Fatal("error: match script")
	}
}

func TestMessageScriptKind(t *testing.T) {
	if MessageScriptKind(`#!action`) != model.MessageScriptOfUndefined {
		t.Fatal("error: script kind")
	}
	if MessageScriptKind(`#!script:javascript`) != model.MessageScriptOfJavascript {
		t.Fatal("error: script kind")
	}
	if MessageScriptKind(`#!script:flowscript`) != model.MessageScriptOfFlowscript {
		t.Fatal("error: script kind")
	}
}

func TestSliceDiff(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"c", "d", "a"}
	diff := SliceDiff(s1, s2)
	if len(diff) != 1 && diff[0] != "b" {
		t.Fatal("error: slice diff")
	}

	var s3 []string
	diff = SliceDiff(s1, s3)
	if len(diff) != 3 {
		t.Fatal("error: slice diff")
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
