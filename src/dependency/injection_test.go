package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hans")

	got := buffer.String()
	want := "Hello, Hans"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
