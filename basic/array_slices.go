package main

import "fmt"

func main() {
	// array는 length 제한이있지만 slice는 length제한이 없다.
	// array
	// names := [5]string{"test", "me", "my"}
	// names[3] = "kakak"
	// names[4] = "hahaha"
	// names[5] = "zzzzz"

	// slice
	names := []string{"test", "me", "my"}

	names = append(names, "test")
	fmt.Println(names)
}
