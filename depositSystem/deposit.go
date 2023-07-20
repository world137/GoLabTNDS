package depositSystem

import (
	"GoLab/account"
	"fmt"
)

type DepositSystem struct {
	// AccountList []account.Account
	AccountMap map[string]account.Account
}

func (d *DepositSystem) Deposit(accountId string, amount int) error { // error is optinal
	targetAccount := account.Account{}
	// for i := 0; i < len(d.AccountList); i++ {
	// 	if d.AccountList[i].Id == accountId {
	// 		targetAccount = &d.AccountList[i]
	// 		break
	// 	}
	// }
	// if targetAccount.Id == "" {
	// 	fmt.Println("Account Not Found")
	// 	return nil
	// }
	_, ok := d.AccountMap[accountId]
	if !ok {
		return fmt.Errorf("Account Not Found")
	} else {
		targetAccount = d.AccountMap[accountId]
	}
	targetAccount.Balance += amount
	d.AccountMap[accountId] = targetAccount

	return nil
}

func (d *DepositSystem) WithDraw(accountId string, amount int) error {
	targetAccount := account.Account{}
	// for i := 0; i < len(d.AccountList); i++ {
	// 	if d.AccountList[i].Id == accountId {
	// 		targetAccount = &d.AccountList[i]
	// 	}
	// }
	_, ok := d.AccountMap[accountId]
	if !ok {
		return fmt.Errorf("Account Not Found")
	} else {
		targetAccount = d.AccountMap[accountId]
	}
	// if targetAccount.Id == "" {
	// 	fmt.Println("Account Not Found")
	// 	return nil
	// }
	targetAccount.Balance -= amount
	d.AccountMap[accountId] = targetAccount

	return nil
}

func (d *DepositSystem) PrintAllAccountData() map[string]account.Account {
	// return d.AccountList
	// returnArray := []account.Account{}
	// for _, v := range d.AccountMap {
	// 	returnArray = append(returnArray, v)
	// }
	return d.AccountMap
}

func (d *DepositSystem) Transfer(from, to string, amount int) error {
	// fromAccount := account.Account{}
	// toAccount := account.Account{}
	// for i, v := range d.AccountList {
	// 	if v.Id == from {
	// 		fromAccount = &d.AccountList[i]
	// 	} else if v.Id == to {
	// 		toAccount = &d.AccountList[i]
	// 	}
	// }

	// if toAccount.Id == "" || fromAccount.Id == "" {
	// 	// err = "Account not found"
	// 	fmt.Println("Account not found")
	// 	fmt.Println(toAccount.Id == "")
	// 	return nil
	// }
	// if fromAccount.Balance <= 0 {
	// 	fmt.Println("Balance < 0")
	// 	return nil
	// }

	// fromAccount.Balance -= amount
	// toAccount.Balance += amount

	_, ok := d.AccountMap[from]
	if !ok {
		return fmt.Errorf("account not found")
	}
	_, ok2 := d.AccountMap[to]
	if !ok2 {
		return fmt.Errorf("account not found")
	}

	d.WithDraw(from, amount)
	d.Deposit(to, amount)

	return nil
}

func (d *DepositSystem) CreateAccount(acc account.Account) *account.Account {
	// for _, v := range d.AccountList {
	// 	if v.Id == accountId {
	// 		panic("Error : Same Id")
	// 	}
	// }
	_, ok := d.AccountMap[acc.AccountId]
	if ok {
		fmt.Println("same account")
	}
	newAccount := account.Account{
		AccountId: acc.AccountId,
		Balance:   0,
	}
	// d.AccountList = append(d.AccountList, newAccount)
	d.AccountMap[acc.AccountId] = newAccount
	return &newAccount
}
