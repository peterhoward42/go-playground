/*
This demonstrates how a default case in a select makes it possible for
attempts to be made to communicate on channels, in the knowledge that the
other party may not be ready, and in that case the attempt will not block.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
getTimeNotifications attempts to provide a (string) alert about the current time
approximately every 5 seconds - depending on GET-in it from an external
http service if it can. It does not expose error conditions. It just provides
the times on the channel when it successfully gets them.
*/
func getTimeNotifications(newTimes chan string) {
	threeSecondTicker := time.NewTicker(5 * time.Second)
	for range threeSecondTicker.C {
		resp, err := http.Get("http://worldtimeapi.org/api/timezone/Europe/London.txt")
		if err != nil {
			continue // Our contract is to ignore errors.
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		newTimes <- string(body)
	}
}

func main() {
	// Main goroutine outputs a message at one second intervals.
	// Inside this loop it attempts to fetch an updated time from a cooperating
	// goroutine, that samples for new times at a slower sampling rate, and
	// furthermore can silently fail to produced new times. At the one
	// second intervals, the main goroutine outputs a message that
	// 1) proves it's still running
	// 2) shows new times when they become known

	newTime := make(chan string)
	go getTimeNotifications(newTime)
	oneSecondTicker := time.NewTicker(time.Second)
	for range oneSecondTicker.C {
		select {
		case timeMessage := <-newTime:
			fmt.Printf("New time arrived: %s\n", timeMessage)
		default:
			fmt.Printf("Still working ...\n")
		}
	}
}
