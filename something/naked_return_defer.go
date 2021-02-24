package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int (return형) {
//두개의 형이 같은경우
func multiply(a, b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// naked return
func lenAndUpper(name string) (length int, upper string) {
	//defer 함수가 실행이 끝난후 실행되는것(callback같네?)
	// 변수넣고는 안됨
	defer fmt.Println("Done")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func repeatWord(words ...string) {
	fmt.Println(words)
}

func main() {
	// 축약형은 함수안에서 변수만 사용 가능
	// var name string = "nakusha"
	// name := "nakusha"

	// fmt.Println(multiply(2, 2))
	// return이 여러가지인경우 안쓰는것은 _ 로 할 수 있다.
	totalLength, upper := lenAndUpper("yeonsu")
	fmt.Println(totalLength, upper)

	// repeatWord("test", "repeat", "really", "how", "Array")
}
