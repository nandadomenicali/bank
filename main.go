package main

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

func main() {

}
