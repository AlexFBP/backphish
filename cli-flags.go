package main

import (
	"flag"
	"fmt"
)

var (
	target string
	times  int
)

const (
	DEFAULT_TIMES = 4
)

func clearFlags() {
	target = ""
	times = DEFAULT_TIMES
}

func parseFlags() {
	clearFlags()

	targetPtr := flag.String("t", "", "Target name")
	timesPtr := flag.Int("n", DEFAULT_TIMES, fmt.Sprintf("Times. Default:%d times", DEFAULT_TIMES))

	flag.Parse()
	target = *targetPtr
	times = *timesPtr
}
