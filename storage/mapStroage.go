package storage

import (
	"GoLab/account"
	"fmt"
	"sort"
)

type mapStroage struct {
	data map[string]account.Account
}

func InitMapStorage() *mapStroage {
	return &mapStroage{
		data: make(map[string]account.Account),
	}
}
func (m *mapStroage) Create(account account.Account) error {
	if account.AccountId == "" {
		return fmt.Errorf("No account id")
	}
	m.data[account.AccountId] = account
	return nil
}
func (m *mapStroage) Read(accountId string) (account.Account, error) {
	if accountId == "" {
	}
	return m.data[accountId], nil
}
func (m *mapStroage) ReadAll() ([]account.Account, error) {
	var returnArray []account.Account
	for _, v := range m.data {
		returnArray = append(returnArray, v)
	}

	sort.SliceStable(returnArray, func(i, j int) bool {
		return returnArray[i].AccountId < returnArray[j].AccountId //condition
	})

	return returnArray, nil
}
func (m *mapStroage) Update(account account.Account) error {
	if account.AccountId == "" {
		return fmt.Errorf("No account id")
	}
	m.data[account.AccountId] = account
	return nil
}
func (m *mapStroage) Delete(account account.Account) error {
	if account.AccountId == "" {
		return fmt.Errorf("No account id")
	}
	delete(m.data, account.AccountId)
	return nil
}
