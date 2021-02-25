package main

import (
	"fmt"

	"github.com/nakusha/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Yeonsu")
	fmt.Println(*account)
}
