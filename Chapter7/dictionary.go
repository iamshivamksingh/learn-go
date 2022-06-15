package main

// var m map[string]string -
// never initialize a map like this as m currently nil and accessing it will cause runtime panic
// do either
// var dictionary = map[string]string{}
// or
// var dictionary = make(map[string]string)
// both approaches creates an empty hash map and point dictionary at it.

type Dictionary map[string]string
type DictionaryErr string // creating our own DictionaryErr
// type which implements error interface

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordNotFound      = DictionaryErr("count not found the word you were looking for")
	ErrWordAlreadyExists = DictionaryErr("cannot add word because it already exists")
)

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
