package storage

import "GoLab/account"

type StroageProvider interface {
	Create(account.Account) error
	Read(string) (account.Account, error)
	ReadAll() ([]account.Account, error)
	Update(account.Account) error
	Delete(account.Account) error
}
