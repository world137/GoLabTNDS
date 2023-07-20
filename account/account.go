package account

type Account struct {
	AccountId string `json:"account_id"`
	Balance   int    `json:"account_balance"`
	Name      string `json:"account_name"`
	Email     string `json:"accoun_email"`
}

// todo encapsulation
func (a Account) GetId() string {
	return a.AccountId
}
func (a *Account) AddBalance(amount int) {
	a.Balance += amount
}
func (a *Account) ReduceBalance(amount int) {
	a.Balance -= amount
}
