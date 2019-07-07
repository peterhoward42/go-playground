package main

import (
        "fmt"
        "time"
        )

func main() {
    // launch 20 goroutines
    doneChan := make(chan int)
    for i := 0; i < 20; i++ {
        go myFn(doneChan, i)
    }
    // stay alive until we have received 20 completion signals back
    for doneSignals := 0; doneSignals < 20; {
        idxFinished := <- doneChan
        doneSignals++
        fmt.Printf("returned from myFn for idx %d\n", idxFinished)
    }
}

func myFn(doneChan chan(int), count int) {
    // block for a time inversely proportional to count, so that last one
    // likely finishes first
    delay := time.Duration(1 + (20 - count)) * time.Second
    time.Sleep(delay)
    doneChan <- count
}
