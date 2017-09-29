package transactions

import "testing"
import "github.com/MatForsberg/GoBudget/accounts"
import "time"
import "github.com/shopspring/decimal"

func TestDebitPostToAccount(t *testing.T) {
  var f Debit
  f.New("foo", 23.45,time.Now())
  var a accounts.Account
  a.Balance = decimal.NewFromFloat(100.05)
  f.PostToAccount(&a)
  if !a.Balance.Equals(decimal.NewFromFloat(76.60)){
    t.Error("Expected 76.60, got ", a.Balance)
  }
}