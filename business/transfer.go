package business

import (
	"GoLab/depositSystem"
	"encoding/json"
	"io"
	"net/http"
)

type transfer struct {
	FromAccountId string `json:"from_account_id"`
	ToAccountId   string `json:"to_account_id"`
	Amount        int    `json:"amount"`
}

func TransferHandler(sys *depositSystem.DepositSystem) func(w http.ResponseWriter, r *http.Request) { // return function
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		bodyjson := &transfer{}
		err = json.Unmarshal(body, bodyjson)
		sys.Transfer(bodyjson.FromAccountId, bodyjson.ToAccountId, bodyjson.Amount) // ต้องการใช้ DepositSystem => รับ parameter

		w.Write([]byte(body))
	}

}
