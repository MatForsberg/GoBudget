package transactions

import "time"
import "github.com/MatForsberg/GoBudget/accounts"
import "github.com/shopspring/decimal"

type Debit Transaction

func (d *Debit) New(name string, amount float64, date time.Time) {
	d.TransactionType = "Debit"
	d.Name = name
	d.Amount = decimal.NewFromFloat(amount)
	d.Date = date
}

func (d Debit) PostToAccount(account *accounts.Account) {
	account.Balance = account.Balance.Sub(d.Amount)
}
