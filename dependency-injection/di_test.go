package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Flo")

	got := buffer.String()
	want := "Hello, Flo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
