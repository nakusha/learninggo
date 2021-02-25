package main

import "fmt"

func canIDrink(age int) bool {
	// if에서 사용할 수 있는 변수를 만들 수 있다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func switchCanIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	default:
		return false
	}

	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// default:
	// 	return false
	// }
}

func main() {
	isDrink := canIDrink(18)
	fmt.Println(isDrink)
}
