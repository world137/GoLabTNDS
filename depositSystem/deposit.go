package depositsystem

import "GoLab/account"

type DepositSystem struct {
	AccountList []account.Account
}

func (d *DepositSystem) Deposit(accountId string, amount int) error { // error is optinal
	return nil
}

func (d *DepositSystem) WithDraw(accountId string, amount int) error {
	return nil
}

func (d *DepositSystem) PrintAllAccountData() []account.Account {
	return nil
}

func (d *DepositSystem) Transfer(from, to string, amount int) error {
	return nil
}

func (d *DepositSystem) CreateAccount(accountId string) *account.Account {
	return nil
}
