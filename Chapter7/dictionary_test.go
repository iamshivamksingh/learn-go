package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"test": "this is just a test",
		"name": "john testing this code"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func TestCustomSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func TestCustomSearchWithErr(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrWordNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("should find added word: ", err)
		}
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		if err == nil {
			t.Fatal("should find added word: ", err)
		}
		assertStrings(t, err.Error(), ErrWordAlreadyExists.Error())
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	newDefinition := "new definition"

	dictionary := Dictionary{word: definition}

	dictionary.Update(word, newDefinition)

	got, _ := dictionary.Search(word)

	if got != newDefinition {
		t.Errorf("got %q want %q", got, newDefinition)
	}
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
