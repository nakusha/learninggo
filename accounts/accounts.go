package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// receiver a 를 이요하여 변수를 수정할 수 있다. struct에 속한 함수
// 선언된 변수를 사용하는것이 아니라 복사해서 사용하기 떄문에 직접적 변화에 작동하지 않는다
// 그래서 포인터를 통해서 해당 주소의 값을 직접 변경해야함
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of user account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw balance with Error handling
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("balnce is not enough")
	}
	a.balance -= amount
	return nil
}
