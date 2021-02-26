package mydict

import "errors"

// type를 이용한 선언 예시
// type Money int
// Money(1) = int 1

// Dcitionary
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search Dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
