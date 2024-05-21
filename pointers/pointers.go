package pointers

import (
	"errors"
	"fmt"
)

type Eth int
type Wallet struct {
	balance Eth
}

func (e Eth) String() string {
	return fmt.Sprintf("%d ETH", e)
}

func (w *Wallet) Deposit(amount Eth) {
	w.balance += amount
}

func (w *Wallet) Balance() Eth {
	return w.balance
}

func (w *Wallet) Withdraw(amount Eth) error {
	if w.balance < amount {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}
