package pointers_test

import (
	"github.com/stretchr/testify/assert"
	"learnGoWithTests/pointers"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposite", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Eth(100))

		actual := wallet.Balance()
		expected := pointers.Eth(100)

		assert.Equalf(t, expected, actual, "%s %s", actual, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(100)
		wallet.Withdraw(pointers.Eth(90))

		actual := wallet.Balance()
		expected := pointers.Eth(10)

		assert.Equalf(t, expected, actual, "%s %d", actual, expected)
	})

}
