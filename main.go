package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	deadline := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())

	fmt.Print(now)
	fmt.Print(deadline)
	fmt.Print("Go home, u are drunk.")
}
