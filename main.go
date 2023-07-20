package main

import (
	"GoLab/account"
	"GoLab/business"
	"GoLab/depositSystem"
	"GoLab/storage"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	defer http.ListenAndServe(":3000", r) // while loop
	// storage := make(map[string]account.Account)
	storage, err := storage.InitPsqlStorage()
	if err != nil {
		fmt.Println(err)
		return
	}
	initStandardRoute(r, storage)

	depositSys := &depositSystem.DepositSystem{
		Provider: storage,
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

func initStandardRoute(r *chi.Mux, provider storage.StroageProvider) { // storageMap => interface storage
	// show all accounts
	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		arr, err := provider.ReadAll()
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		res, err := json.Marshal(arr)
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

		resBody, err := provider.Read(accountId)
		if err != nil {
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

		account := account.Account{}
		err = json.Unmarshal(body, &account)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		// accFromRead, err := provider.Read(accountId)
		accFromRead, _ := provider.Read(accountId)

		if accFromRead.AccountId == "" {
			err := provider.Create(account)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			} else {
				err = provider.Update(account)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				}
			}
		}

		err = provider.Update(account)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		account.AccountId = accountId

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
			// storage[account.AccountId] = account
			// accFromRead, err := provider.Read(account.AccountId)
			accFromRead, _ := provider.Read(account.AccountId)

			// if err != nil {
			// 	http.Error(w, err.Error(), 400)
			// 	return
			// }
			if accFromRead.AccountId == "" {
				err := provider.Create(account)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				} else {
					err = provider.Update(account)
					if err != nil {
						http.Error(w, err.Error(), 400)
						return
					}
				}
			}
		}

		w.Write([]byte("success"))
	})

	// update account by accountId
	r.Patch("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "id")
		accFromRead, err := provider.Read(accountId)

		if accFromRead.AccountId == "" {
			http.Error(w, err.Error(), 400)
			return
		}

		// if _, ok := storage[accountId]; !ok {
		// 	http.Error(w, "account not found", 400)
		// 	return
		// }

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		account := account.Account{}
		err = json.Unmarshal(body, &account)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		// storage[accountId] = *account

		// fmt.Println(storage)
		err = provider.Update(account)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.Write([]byte("success"))
	})

	// remove account by accountId
	r.Delete("/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "id")

		// if _, ok := storage[accountId]; !ok {
		// 	http.Error(w, "account not found", 400)
		// 	return
		// }

		// delete(storage, accountId)

		// fmt.Println(storage)
		acc, err := provider.Read(accountId)
		if acc.AccountId == "" {
			http.Error(w, err.Error(), 400)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = provider.Delete(account.Account{AccountId: accountId})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.Write([]byte("success"))
	})
}
