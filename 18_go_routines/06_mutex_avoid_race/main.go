package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var counter int
var mt sync.Mutex

func main() {

	wg.Add(2)
	go incrementor("foo")
	go incrementor("bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 20; i++ {
		mt.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mt.Unlock()
	}

	wg.Done()
}
