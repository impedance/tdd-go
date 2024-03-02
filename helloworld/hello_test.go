package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Chris", "")
    want := "Hello, Chris"

    assertCorrectMessage(t, got, want)
  })
  t.Run("Hello workld when empty stirng supplied", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"

    assertCorrectMessage(t, got, want)
  })

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Elodie", "Spanish")
    want := "Hola, Elodie"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in French", func(t *testing.T) {
    got := Hello("Sanya", "French")
    want := "Bonjour, Sanya"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
