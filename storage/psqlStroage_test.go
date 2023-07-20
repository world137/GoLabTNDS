package storage

import (
	"GoLab/account"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

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
		// TODO: Add test cases.
		{
			name:    "case1: defalut read",
			want:    []account.Account{{AccountId: "002", Name: "world2", Email: "world2@gmail.com", Balance: 200}, {AccountId: "001", Name: "world", Email: "world@gmail.com", Balance: 100000}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitPsqlStorage()
			if err != nil {
				t.Error("init psqlStorage error", err)
			}
			got, err := db.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("psqlStorage.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("psqlStorage.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
