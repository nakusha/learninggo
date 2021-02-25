package main

import "fmt"

func superAdd(numbers ...int) int {
	// range는 index와 number를 같이 return 한다.
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	total := 0

	for _, number := range numbers {
		// total = total + number
		total += number
	}

	return total
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}
