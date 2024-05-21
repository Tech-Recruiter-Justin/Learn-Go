package pointers_test

import (
	"github.com/stretchr/testify/assert"
	"learnGoWithTests/pointers"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}

		wallet.Deposit(10)

		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := pointers.Wallet{}

		wallet.Deposit(100)
		err := wallet.Withdraw(90)

		assert.Nil(t, err, "expected no error but got %s", err)
		assertBalance(t, wallet, pointers.Eth(10))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := pointers.Wallet{}
		var startingBalance = pointers.Eth(15)

		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(100)

		expectedErrorMsg := "insufficient funds"
		assert.EqualErrorf(t, err, expectedErrorMsg, "expected an error with msg '%s' but got '%s' instead", expectedErrorMsg, err)
	})
}

func assertBalance(tb testing.TB, wallet pointers.Wallet, expectedBalance pointers.Eth) {
	tb.Helper()
	assert.Equalf(tb, expectedBalance, wallet.Balance(), "expected: %s, actual: %s", expectedBalance, wallet.Balance())
}
