package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you looked")

func (d Dictionary) Search(word string) (string, error) {
  definintion, ok := d[word]
  if !ok {
    return "", ErrNotFound
  }

  return definintion, nil
}

func (d Dictionary) Add(word string, definition string) {
  d[word] = definition
}

