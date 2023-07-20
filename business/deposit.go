package business

import (
	"GoLab/depositSystem"
	"encoding/json"
	"io"
	"net/http"
)

type deposit struct {
	AccountId string `json:"account_id"`
	Amount    int    `json:"amount"`
}

func DepositHandler(sys *depositSystem.DepositSystem) func(w http.ResponseWriter, r *http.Request) { // return function
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		bodyjson := &deposit{}
		err = json.Unmarshal(body, bodyjson)
		sys.Deposit(bodyjson.AccountId, bodyjson.Amount) // ต้องการใช้ DepositSystem => รับ parameter

		w.Write([]byte(body))

	}

}
