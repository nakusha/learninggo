package main

import (
	"fmt"

	"github.com/nakusha/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	word := "hello"
	dictionary.Add(word, "Greeting")
	dictionary.Update(word, "Change")
	value, _ := dictionary.Search(word)

	fmt.Println(value)

}
