package main

import "fmt"

type Account struct {
	accountHolder string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (a *Account) cashOut(cashOutAmount float64) string {
	cashOut := cashOutAmount > 0 && cashOutAmount <= a.balance

	if cashOut {
		a.balance -= cashOutAmount
		return "Cash out made successfully"
	} else {
		return "Insufficient funds"
	}
}

func (a *Account) cashDeposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		a.balance += depositAmount
		return "Deposit made successfully. Your balance is:", a.balance
	} else {
		return "Deposit amount cannot be less than 0", a.balance
	}
}

func (a *Account) transfer(transferAmount float64, destinationAccount *Account) bool {
	if transferAmount < a.balance && transferAmount > 0 {
		a.balance -= transferAmount
		destinationAccount.cashDeposit(transferAmount)
		return true
	} else {
		return false
	}
}

func main() {
	fernandaAccount := Account{accountHolder: "Fernanda", balance: 500}
	userTestAccount := Account{accountHolder: "UserTester", balance: 100}

	status := fernandaAccount.transfer(100, &userTestAccount)

	fmt.Println(status)
	fmt.Println(fernandaAccount)
	fmt.Println(userTestAccount)
}
