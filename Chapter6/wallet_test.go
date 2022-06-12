package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		// Creating an empty wallet
		wallet := Wallet{}

		// Depositing money into the wallet
		wallet.Deposit(Bitcoin(10))

		// Asserting the balance
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		// Creating a wallet with 20 Bitcoins
		wallet := Wallet{balance: Bitcoin(20)}

		// Withdrawing the money from the wallet
		wallet.Withdraw(Bitcoin(10))

		// Asserting the balance
		assertBalance(t, wallet, Bitcoin(10))
	})

}
