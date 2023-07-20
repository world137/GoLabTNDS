package storage

// import (
// 	"GoLab/account"
// 	"database/sql"
// 	"testing"

// 	_ "github.com/lib/pq"
// )

// func Test_psqlStorage_Delete(t *testing.T) {
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
// 			name: "case1: defalut delete",
// 			args: args{
// 				account.Account{
// 					AccountId: "008",
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
// 			if err := db.Delete(tt.args.account); (err != nil) != tt.wantErr {
// 				t.Errorf("psqlStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
