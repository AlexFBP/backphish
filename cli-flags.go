package main

import (
	"flag"
	"fmt"
)

var (
	target string
	times  int
	mock   bool
)

const (
	DEFAULT_TIMES = 4
	DEFAULT_MOCK  = "http://localhost:1080"
)

func clearFlags() {
	target = ""
	times = DEFAULT_TIMES
	mock = false
}

func parseFlags() {
	clearFlags()

	targetPtr := flag.String("t", "", "Target name")
	timesPtr := flag.Int("n", DEFAULT_TIMES, fmt.Sprintf("Times. Default:%d times", DEFAULT_TIMES))
	mockPtr := flag.Bool("m", false, "Use mock server in HTTP requests. Default: false")

	flag.Parse()
	target = *targetPtr
	times = *timesPtr
	mock = *mockPtr
}
