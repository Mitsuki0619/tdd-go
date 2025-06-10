package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("入金", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("出金", func(t *testing.T) {
		wallet := Wallet{
			balance: BitCoin(20),
		}
		err := wallet.Withdraw(10)

		assertBalance(t, wallet, 10)
		assertNoError(t, err)
	})

	t.Run("予算額を超える出金はエラーとなる", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{
			balance: startingBalance,
		}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrINsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
