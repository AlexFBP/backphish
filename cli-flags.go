package main

import (
	"flag"
)

var (
	target string
	times  int
)

const (
	DEFAULT_TIMES = 1
)

func clearFlags() {
	target = ""
	times = DEFAULT_TIMES
}

func parseFlags() {
	clearFlags()

	targetPtr := flag.String("t", "", "Target name")
	timesPtr := flag.Int("n", DEFAULT_TIMES, "Times. Default")

	flag.Parse()
	target = *targetPtr
	times = *timesPtr
}
