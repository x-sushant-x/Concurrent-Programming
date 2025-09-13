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

	fmt.Printf("Starting Balance: %d\n", balance)

	wg.Add(1)

	for i := range 100 {
		go func(ii int, transChan chan bool) {
			transactionAmount := rand.Intn(25)
			transaction(transactionAmount)
			if ii == 99 {
				transChan <- true
			}
		}(i, transChan)
	}

	if <-transChan {
		wg.Done()
	}

	close(transChan)
	fmt.Printf("Final Balancer: %d\n", balance)
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
