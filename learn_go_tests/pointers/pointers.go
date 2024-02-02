package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	newBalance := w.balance - amount

	if newBalance < 0 {
		return errors.New("withdrawal amount exceeded")
	} else {
		w.balance = newBalance
		return nil
	}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// here we are defining a method on a type which is similar to as in a struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
