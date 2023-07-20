package main

import (
	"GoLab/account"
	"GoLab/business"
	"GoLab/depositSystem"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type StroageProvider interface {
	Create(account.Account) error
	Read(string) account.Account
	ReadAll() []account.Account
	Update(account.Account) error
	Delete(account.Account) error
}

type mapStroage struct {
	data map[string]account.Account
}

func (m *mapStroage) Create(account account.Account) error {
	if account.AccountId == "" {
		return fmt.Errorf("No account id")
	}
	m.data[account.AccountId] = account
	return nil
}
func (m *mapStroage) Read(accountId string) account.Account {
	if accountId == "" {
	}
	return m.data[accountId]
}
func (m *mapStroage) ReadAll() []account.Account {
	var returnArray []account.Account
	for _, v := range m.data {
		returnArray = append(returnArray, v)
	}

	sort.SliceStable(returnArray, func(i, j int) bool {
		return returnArray[i].AccountId < returnArray[j].AccountId //condition
	})

	return returnArray
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

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	defer http.ListenAndServe(":3000", r) // while loop
	storage := make(map[string]account.Account)

	initStandardRoute(r, storage)

	depositSys := &depositSystem.DepositSystem{
		AccountMap: storage,
	}

	initRoute(r, depositSys)

}

func initRoute(r *chi.Mux, depositSys *depositSystem.DepositSystem) {
	// { "account_id": "001", "amout":100 }
	r.Post("/transactions/deposit", business.DepositHandler(depositSys))
	// { "account_id": "001", "amout":100 }
	r.Post("/transactions/withdraw", business.WithDrawHandler(depositSys))
	// { "from_account_id": "001", "to_account_id": "001", "amout":100 }
	r.Post("/transactions/transfer", business.TransferHandler(depositSys))
}

func initStandardRoute(r *chi.Mux, storage map[string]account.Account) {
	// show all accounts
	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		var resArr []account.Account

		for _, val := range storage {
			resArr = append(resArr, val)
		}

		res, err := json.Marshal(resArr)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		w.Write(res)
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(res)
	})

	// show only accountId
	r.Get("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "id")

		resBody, ok := storage[accountId]
		if !ok {
			http.Error(w, "account not found", 400)
			return
		}

		res, err := json.Marshal(resBody)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.Write(res)
	})

	// insert or update account by id
	r.Put("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "id")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		account := &account.Account{}
		err = json.Unmarshal(body, account)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		storage[accountId] = *account

		fmt.Println(storage)

		w.Write([]byte("success"))
	})

	// insert or update all accouts
	r.Put("/accounts", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		accounts := []account.Account{}
		err = json.Unmarshal(body, &accounts)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		for _, account := range accounts {
			storage[account.AccountId] = account
		}

		fmt.Println(storage)

		w.Write([]byte("success"))
	})

	// update account by accountId
	r.Patch("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "id")

		if _, ok := storage[accountId]; !ok {
			http.Error(w, "account not found", 400)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		account := &account.Account{}
		err = json.Unmarshal(body, account)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		storage[accountId] = *account

		fmt.Println(storage)

		w.Write([]byte("success"))
	})

	// remove account by accountId
	r.Delete("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "id")

		if _, ok := storage[accountId]; !ok {
			http.Error(w, "account not found", 400)
			return
		}

		delete(storage, accountId)

		fmt.Println(storage)

		w.Write([]byte("success"))
	})
}
