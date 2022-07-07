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
		got := hello("Sanyi")
		want := "Hello, Sanyi"

		handleAssertion(t, got, want)
	})
	t.Run("default to world if empty", func(t *testing.T) {
		got := hello("")
		want := "Hello, world"

		handleAssertion(t, got, want)
	})
}
