package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	balance       int
	transactionNo int
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	balance = 1000
	transactionNo = 0

	transChan := make(chan bool)
	balanceChan := make(chan int)

	fmt.Printf("Starting Balance: %d\n", balance)

	wg.Add(1)

	for i := range 100 {
		go func(ii int) {
			transactionAmount := rand.Intn(25)
			balanceChan <- transactionAmount

			if ii == 99 {
				fmt.Println("Should be quit time")
				transChan <- true
				wg.Done()
			}
		}(i)
	}

	breakPoint := false
	for {
		if breakPoint {
			break
		}

		select {
		case amt := <-balanceChan:
			fmt.Println("Transaction for $", amt)

			if (balance - amt) < 0 {
				fmt.Println("Transaction failed!")
			} else {
				balance = balance - amt
				fmt.Println("Transaction succeeded")
			}
			fmt.Println("Balance now $", balance)

		case status := <-transChan:
			if status {
				fmt.Println("Done")
				breakPoint = true
				close(transChan)
			}
		}
	}

	fmt.Printf("Final Balance: %d\n", balance)
}

func transaction(amount int) bool {
	approved := false

	if (balance - amount) < 0 {
		approved = false
	} else {
		approved = true
		balance = balance - amount
	}

	approveText := "declines"

	if approved {
		approveText = "approved"
	}

	transactionNo = transactionNo + 1

	fmt.Println(transactionNo, "Transaction for $", amount, approveText)
	fmt.Println("\tRemaining balance $", balance)
	return approved
}
