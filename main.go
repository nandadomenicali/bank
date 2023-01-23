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

func main() {
	myAccount := Account{}
	myAccount.accountHolder = "Fernanda"
	myAccount.balance = 500

	fmt.Println(myAccount.balance)
	status, value := myAccount.cashDeposit(100)
	fmt.Println(status, value)
}
