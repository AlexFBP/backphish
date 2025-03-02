package common

import (
	"testing"
)

func TestParseOptions(t *testing.T) {
	opts := "ble,bla,bli"
	var m MirrorOptions
	m.ParseOptions(opts)
	if len(m.options) != 3 || m.options[1] != "bla" {
		t.Fatalf("wrong second object")
	}
}

func TestHasOptions(t *testing.T) {
	string_opts := "back,yes"
	var m MirrorOptions
	m.ParseOptions(string_opts)
	if !m.HasOption("yes ") {
		t.Fatalf("Expected option not found")
	}
}
