package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world!"
	got := Hello()
	if got != expected {
		t.Errorf("Hello() = %q; want %q", got, expected)
	}
}
