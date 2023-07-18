package main

import (
	"GoLab/account"
	"GoLab/depositSystem"
	"reflect"
	"testing"
)

func TestDepositSystem_PrintAllAccountData(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func() []account.Account
		want     []account.Account
	}{
		{
			name:     "test1_create_account_and_deposit",
			testFunc: testCreateAccountAndDeposit,
			want: []account.Account{
				{
					Id:      "001",
					Balance: 254,
				},
			},
		},
		{
			name:     "test2_create_multiple_accounts_and_deposit_withdraw",
			testFunc: testMultipleAccountsAndDepositWithdraw,
			want: []account.Account{
				{
					Id:      "001",
					Balance: 862,
				},
				{
					Id:      "002",
					Balance: 190,
				},
				{
					Id:      "003",
					Balance: 230,
				},
				{
					Id:      "004",
					Balance: 657,
				},
				{
					Id:      "005",
					Balance: 335,
				},
			},
		},
		{
			name:     "test3_create_account_and_withdraw",
			testFunc: testCreateAccountAndWithdraw,
			want: []account.Account{
				{
					Id:      "001",
					Balance: 50,
				},
				{
					Id:      "002",
					Balance: 78,
				},
				{
					Id:      "003",
					Balance: 78,
				},
				{
					Id:      "004",
					Balance: 0,
				},
				{
					Id:      "005",
					Balance: 0,
				},
			},
		},
		{
			name:     "test4_create_accounts_and_transfer",
			testFunc: testCreateAccountsAndTransfer,
			want: []account.Account{
				{
					Id:      "001",
					Balance: 384,
				},
				{
					Id:      "002",
					Balance: 814,
				},
				{
					Id:      "003",
					Balance: 196,
				},
				{
					Id:      "004",
					Balance: 755,
				},
				{
					Id:      "005",
					Balance: -82,
				},
			},
		},
		{
			name:     "test5_create_accounts_and_transfer_insufficient_funds",
			testFunc: testTransferInsufficientFunds,
			want: []account.Account{
				{
					Id:      "001",
					Balance: 10,
				},
				{
					Id:      "002",
					Balance: 800,
				},
				{
					Id:      "003",
					Balance: 300,
				},
				{
					Id:      "004",
					Balance: -100,
				},
				{
					Id:      "005",
					Balance: -300,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualAccountList := tt.testFunc()
			if !reflect.DeepEqual(tt.want, actualAccountList) {
				t.Errorf("DepositSystem = %v, want %v", actualAccountList, tt.want)
			}
		})
	}
}
func testCreateAccountAndDeposit() []account.Account {
	d := &depositSystem.DepositSystem{
		AccountList: []account.Account{},
	}
	d.CreateAccount("001")
	d.Deposit("001", 254)
	return d.PrintAllAccountData()
}
func testMultipleAccountsAndDepositWithdraw() []account.Account {
	d := &depositSystem.DepositSystem{
		AccountList: []account.Account{},
	}
	d.CreateAccount("001")
	d.CreateAccount("002")
	d.CreateAccount("003")
	d.CreateAccount("004")
	d.CreateAccount("005")
	d.Deposit("001", 842)
	d.Deposit("002", 220)
	d.WithDraw("001", 400)
	d.Deposit("002", 120)
	d.Deposit("003", 430)
	d.WithDraw("001", 123)
	d.Deposit("001", 543)
	d.Deposit("004", 657)
	d.Deposit("005", 335)
	d.WithDraw("002", 150)
	d.WithDraw("003", 200)
	return d.PrintAllAccountData()
}
func testCreateAccountAndWithdraw() []account.Account {
	d := &depositSystem.DepositSystem{
		AccountList: []account.Account{},
	}
	d.CreateAccount("001")
	d.CreateAccount("002")
	d.CreateAccount("003")
	d.CreateAccount("004")
	d.CreateAccount("005")
	d.Deposit("001", 154)
	d.Deposit("002", 78)
	d.WithDraw("001", 100)
	d.Deposit("003", 78)
	d.WithDraw("001", 4)
	return d.PrintAllAccountData()
}
func testCreateAccountsAndTransfer() []account.Account {
	d := &depositSystem.DepositSystem{
		AccountList: []account.Account{},
	}
	d.CreateAccount("001")
	d.CreateAccount("002")
	d.CreateAccount("003")
	d.CreateAccount("004")
	d.CreateAccount("005")
	d.Deposit("001", 376)
	d.Deposit("002", 694)
	d.Deposit("003", 219)
	d.Transfer("001", "002", 120)
	d.Deposit("004", 732)
	d.Deposit("005", 46)
	d.Transfer("003", "004", 23)
	d.Transfer("005", "001", 128)
	return d.PrintAllAccountData()
}
func testTransferInsufficientFunds() []account.Account {
	d := &depositSystem.DepositSystem{
		AccountList: []account.Account{},
	}
	d.CreateAccount("001")
	d.CreateAccount("002")
	d.CreateAccount("003")
	d.CreateAccount("004")
	d.CreateAccount("005")
	d.Deposit("001", 210)
	d.Transfer("001", "002", 300)
	d.Deposit("003", 500)
	d.Transfer("003", "002", 200)
	d.Transfer("003", "001", 100)
	d.Transfer("004", "002", 300)
	d.Transfer("005", "004", 200)
	d.Transfer("005", "003", 100)
	return d.PrintAllAccountData()
}
