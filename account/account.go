package account

type Account struct {
	Id      string `json:"account_id"`
	Balance int    `json:"account_balance"`
	Name    string `json:"account_name"`
	Email   string `json:"accoun_email"`
}

// todo encapsulation
