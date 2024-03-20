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
  t.Run("new word", func(t *testing.T) {

    dictionary := Dictionary{}
    word := "test"
    definition := "this is just test"

    err := dictionary.Add(word, definition)

    assertError(t, err, nil)
    assertDefintion(t, dictionary, word, definition)
  })

  t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is just test"
    dictionary := Dictionary{word: definition}

    err := dictionary.Add(word, "new test")
    assertError(t, err, ErrWordExists)
    assertDefintion(t, dictionary, word, definition)
  })
}

func TestUpdate(t *testing.T) {
  t.Run("existing word", func(t *testing.T) {

  word := "test"
  definition := "this is just test"
  dictionary := Dictionary{word: definition}
  newDefinition := "new definition"

  err := dictionary.Update(word, newDefinition)

    assertError(t, err, nil)
  assertDefintion(t, dictionary, word, newDefinition)
  })

  t.Run("new word", func(t *testing.T) {

  word := "test"
  definition := "this is just test"
  dictionary := Dictionary{}

    err := dictionary.Update(word, definition)
    assertError(t, err,ErrWordDoesNotExist)
  })

}

func TestDelete(t *testing.T) {
  word := "test"
  definition := "this is test"
  dictionary := Dictionary{word: definition}

  dictionary.Delete(word)

  _, err := dictionary.Search(word)
  if err != ErrNotFound {
    t.Errorf("Expected %q to be deleted", word)
  }
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

func assertDefintion (t testing.TB, dictionary Dictionary, word, definition string) {
  t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

