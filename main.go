package main

import (
	"bank/account"
	"fmt"
)

func main() {
	fernandaAccount := account.DependentAccount{}
	fernandoAccount := account.Account{}

	fernandaAccount.CashDeposit(100)
	fernandaAccount.CashOut(50)

	fmt.Println(fernandaAccount)
	fmt.Println(fernandoAccount)

}
