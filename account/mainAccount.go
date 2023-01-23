package account

import "bank/customers"

type Account struct {
	AccountHolder               customers.Customer
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (a *Account) CashOut(cashOutAmount float64) string {
	cashOut := cashOutAmount > 0 && cashOutAmount <= a.balance

	if cashOut {
		a.balance -= cashOutAmount
		return "Cash out made successfully"
	} else {
		return "Insufficient funds"
	}
}

func (a *Account) CashDeposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		a.balance += depositAmount
		return "Deposit made successfully. Your balance is:", a.balance
	} else {
		return "Deposit amount cannot be less than 0", a.balance
	}
}

func (a *Account) Transfer(transferAmount float64, destinationAccount *Account) bool {
	if transferAmount < a.balance && transferAmount > 0 {
		a.balance -= transferAmount
		destinationAccount.CashDeposit(transferAmount)
		return true
	} else {
		return false
	}
}

func (a *Account) GetBalance() float64 {
	return a.balance
}
