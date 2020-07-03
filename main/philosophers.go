package main

import (
	"fmt"
	"sync"
)

type Chopstick struct { sync.Mutex }

type Philosopher struct {
	number int
	leftCS, rightCS * Chopstick
	channel * chan bool
	wg * sync.WaitGroup
}

type Host struct {
	philosophers []bool
}

func (philosopher Philosopher) Eat() {
	for i := 0; i < 3; i++ {
		var answer bool
		for {
			*philosopher.channel <- true
			answer = <- *philosopher.channel
			if answer == true {
				break
			}
		}
		fmt.Println("Broke!")
		philosopher.leftCS.Lock()
		philosopher.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", philosopher.number)
		philosopher.wg.Done()
		fmt.Printf("finishing eating %d\n", philosopher.number)
		philosopher.leftCS.Lock()
		philosopher.rightCS.Unlock()
		*philosopher.channel <- false
		<- *philosopher.channel
	}
}

func(host Host) processRequest(request bool, philosopher int) {
	//fmt.Println("Processing request ", request, " of philosopher ", philosopher)
	if request == false {
		host.philosophers[philosopher] = false
		channels[philosopher] <- true
	} else if !host.philosophers[(philosopher + 1) % 5] && !host.philosophers[(philosopher - 1 + 5) % 5] {
		host.philosophers[philosopher] = true
		channels[philosopher] <- true
	} else {
		channels[philosopher] <- false
	}
}

func (host Host) managePhilosophers() {
	for {
		select {
		case request := <- channels[0]: host.processRequest(request, 0)
		case request := <- channels[1]: host.processRequest(request, 1)
		case request := <- channels[2]: host.processRequest(request, 2)
		case request := <- channels[3]: host.processRequest(request, 3)
		case request := <- channels[4]: host.processRequest(request, 4)
		}
	}

}

var channels [5]chan bool

func main() {
	for i := 0; i < 5; i++ {
		channels[i] = make(chan bool)
	}
	chopsticks := make([] * Chopstick, 5)
	philosophers := make([] * Philosopher, 5)
	var wg sync.WaitGroup
	wg.Add(15)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			i + 1, chopsticks[i],
			chopsticks[(i + 1) % 5],&(channels[i]),&wg}
	}
	host := Host{make([]bool, 5)}
	go host.managePhilosophers()
	for i := 0; i < 5; i++ {
		go philosophers[i].Eat()
	}
	wg.Wait()
}