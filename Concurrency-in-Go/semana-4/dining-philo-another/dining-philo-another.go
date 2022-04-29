package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// put wait group and host channel in global scope
var waitGroup sync.WaitGroup

var hostPermissionChannel = make(chan int, 2)
var doneEatingChannel = make(chan int)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	name           int
	chopstickLeft  *chopstick
	chopstickRight *chopstick
	timesEaten     int
}

// this is a pointer receiver, since the philosopher instance
// keeps track of how many times its eaten in 'timesEaten' field
func (p *philosopher) startEating() {
	maxTimesEaten := 3
	for p.timesEaten < maxTimesEaten {
		// ask permission from host
		hostPermissionChannel <- p.name
		fmt.Printf("Philosopher %d starting to eat (times eaten so far: %d/%d)\n", p.name, p.timesEaten, maxTimesEaten)

		// choose at random which chopstick to use first
		rand.Seed(time.Now().UnixNano())
		dice := rand.Intn(2)

		if dice == 0 {
			p.chopstickLeft.Lock()
			p.chopstickRight.Lock()
		} else {
			p.chopstickRight.Lock()
			p.chopstickLeft.Lock()
		}

		// eat, and increment timesEaten
		p.timesEaten++

		// unlock chopsticks
		p.chopstickLeft.Unlock()
		p.chopstickRight.Unlock()
		doneEatingChannel <- p.name

		fmt.Printf("Philosopher %d finished eating (times eaten so far: %d/%d)\n", p.name, p.timesEaten, maxTimesEaten)
	}
	defer waitGroup.Done()
}

/*
example usage:
$ go run *.go
Philosopher 1 starting to eat (times eaten so far: 0/3)
Philosopher 2 starting to eat (times eaten so far: 0/3)
Philosopher 0 starting to eat (times eaten so far: 0/3)
Philosopher 3 starting to eat (times eaten so far: 0/3)
...
Philosopher 3 starting to eat (times eaten so far: 2/3)
Philosopher 3 finished eating (times eaten so far: 3/3
*/
func main() {
	n := 5

	// create philosophers & chopsticks
	chopsticks := make([]chopstick, n)
	philosophers := make([]philosopher, n)

	for i := 0; i < n; i++ {
		chopsticks[i] = chopstick{}
		philosophers[i] = philosopher{i, &chopsticks[i], &chopsticks[(i+1)%5], 0}
	}

	// start the "host" goroutine
	go (func() {
		for {
			// block until a philosopher sends a "done eating" message
			<-doneEatingChannel
			<-hostPermissionChannel
		}
	})()

	// let the philosophers start eating concurrently
	for i := range philosophers {
		waitGroup.Add(1)
		go philosophers[i].startEating()
	}
	waitGroup.Wait()
}
