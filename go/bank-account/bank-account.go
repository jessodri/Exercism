package account

import "sync"

// Account struct
type Account struct {
	open    bool
	balance int64
	mux     sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{open: true, balance: initialDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open != true {
		return 0, false
	}
	a.open = false
	payout = a.balance
	a.balance = 0
	return payout, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open != true {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.open {
		return 0, false
	}
	if amount < 0 {
		if (a.balance + amount) < 0 {
			return a.balance, false
		}
		a.balance += amount
		return a.balance, true
	}
	if amount < 0 {
		a.balance += amount
		return a.balance, true
	}
	a.balance += amount
	return a.balance, true
}
