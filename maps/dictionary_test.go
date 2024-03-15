package main

import "testing"

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just test"}


  t.Run("known word", func(t *testing.T) {

  got, _ := dictionary.Search("test")
  want := "this is just test"
  assertStrings(t, got, want)
  })

  t.Run("unknown word", func(t *testing.T) {
    _, got := dictionary.Search("unknown")

    assertError(t, got, ErrNotFound)
  })
}

func TestAdd(t *testing.T) {
  dictionary := Dictionary{}
  dictionary.Add("test", "this is just test")

  want := "this is just test"
  got, err := dictionary.Search("test")

  if err != nil {
    t.Fatal("shold find added word: ", err)
  }
  assertStrings(t, got, want)
}

func assertError(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
    t.Errorf("got error %q want %q", got, want)
  }
}

func assertStrings(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }

}
