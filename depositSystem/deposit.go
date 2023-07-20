package depositSystem

import (
	"GoLab/account"
	"GoLab/storage"
	"fmt"
)

// Packge depositSystem
type DepositSystem struct {
	Provider storage.StroageProvider
}

// AccountID gonna be 001, 002, 003, 004
func (d *DepositSystem) CreateAccount(acct account.Account) *account.Account {
	acctFromRead, err := d.Provider.Read(acct.AccountId)
	if acctFromRead.AccountId == "" {
		err := d.Provider.Create(acct)
		if err != nil {
			return &account.Account{}
		}
	}
	if err != nil {
		return &account.Account{}
	}

	return &acct
}

func (d *DepositSystem) Deposit(accountId string, amount int) error {
	acct, err := d.Provider.Read(accountId)
	if err != nil {
		return err
	}
	acct.AddBalance(amount)

	err = d.Provider.Update(acct)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepositSystem) WithDraw(accountId string, amount int) error {
	acct, err := d.Provider.Read(accountId)
	if err != nil {
		return err
	}
	acct.ReduceBalance(amount)

	err = d.Provider.Update(acct)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepositSystem) PrintAllAccountData() map[string]account.Account {

	return nil
}

func (d *DepositSystem) Transfer(from, to string, amount int) error {
	acctFrom, err := d.Provider.Read(from)
	if acctFrom.AccountId == "" {
		return fmt.Errorf("from account not found")
	}
	if err != nil {
		return err
	}

	acctTo, err := d.Provider.Read(from)

	if acctTo.AccountId == "" {
		return fmt.Errorf("to account not found")
	}
	if err != nil {
		return err
	}

	/* attention */
	err = d.WithDraw(from, amount)
	if err != nil {
		return err
	}
	err = d.Deposit(to, amount)
	if err != nil {
		return err
	}
	return nil
}
