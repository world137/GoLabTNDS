package depositSystem

import (
	"GoLab/account"
	"fmt"
)

type DepositSystem struct {
	AccountList []account.Account
}

func (d *DepositSystem) Deposit(accountId string, amount int) error { // error is optinal
	targetAccount := &account.Account{}
	for i := 0; i < len(d.AccountList); i++ {
		if d.AccountList[i].Id == accountId {
			fmt.Printf("address account 1 %p \n", &d.AccountList[i])
			targetAccount = &d.AccountList[i]
			break
		}
	}
	if targetAccount.Id == "" {
		fmt.Println("Account Not Found")
		return nil
	}
	fmt.Printf("address account 2 %p \n", targetAccount)
	targetAccount.Balance += amount
	fmt.Println(targetAccount.Balance)

	return nil
}

func (d *DepositSystem) WithDraw(accountId string, amount int) error {
	targetAccount := &account.Account{}
	for i := 0; i < len(d.AccountList); i++ {
		if d.AccountList[i].Id == accountId {
			targetAccount = &d.AccountList[i]
		}
	}
	if targetAccount.Id == "" {
		fmt.Println("Account Not Found")
		return nil
	}
	targetAccount.Balance -= amount
	return nil
}

func (d *DepositSystem) PrintAllAccountData() []account.Account {
	return d.AccountList
}

func (d *DepositSystem) Transfer(from, to string, amount int) error {
	fromAccount := &account.Account{}
	toAccount := &account.Account{}
	for i, v := range d.AccountList {
		if v.Id == from {
			fromAccount = &d.AccountList[i]
		}
		if v.Id == to {
			toAccount = &d.AccountList[i]
		}
	}
	if toAccount == fromAccount {
		// err = "Same Account"
		fmt.Println("Same Account")
		return nil
	}
	if toAccount.Id == "" || fromAccount.Id == "" {
		// err = "Account not found"
		fmt.Println("Account not found")
		fmt.Println(toAccount.Id == "")
		return nil
	}
	// if fromAccount.Balance <= 0 {
	// 	fmt.Println("Balance < 0")
	// 	return nil
	// }
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	return nil
}

func (d *DepositSystem) CreateAccount(accountId string) *account.Account {
	for _, v := range d.AccountList {
		if v.Id == accountId {
			panic("Error : Same Id")
		}
	}
	newAccount := account.Account{
		Id:      accountId,
		Balance: 0,
	}
	d.AccountList = append(d.AccountList, newAccount)
	return &newAccount
}
