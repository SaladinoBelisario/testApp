package main

import (
	"fmt"
	"sync"
	"time"
)

type chopS struct{ sync.Mutex }

type philosopher struct {
	id              int
	leftCS, rightCS *chopS
}

var wGroup sync.WaitGroup

func main() {

	count := 5

	sticks := make([]*chopS, count)
	for i := 0; i < count; i++ {
		sticks[i] = new(chopS)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{id: i, leftCS: sticks[i], rightCS: sticks[(i+1)%count]}
		wGroup.Add(1)
		go philosophers[i].eat()
	}

	wGroup.Wait()
}

func (p philosopher) eat() {

	for j := 0; j < 3; j++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher %d is eating\n", p.id+1)
		time.Sleep(time.Second)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Printf("Philosopher %d has finished eating\n", p.id+1)
		time.Sleep(time.Second)
	}

	wGroup.Done()
}
