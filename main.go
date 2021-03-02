package main

import (
	"fmt"
	"time"
)

func main() {
	go counter("ys")
	go counter("nakusha")
	counter("nak")
}

// goroutines
func counter(persion string) {
	for i := 0; i < 10; i++ {
		fmt.Println(persion, "is kind", i)
		time.Sleep(time.Second)
	}
}
