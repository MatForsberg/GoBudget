package transactions

import "testing"
import "github.com/MatForsberg/GoBudget/accounts"
import "time"
import "github.com/shopspring/decimal"

func TestCreditPostToAccount(t *testing.T) {
  var f Credit
  f.New("foo", 23.45,time.Date(1,11,1,0,0,0,0,time.Local))
  var a accounts.Account
  a.Balance = decimal.NewFromFloat(100.05)
  f.PostToAccount(&a)
  if !a.Balance.Equals(decimal.NewFromFloat(123.50)) {
    t.Error("Expected 123.50, got ", a.Balance)
  }
}