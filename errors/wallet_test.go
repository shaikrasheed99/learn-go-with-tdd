package errors

import (
	"testing"

	"gotest.tools/assert"
)

func TestWallet(t *testing.T) {
	type TestCase struct {
		name           string
		bitcoin        Bitcoin
		depositAmount  Bitcoin
		withdrawAmount Bitcoin
		want           Bitcoin
	}

	testCases := []TestCase{
		{
			name:          "Should be able to deposit bitcoins in wallet",
			bitcoin:       Bitcoin(10),
			depositAmount: Bitcoin(10),
			want:          Bitcoin(10),
		},
		{
			name:           "Should be able to withdraw bitcoins from wallet",
			bitcoin:        Bitcoin(10),
			depositAmount:  Bitcoin(20),
			withdrawAmount: Bitcoin(10),
			want:           Bitcoin(10),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			wallet := Wallet{}

			wallet.Deposit(testCase.depositAmount)
			wallet.Withdraw(testCase.withdrawAmount)
			got := wallet.Balance()

			assert.Equal(t, testCase.want, got)
		})
	}

	t.Run("Should be able to give insufficient balance error when we withdraw bitcoins from wallet", func(t *testing.T) {
		bitcoin := Bitcoin(10)
		wallet := Wallet{}
		want := "insufficient balance in wallet"

		error := wallet.Withdraw(bitcoin)

		if error == nil {
			t.Error()
		}

		assert.Equal(t, want, error.Error())
	})
}
