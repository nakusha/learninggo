package main

import (
	"fmt"

	"github.com/nakusha/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Yeonsu")
	account.Deposit(100)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(10)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Println(account)
}
