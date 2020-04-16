package pe

import (
	"errors"
	"fmt"
)

//ErrInsufficientFunds implicates that there is no enough moneys to withdraw
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Bitcoin is moneys
type Bitcoin int

//Wallet ...
type Wallet struct {
	balance Bitcoin
}

//Deposit ...
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Stringer ...
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
