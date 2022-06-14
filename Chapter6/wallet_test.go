package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		// Creating an empty wallet
		wallet := Wallet{}

		// Depositing money into the wallet
		wallet.Deposit(Bitcoin(10))

		// Asserting the balance
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		// Creating a wallet with 20 Bitcoins
		wallet := Wallet{startingBalance}

		// Withdrawing the money from the wallet
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		// Asserting the balance
		assertBalance(t, wallet, startingBalance)

	})

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
