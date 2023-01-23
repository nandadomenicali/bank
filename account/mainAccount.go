package account

type Account struct {
	AccountHolder string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (a *Account) CashOut(cashOutAmount float64) string {
	cashOut := cashOutAmount > 0 && cashOutAmount <= a.Balance

	if cashOut {
		a.Balance -= cashOutAmount
		return "Cash out made successfully"
	} else {
		return "Insufficient funds"
	}
}

func (a *Account) CashDeposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		a.Balance += depositAmount
		return "Deposit made successfully. Your balance is:", a.Balance
	} else {
		return "Deposit amount cannot be less than 0", a.Balance
	}
}

func (a *Account) Transfer(transferAmount float64, destinationAccount *Account) bool {
	if transferAmount < a.Balance && transferAmount > 0 {
		a.Balance -= transferAmount
		destinationAccount.CashDeposit(transferAmount)
		return true
	} else {
		return false
	}
}
