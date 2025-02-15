package common

import (
	"log"
	"sync"
	"time"
)

// Calls a handler repeatedly in a limit time and if handler needs it.
//
// The handler function will be called at most "limit"/"each" times
// while it returns true
func TimedCall(each, limit time.Duration, handler func() (iterateAgain bool)) {
	const FORMAT = time.TimeOnly

	nextAttempt := time.Now()
	timeout := nextAttempt.Add(limit)
	if CanLog(LOG_VERBOSE) {
		log.Printf("start/next: %s - timeout: %s", nextAttempt.Format(FORMAT), timeout.Format(FORMAT))
	}
	for {
		now := time.Now()
		if now.After(timeout) {
			return
		}
		if now.After(nextAttempt) {
			// Call handler
			if !handler() {
				return
			}

			// Calculate next attempt time in terms of duration and ellapsed time
			now = time.Now()
			times := now.Sub(nextAttempt) / each
			nextAttempt = nextAttempt.Add(each * (times + 1))
			if CanLog(LOG_VERBOSE) {
				log.Printf("next: %s", nextAttempt.Format(FORMAT))
			}
		} else {
			time.Sleep(nextAttempt.Sub(now))
		}
	}
}

type SafeCounter struct {
	mu sync.Mutex
	v  int
}

func NewSafeCounter(init int) *SafeCounter {
	return &SafeCounter{v: init}
}

func (c *SafeCounter) Add() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.v++
	return c.v
}

func (c *SafeCounter) Read() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.v
}
