package main

import (
	"fmt"

	"github.com/nakusha/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	definition, err := dictionary.Search("second")
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
