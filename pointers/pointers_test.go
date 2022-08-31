package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got: %s want: %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("got error but didn't want one")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})
	t.Run("withdraw insufficient amount", func(t *testing.T) {
		startingAmount := Bitcoin(20)
		wallet := Wallet{balance: startingAmount}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientBalance)
		assertBalance(t, wallet, startingAmount)
	})
}
