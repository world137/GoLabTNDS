package storage

import (
	"GoLab/account"
)

type psqlStorage struct {
}

func InitPsqlStorage() *psqlStorage {
	return &psqlStorage{}
}
func (m *psqlStorage) Create(account account.Account) error {

	return nil
}
func (m *psqlStorage) Read(accountId string) (account.Account, error) {

	return account.Account{}, nil
}
func (m *psqlStorage) ReadAll() ([]account.Account, error) {

	return []account.Account{}, nil
}
func (m *psqlStorage) Update(account account.Account) error {

	return nil
}
func (m *psqlStorage) Delete(account account.Account) error {
	return nil
}
