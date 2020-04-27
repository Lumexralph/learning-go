package bank3

import (
	"sync"
)

var (
	mu      sync.Mutex // To guard the variable that will be shared
	balance int
)

func Deposit(amount int) {
	mu.Lock() // lock any access to the variable or the process until it is done
	balance += amount
	mu.Unlock() // unlock the access to the variable for access from any goroutine
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()

	return b
}
