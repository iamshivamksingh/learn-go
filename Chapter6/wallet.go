package main

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

// The syntax for creating a method on a type declaration
// is the same as it is on a struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// pointer to wallet struct
func (w *Wallet) Balance() Bitcoin {
	// implicit dereference
	// automatically dereferenced
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
