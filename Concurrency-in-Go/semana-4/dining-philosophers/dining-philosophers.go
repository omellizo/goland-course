package main

import (
	"fmt"
	"sync"
)

type CHopS struct {
	sync.Mutex
}

type Philo struct {
	leffCS, rightCS *CHopS
}

func (p Philo) eat() {
	for {
		p.leffCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leffCS.Unlock()
	}
}

func main() {
	CSticks := make([]*CHopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(CHopS)
	}

	philos := make([] *Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i + 1) % 5]}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	
}