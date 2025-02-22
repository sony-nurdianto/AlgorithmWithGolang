package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}

type Account struct {
	details      *AccountDetails
	CustomerName string
}

func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{
		id:          id,
		accountType: accountType,
	}
}

func (account *Account) getId() string {
	return account.details.id
}

func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	account := &Account{
		CustomerName: "John smith",
	}
	account.setDetails("4532", "current")

	if jsonAccount, err := json.MarshalIndent(account, "", "   "); err == nil {
		fmt.Printf("Private class Hidden: %s\n", string(jsonAccount))
	}

	fmt.Println("Account Id: ", account.getId())
	fmt.Println("Account Type: ", account.getAccountType())
}
