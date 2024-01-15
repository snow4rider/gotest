package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update non-existing word")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search searches for a word in the dictionary and returns its definition.
// If the word is not found, it returns an error.
func (d Dictionary) Search(word string) (string, error) {

	// Check if the word exists in the dictionary
	definition, ok := d[word]

	// If the word is not found, return an error
	if !ok {
		return "", ErrNotFound
	}

	// Return the definition
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
