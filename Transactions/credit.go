package transactions

import "time"
import "github.com/MatForsberg/GoBudget/accounts"
import "github.com/shopspring/decimal"

type Credit Transaction

func (c *Credit) New(name string, amount float64, date time.Time){
	c.TransactionType = "Credit"
	c.Name = name
	c.Amount = decimal.NewFromFloat(amount)
	c.Date = date
	c.posted = false
}

func (c Credit) PostToAccount(account *accounts.Account ){
	account.Balance = account.Balance.Add(c.Amount)
}