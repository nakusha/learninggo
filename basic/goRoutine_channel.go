package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [3]string{"ys", "nakusha", "nak"}

	for _, person := range people {
		go isKind(person, c)

	}
	for i := 0; i < len(people); i++ {
		fmt.Println("Waitting for ", i)
		fmt.Println(<-c)
	}
}

func isKind(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is Kind"
}

// goroutines
func counter(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is kind", i)
		time.Sleep(time.Second)
	}
}
