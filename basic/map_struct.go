package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// map
	// testMap := map[string]string{
	// 	"name": "yeonsu",
	// 	"age":  "20",
	// }

	// for key, value := range testMap {
	// 	fmt.Println(key, value)
	// }

	//struct
	favFood := []string{"ramen", "steak"}
	me := person{name: "yeonsu", age: 37, favFood: favFood}
	fmt.Println(me.name, me.age)
}
