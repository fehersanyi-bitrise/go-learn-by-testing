package main

import "testing"

func TestHello(t *testing.T) {

	handleAssertion := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Sanyi", "English")
		want := "Hello, Sanyi"

		handleAssertion(t, got, want)
	})
	t.Run("default to world if empty", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello, world"

		handleAssertion(t, got, want)
	})
	t.Run("hungarian support", func(t *testing.T) {
		got := hello("Bélám", "Hungarian")
		want := "Szevasz, Bélám"

		handleAssertion(t, got, want)
	})
	t.Run("french support", func(t *testing.T) {
		got := hello("Belle", "French")
		want := "Bonjour, Belle"

		handleAssertion(t, got, want)
	})
}
