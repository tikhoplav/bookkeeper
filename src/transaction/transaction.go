package transaction

import (
	"time"
)

type Transaction struct {
	ID       uint64    `json:"id"`
	DateTime time.Time `json:"datetime"`
	DebitID  uint64    `json:"debit_id"`
	CreditID uint64    `json:"credit_id"`
	Amount   float64   `json:"amount"`
	Note     string    `json:"note"`
}
