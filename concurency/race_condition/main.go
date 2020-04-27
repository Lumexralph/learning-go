package main

import (
	"fmt"

	"github.com/Lumexralph/go_learning/concurency/race_condition/bank"
)

func main() {
	// Person A
	go func() {
		bank.Deposit(200)                 // A1
		fmt.Println("= ", bank.Balance()) // A2
	}()

	// Person B
	go bank.Deposit(300)
}
