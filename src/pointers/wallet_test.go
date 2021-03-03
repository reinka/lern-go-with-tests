package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test Wallet Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("Test Wallet Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(35)
		err := wallet.Withdraw(5)
		assertNoError(t, err)
		assertBalance(t, &wallet, Bitcoin(30))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBallance := Bitcoin(15)
		wallet := Wallet{startingBallance}
		err := wallet.Withdraw(Bitcoin(50))

		assertBalance(t, &wallet, startingBallance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

var assertBalance = func(t testing.TB, wallet *Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but did not get one")
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
