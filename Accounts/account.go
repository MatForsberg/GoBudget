package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/shopspring/decimal"
)

type Account struct {
	Name    string
	Balance decimal.Decimal
	Enabled bool
}

func (a Account) Create() (e error) {
	dat, e := ioutil.ReadFile("Accounts")

	var accounts []Account
	e = json.Unmarshal(dat, &accounts)

	accounts = append(accounts, a)
	as, e := json.Marshal(accounts)
	e = ioutil.WriteFile("Accounts", as, os.ModePerm)
	return e
}

func GetAll() (accounts []Account) {
	dat, _ := ioutil.ReadFile("Accounts")

	json.Unmarshal(dat, &accounts)
	return accounts
}

func (a *Account) Print() {
	var enabled string
	if a.Enabled {
		enabled = "Enabled"
	} else {
		enabled = "Disabled"
	}
	fmt.Printf("%-15.10s $%s %s\n", a.Name, a.Balance.StringFixed(2), enabled)
}
