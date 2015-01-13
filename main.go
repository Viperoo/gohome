package main

import (
	"fmt"
	"time"
)

func main() {

	for t := range time.Tick(1 * time.Second) {
		deadline := time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, 0, t.Location())
		now := time.Now()

		difference := deadline.Unix() - now.Unix()
		hours := (difference / 3600)

		minutes := ((difference / 60) % 60)
		seconds := difference % 60

		fmt.Printf("\x1b[31;1mETA: %[1]v hours %[2]v minutes %[3]v seconds (%[4]v)\n\x1b[0m", hours, minutes, seconds, difference)
	}

}
