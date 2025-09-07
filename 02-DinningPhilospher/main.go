package main

import (
	"fmt"
	"sync"
	"time"
)

var numPhilosophers = 5

type philospher struct {
	id                  int
	leftFork, rightFork *fork
}

type fork struct {
	sync.Mutex
}

func (p philospher) eat(wg *sync.WaitGroup, host chan struct{}) {
	defer wg.Done()

	// Philospher will eat 3 times
	for range 3 {
		host <- struct{}{}

		p.leftFork.Lock()
		p.rightFork.Lock()

		fmt.Printf("Philosopher %d is eating\n", p.id)
		time.Sleep(time.Second) // eating
		fmt.Printf("Philosopher %d finished eating\n", p.id)

		p.leftFork.Unlock()
		p.rightFork.Unlock()

		<-host

		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup
	forks := make([]*fork, numPhilosophers)

	for i := range numPhilosophers {
		forks[i] = new(fork)
	}

	philosophers := make([]*philospher, numPhilosophers)

	for i := range numPhilosophers {
		philosophers[i] = &philospher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}

	host := make(chan struct{}, numPhilosophers-1)

	for i := range numPhilosophers {
		wg.Add(1)
		go philosophers[i].eat(&wg, host)
	}

	wg.Wait()

	fmt.Println("Dinning Over")
}
