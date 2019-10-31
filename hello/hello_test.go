package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jeff", "English")
		want := "Hello, Jeff!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, World! when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in spanish", func(t *testing.T) {
		got := Hello("Jeff", "Spanish")
		want := "Hola, Jeff!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in french", func(t *testing.T) {
		got := Hello("Jeff", "French")
		want := "Bonjour, Jeff!"
		assertCorrectMessage(t, got, want)
	})

}
