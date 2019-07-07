/*
This explores a "quitable generator" channels pattern idiom suggested in
https://tour.golang.org/concurrency/5

The Tour example is a step-wise fibonacci series generater that doesn't in of
itself know when to stop. Instead the go routine that is calling it stops reading
successive results from the payload channel, and instead sends a quit message
into its quit channel.

I modified it to make sure I've grocked it:
    1) concatenate letter 'A's instead of calculating Fibonacci series.
    2) swap which side is the non-main go routine (arbitrary) - i.e. run
       the generator in a new goroutine rather than the control loop.
    3) make the quit channel bool insead of int
*/

package main

import (
	"fmt"
)

// generator repeatedly sends a string like "AAA" to the results channel,
// where there is one more A in the string each time.
// It continues to do so until it receives (any) boolean on the quit
// channel; at which point it returns.
func generator(results chan string, quit chan bool) {
	s := "A"
	for {
		select {
		case results <- s:
            s = s + "A"
		case <-quit:
			return
		}
	}
}

func main() {
	results := make(chan string)
	quit := make(chan bool)
	go generator(results, quit)
	// Our control loop will fetch 5 results then tell the generator to stop.
	var latestResult string
	for i := 0; i < 5; i++ {
		latestResult = <- results
	}
	quit <- true
	fmt.Printf("Result: %s\n", latestResult)
}
