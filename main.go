package main

import (
	"flag"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	interval := flag.Duration("interval", 30 * time.Second, "Seconds to wait until next movement")
	duration := flag.Duration("duration", 10 * time.Minute, "Total time to run the program")

	flag.Parse()

	end := time.Now().Add(*duration)

	for time.Now().Before(end) {
		robotgo.ScrollMouse(10, "up")
		robotgo.ScrollMouse(10, "down")

		time.Sleep(*interval)
	}

}
