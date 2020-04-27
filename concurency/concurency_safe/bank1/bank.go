package bank1

var deposits = make(chan int)    // Send amount to deposit
var balances = make(chan int)    // Receive balance

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int    // balance is confined to the teller goroutine

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	// start the monitor goroutine
	go teller()
}