package transactions

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionType string
	Name            string
	Amount          decimal.Decimal
	Date            time.Time
}

func (t *Transaction) Print() {
	fmt.Printf("%s.1 %s %d", t.TransactionType, t.Name, t.Amount)
}
