package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer: fires once after the specified duration.
	timer := time.NewTimer(2 * time.Second)
	<-timer.C // blocks until timer fires
	fmt.Println("Timer fired")

	// Ticker: fires repeatedly at intervals.
	ticker := time.NewTicker(500 * time.Millisecond)

	// Run a goroutine to handle ticks.
	go func() {
		// range over ticker.C would run forever; we stop it externally.
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Let the ticker run for a while, then stop it.
	time.Sleep(2 * time.Second)
	ticker.Stop() // stops the ticker; the goroutine will exit because range stops
	fmt.Println("Ticker stopped")
}
