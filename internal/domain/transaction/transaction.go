package transaction

type Transaction struct {
	ID     string
	From   string
	To     string
	Amount float64
	Time   int64
}
