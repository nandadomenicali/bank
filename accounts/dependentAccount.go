package accounts

import c "bank/customers"

type DependentAccount struct {
	AccountHolder                          c.Customer
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (a *DependentAccount) CashOut(cashOutAmount float64) string {
	cashOut := cashOutAmount > 0 && cashOutAmount <= a.balance

	if cashOut {
		a.balance -= cashOutAmount
		return "Cash out made successfully"
	} else {
		return "Insufficient funds"
	}
}

func (a *DependentAccount) CashDeposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		a.balance += depositAmount
		return "Deposit made successfully. Your balance is:", a.balance
	} else {
		return "Deposit amount cannot be less than 0", a.balance
	}
}

func (a *DependentAccount) GetBalance() float64 {
	return a.balance
}
