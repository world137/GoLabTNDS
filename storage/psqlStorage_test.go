package storage

import (
	"GoLab/account"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestInitPsqlStorage(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "case1 : connect database",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InitPsqlStorage()
			if err != nil {
				t.Errorf("InitPsqlStorage() error = %v", err)
				return
			}
		})
	}
}

func Test_psqlStorage_Create(t *testing.T) {
	type args struct {
		account account.Account
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1: insert account id 001",
			args: args{
				account.Account{
					AccountId: "0011",
					Balance:   100,
					Name:      "test1",
					Email:     "test1@mail.com",
				},
			},
			wantErr: false,
		},
		{
			name: "case 2: insert account id 002",
			args: args{
				account.Account{
					AccountId: "0022",
					Balance:   200,
					Name:      "test2",
					Email:     "test2@mail.com",
				},
			},
			wantErr: false,
		},

		{
			name: "case 2: insert account id 003",
			args: args{
				account.Account{
					AccountId: "0033",
					Balance:   300,
					Name:      "test3",
					Email:     "test3@mail.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			if err := db.Create(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := db.Delete(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_psqlStorage_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		account account.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case 1 : update account 001",
			args: args{
				account: account.Account{
					AccountId: "001",
					Name:      "test1",
					Email:     "test1@mail.com",
					Balance:   1000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			if err := db.Create(account.Account{AccountId: tt.args.account.AccountId}); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := db.Update(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := db.Delete(account.Account{AccountId: tt.args.account.AccountId}); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_psqlStorage_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		account account.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1 : delete account 001",
			args: args{
				account: account.Account{
					AccountId: "001",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			if err := db.Create(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := db.Delete(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_psqlStorage_Read(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		accountId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "case 1 : account id 0011",
			args: args{
				accountId: "0011",
			},
			want: account.Account{
				AccountId: "0011",
				Name:      "test1",
				Email:     "test1@mail.com",
				Balance:   100,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			if err := db.Create(tt.want); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := db.Read(tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("psqlStorage.Read() = %v, want %v", got, tt.want)
			}
			if err := db.Delete(tt.want); (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_psqlStorage_ReadAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []account.Account
		wantErr bool
	}{
		{
			name: "case 1 : account id 0011",
			want: []account.Account{
				{
					AccountId: "0011",
					Name:      "test1",
					Email:     "test1@mail.com",
					Balance:   100,
				},
				{
					AccountId: "0022",
					Name:      "test2",
					Email:     "test2@mail.com",
					Balance:   200,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			for _, value := range tt.want {
				if err := db.Create(value); (err != nil) != tt.wantErr {
					t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			got, err := db.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("psqlStorage.Read() = %v, want %v", got, tt.want)
			}
			for _, value := range tt.want {
				if err := db.Delete(value); (err != nil) != tt.wantErr {
					t.Errorf("psqlStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
