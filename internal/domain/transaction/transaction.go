package transaction

type Transaction struct {
	ID     string  `json:"id"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Time   int64   `json:"time"`
}

func NewTransaction(id, from, to string, amount float64, timestamp int64) *Transaction {
	return &Transaction{
		ID:     id,
		From:   from,
		To:     to,
		Amount: amount,
		Time:   timestamp,
	}
}
