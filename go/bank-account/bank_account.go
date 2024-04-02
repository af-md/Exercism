package account

import "sync"

// Define the Account type here.

type Account struct {
	balance       int64
	accountClosed bool
	lock          sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
	}
}

func (a *Account) Balance() (int64, bool) {
	if a.accountClosed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {

	// lock the function to prevent race conditions
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.accountClosed {
		return 0, false
	}

	if a.balance == 0 && amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
   
    // lock the function to prevent race conditions
    a.lock.Lock()
	defer a.lock.Unlock()
	
    if a.accountClosed {
		return 0, false
	}

	var currentBalance = a.balance
	a.balance = 0
	a.accountClosed = true
	return currentBalance, a.accountClosed
}
