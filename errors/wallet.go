package errors

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientBalance = errors.New("insufficient balance in wallet")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientBalance
	}

	w.balance -= amount

	return nil
}
