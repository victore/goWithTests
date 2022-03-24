package main

import "testing"

// The test function must start with the word Test
// The test function takes one argument only t *testing.T
func TestHello(t *testing.T) {
	// helper function - testing.TB is an interface that *testing.T and *testing.B both satisfy
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // is needed to tell the test suite that this method is a helper.
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Victor")
		want := "Hello, Victor"
		assertCorrectMessage(t, got, want)

	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
