// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.

package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
	close(done)
}

func otherwork(done chan bool) chan bool {
	otc := make(chan bool)
	go func() {

		fmt.Println("I am at otherwork...waiting to worker done.")
		<-done
		fmt.Println("Hurry!!, Worker done. Now will do other work")
		otc <- true
		close(otc)
	}()
	return otc
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)

	ot := otherwork(done)
	fmt.Println("Here, I am doing other things")
	for value := range ot {
		fmt.Println(value)
	}
}
