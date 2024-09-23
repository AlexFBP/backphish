package main

import (
	"flag"
	"fmt"
)

var (
	target    string
	times     int
	mock      bool
	threadQty int
)

const (
	DEFAULT_TIMES     = 4
	DEFAULT_MOCK      = "http://localhost:1080"
	DEFAULT_PROCESSES = 1
)

func clearFlags() {
	target = ""
	times = DEFAULT_TIMES
	mock = false
	threadQty = DEFAULT_PROCESSES
}

func parseFlags() {
	clearFlags()

	targetPtr := flag.String("a", "", "Attack name")
	timesPtr := flag.Int("n", DEFAULT_TIMES, fmt.Sprintf("Times. Default:%d times", DEFAULT_TIMES))
	mockPtr := flag.Bool("m", false, "Use mock server in HTTP requests. Default: false")
	threadQtyPtr := flag.Int("p", DEFAULT_PROCESSES, fmt.Sprint("Quantity of simultaneous processes. Default:", DEFAULT_PROCESSES))

	flag.Parse()
	target = *targetPtr
	times = *timesPtr
	mock = *mockPtr
	threadQty = *threadQtyPtr
}
