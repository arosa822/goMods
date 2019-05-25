package main

import (
	"fmt"
	"sync"
)

var x = 0
var y = 0

// introduces race condition when function tries to acces the value of x
// concurrently
func increment_race(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

// solves the issue of race condition with the use of mutex
func increment_noRace(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var w2 sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		w2.Add(1)
		go increment_race(&w)
		go increment_noRace(&w2, &m)
	}
	w.Wait()
	w2.Wait()
	fmt.Println("final value of x", x)

	fmt.Println("final value of y", y)
	fmt.Println("vim-go")
}
