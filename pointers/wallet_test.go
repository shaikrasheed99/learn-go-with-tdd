package pointers

import (
	"testing"

	"gotest.tools/assert"
)

func TestWallet(t *testing.T) {
	t.Run("Should be able to deposit bitcoin into wallet", func(t *testing.T) {
		bitcoin := Bitcoin(10)
		wallet := Wallet{}
		want := Bitcoin(10)

		wallet.Deposit(bitcoin)
		got := wallet.Balance()

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to withdraw bitcoin into wallet", func(t *testing.T) {
		bitcoin := Bitcoin(10)
		wallet := Wallet{}
		want := Bitcoin(10)

		wallet.Deposit(bitcoin + bitcoin)
		wallet.Withdraw(bitcoin)
		got := wallet.Balance()

		assert.Equal(t, want, got)
	})
}
