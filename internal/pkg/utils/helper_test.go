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
	if MessageScriptKind(`#!script:dsl`) != model.MessageScriptOfDSL {
		t.Fatal("error: script kind")
	}
}
