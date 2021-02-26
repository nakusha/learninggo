package mydict

import "errors"

// type를 이용한 선언 예시
// type Money int
// Money(1) = int 1

// Dcitionary
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("Word already Exists")
	errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant delete non-existing word")
)

// Search Dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add word to deictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}

// Update Word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, word)
	}
	return nil
}
