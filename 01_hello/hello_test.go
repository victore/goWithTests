package main

import "testing"

// The test function must start with the word Test
// The test function takes one argument only t *testing.T
func TestHello(t *testing.T) {
	// got := Hello()
	// want := "Hello, world"
	got := Hello("Victor")
	want := "Hello, Victor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
