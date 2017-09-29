package transactions

import "time"
import "github.com/MatForsberg/GoBudget/accounts"
import "github.com/shopspring/decimal"

type Action interface{
	New(name string, amount decimal.Decimal, date time.Time)
	PostToAccount(account *accounts.Account )
}