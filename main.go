package main

import (
	"bank/accounts"
	"fmt"
)

func payInvoice(account verifyAccount, invoiceValue float64) {
	account.CashOut(invoiceValue)
}

type verifyAccount interface {
	CashOut(value float64) string
}

func main() {
	fernandaAccount := accounts.Account{}
	fernandaAccount.CashDeposit(100)
	payInvoice(&fernandaAccount, 50)

	fmt.Println(fernandaAccount.GetBalance())
}
