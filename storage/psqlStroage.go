package storage

import (
	"GoLab/account"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type psqlStorage struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
	dbname   = "training"
)

func InitPsqlStorage() (*psqlStorage, error) {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlconn := []string{
		"host=" + host,
		"port=" + strconv.Itoa(port),
		"user=" + user,
		"password=" + password,
		"dbname=" + dbname,
		"sslmode=disable",
	}
	// open db
	db, err := sql.Open("postgres", strings.Join(psqlconn, " "))
	if err != nil {
		return nil, err
	}
	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!")
	return &psqlStorage{
		db: db,
	}, nil
}
func (m *psqlStorage) Create(account account.Account) error {
	// result, err := m.db.Exec("INSERT INTO account VALUES ('" + account.AccountId + "','" + account.Name + "','" + account.Email + "','" + strconv.Itoa(account.Balance) + "')")
	result, err := m.db.Exec("INSERT INTO account VALUES($1,$2,$3,$4)", account.AccountId, account.Name, account.Email, account.Balance)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
func (m *psqlStorage) Read(accountId string) (account.Account, error) {
	row := m.db.QueryRow("SELECT * FROM account WHERE account_id=$1;", accountId)
	acc := account.Account{}
	err := row.Scan(&acc.AccountId, &acc.Name, &acc.Email, &acc.Balance)
	if err != nil {
		return account.Account{}, err
	}
	return acc, nil
}
func (m *psqlStorage) ReadAll() ([]account.Account, error) {
	var list []account.Account
	rows, err := m.db.Query("SELECT * FROM account")
	if err != nil {
		return []account.Account{}, err
	}
	for rows.Next() {
		acct := account.Account{}
		rows.Scan(&acct.AccountId, &acct.Name, &acct.Email, &acct.Balance)
		if err != nil {
			return []account.Account{}, err
		}
		defer rows.Close()
		list = append(list, acct)
	}
	return list, nil
}
func (m *psqlStorage) Update(account account.Account) error {
	result, err := m.db.Exec("UPDATE account SET balance=$1 WHERE account_id=$2;", account.Balance, account.AccountId)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
func (m *psqlStorage) Delete(account account.Account) error {
	result, err := m.db.Exec("DELETE FROM account WHERE account_id = $1", account.AccountId)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
