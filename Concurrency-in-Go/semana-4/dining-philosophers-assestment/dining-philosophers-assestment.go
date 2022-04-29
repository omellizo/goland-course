package main

import (
	"fmt"
	"sync"
)

type Philo struct {
	num int
	timesHasEaten int
}

func (p *Philo) eat(channelsStart chan int, channelsEnd chan int, chopsticksInUse chan int, chopsticksFree chan int) {
	for i := 0; i < 3; i++ {
		
		<- chopsticksInUse
		<- chopsticksInUse
		<- channelsStart

		fmt.Println("starting to eat ", p.num)
		p.timesHasEaten = p.timesHasEaten + 1
		fmt.Println("finishing eating ", p.num)

		channelsEnd <- 1
		chopsticksFree <- 1
		chopsticksFree <- 1

	}
	wg.Done()
}

var wg sync.WaitGroup 

func managementChannels(ch1 chan int, ch2 chan int){
	ch1 <- 1
	<- ch2
}

func host(channelsStart chan int, channelsEnd chan int){
	for {
		for i := 0; i < 2; i++ {
			go managementChannels(channelsStart, channelsEnd)	
		}
	}
}

func chopsticks(chopsticksInUse chan int, chopsticksFree chan int){
	for {
		for i := 0; i < 5; i++ {
			go managementChannels(chopsticksInUse, chopsticksFree)	
		}
	}
}

func main() {
	channelsStart := make(chan int, 2)
	channelsEnd := make(chan int, 2)
	chopsticksInUse := make(chan int, 5)
	chopsticksFree := make(chan int, 5)

	wg.Add(5)
	philos := make([] *Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, 0}
	}

	go host(channelsStart, channelsEnd)
	go chopsticks(chopsticksInUse, chopsticksFree)

	for i := 0; i < 5; i++ {
		go philos[i].eat(channelsStart, channelsEnd, chopsticksInUse, chopsticksFree)
	}
	wg.Wait()
	
	for i := 0; i < 5; i++ {
		fmt.Println("Philosopher", philos[i].num, "has eaten", philos[i].timesHasEaten, "times")
	}
}