package model

import "time"

type Block struct {
	Index     int64
	Timestamp time.Time
	//Transactions []Transaction
	PrevHash string
	Hash     string
	Nonce    int64
}
