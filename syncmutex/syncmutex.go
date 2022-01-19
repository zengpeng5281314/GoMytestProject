package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func set(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = i
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

//sync.Mutex 互斥锁
func main() {
	numGR := 5
	var wg sync.WaitGroup
	fmt.Printf("%d", v1)
	for i := 1; i <= numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			set(i)
			fmt.Printf("----%d", read())
		}(i)
	}
	wg.Wait()
}
