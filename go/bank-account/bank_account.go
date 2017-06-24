package account

import (
	"sync"
)

const testVersion = 1

// Account : For Bank Account
type Account struct {
	balance int64
	close   bool
	mux     sync.Mutex
}

// Open : Open new account
func Open(initialDeposit int64) *Account {

	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close : Close account
func (acc *Account) Close() (payout int64, ok bool) {

	acc.mux.Lock()
	defer acc.mux.Unlock()

	// if acc already close
	if acc.close {
		return 0, false
	}
	acc.close = true
	return acc.balance, true
}

// Balance : get Balance
func (acc *Account) Balance() (balance int64, ok bool) {

	acc.mux.Lock()
	defer acc.mux.Unlock()
	if acc.close {
		return 0, false
	}
	return acc.balance, true
}

// Deposit :
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {

	acc.mux.Lock()
	defer acc.mux.Unlock()
	if acc.close {
		return 0, false
	}

	if acc.balance+amount < 0 {
		return acc.balance, false
	}
	acc.balance += amount

	return acc.balance, true
}
