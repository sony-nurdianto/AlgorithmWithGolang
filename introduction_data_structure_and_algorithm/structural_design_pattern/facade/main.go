package main

import "fmt"

type Account struct {
	id         string
	acountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.acountType = accountType
	return account
}

func (account *Account) getById(_ string) *Account {
	fmt.Println("Get Account By Id")
	return account
}

func (account *Account) deleteById(_ string) {
	fmt.Println("Delete Account by Id")
}

type Customer struct {
	name string
	id   string
}

func (customer *Customer) create(name string) *Customer {
	customer.name = name
	return customer
}

type Transaction struct {
	id            string
	amount        float64
	srcAccountId  string
	destAccountId string
}

func (transaction *Transaction) create(
	srcAccountId string,
	destAccountId string,
	amount float64,
) *Transaction {
	fmt.Println("Create Transaction")

	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount

	return transaction
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{
		account:     &Account{},
		customer:    &Customer{},
		transaction: &Transaction{},
	}
}

func (facade *BranchManagerFacade) createCustomerAccount(
	customerName string,
	accountType string,
) (*Customer, *Account) {
	customer := facade.customer.create(customerName)
	account := facade.account.create(accountType)

	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(
	srcAccountId string,
	destAccountId string,
	amount float64,
) *Transaction {
	transaction := facade.transaction.create(srcAccountId, destAccountId, amount)

	return transaction
}

func main() {
	facade := NewBranchManagerFacade()
	var customer *Customer
	var account *Account

	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")

	fmt.Println(customer.name)
	fmt.Println(account.acountType)

	transaction := facade.createTransaction("21456", "87345", 1000)
	fmt.Println(transaction.amount)
}
