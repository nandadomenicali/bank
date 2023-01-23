package main

import (
	a "bank/account"
	"fmt"
)

func main() {
	fernandaAccount := a.Account{AccountHolder: "Fernanda", Balance: 500}
	userTestAccount := a.Account{AccountHolder: "UserTester", Balance: 100}

	status := fernandaAccount.Transfer(100, &userTestAccount)

	fmt.Println(status)
	fmt.Println(fernandaAccount)
	fmt.Println(userTestAccount)
}
