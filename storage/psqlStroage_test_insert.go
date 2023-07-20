package storage

// import (
// 	"GoLab/account"
// 	"database/sql"
// 	"testing"

// 	_ "github.com/lib/pq"
// )

// func TestInitPsqlStorage(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		// want    *psqlStorage
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name:    "case1: connect database",
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := InitPsqlStorage()
// 			if err != nil {
// 				t.Errorf("InitPsqlStorage() error = %v", err)
// 				return
// 			}
// 		})
// 	}
// }

// func Test_psqlStorage_Create(t *testing.T) {
// 	type fields struct {
// 		db *sql.DB
// 	}
// 	type args struct {
// 		account account.Account
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "case1: defalut insert",
// 			args: args{
// 				account.Account{
// 					AccountId: "007",
// 					Name:      "world7",
// 					Email:     "world7@gmail.com",
// 					Balance:   10000000,
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "case2: defalut insert",
// 			args: args{
// 				account.Account{
// 					AccountId: "008",
// 					Name:      "world8",
// 					Email:     "world8@gmail.com",
// 					Balance:   80000000,
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "case3: defalut insert",
// 			args: args{
// 				account.Account{
// 					AccountId: "009",
// 					Name:      "world9",
// 					Email:     "world9@gmail.com",
// 					Balance:   90000000,
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			db, err := InitPsqlStorage()
// 			if err != nil {
// 				t.Error("init psqlStorage error", err)
// 			}
// 			if err := db.Create(tt.args.account); (err != nil) != tt.wantErr {
// 				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// // func Test_psqlStorage_Read(t *testing.T) {
// // 	type args struct {
// // 		account account.Account
// // 	}
// // 	tests := []struct {
// // 		name string
// // 		args args
// // 		// want    account.Account
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.

// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			db, err := InitPsqlStorage()
// // 			if err != nil {
// // 				t.Error("init psqlStorage error", err)
// // 			}
// // 			if err := db.Create(tt.args.account); err != nil {
// // 				t.Errorf("psqlStorage.Create() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }
