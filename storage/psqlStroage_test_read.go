package storage

import (
	"GoLab/account"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

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
		// TODO: Add test cases.
		{
			name: "case1: defalut read",
			args: args{
				accountId: "002",
			},
			want:    account.Account{AccountId: "002", Name: "world2", Email: "world2@gmail.com", Balance: 200},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			got, err := db.Read(tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("psqlStorage.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
