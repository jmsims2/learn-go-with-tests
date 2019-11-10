package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Fatal("got %s want %s", got, want)
		}

		if err.Error() != "cannot withdraw, insufficient funds" {
			t.Errorf("got %q, want %q", err.Error(), "cannot withdraw, insufficient funds")
		}
	})

	t.Run("overdraft", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if err.Error() != "cannot withdraw, insufficient funds" {
			t.Errorf("got %q, want %q", err.Error(), "cannot withdraw, insufficient funds")
		}
	})
}
