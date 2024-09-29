package common

import (
	"fmt"
	"log"
	"math/rand" // "crypto/rand"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	mockServer string
)

func init() {
	rand.Seed(time.Now().UnixMilli())
	SetMockServer("")
}

func CheckURI(testUrl string) (err error) {
	_, err = url.ParseRequestURI(testUrl)
	return
}

func CheckURL(testUrl string) (err error) {
	_, err = url.Parse(testUrl)
	return
}

func SetMockServer(serverUrl string) (err error) {
	if serverUrl != "" {
		if err = CheckURI(serverUrl); err != nil {
			return
		}
	}
	mockServer = serverUrl
	if mockServer != "" && CanLog(LOG_NORMAL) {
		fmt.Printf("Using mockServer: %s\n", mockServer)
	}
	return
}

func ArgsHaveTimes(args ...string) int {
	const DEFAULT = 0
	if len(args) == 1 {
		if v, err := strconv.Atoi(args[0]); err == nil {
			return v
		}
	}
	return DEFAULT
}

type AttemptHander func()
type dummyType struct{}

func AttackRunner(attemptHandle AttemptHander) error {
	q := conf.GetTimes()
	attempts := NewSafeCounter(0)
	maxGoRoutines := conf.threadQty
	activeRoutines := 0

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if CanLog(LOG_NORMAL) {
			log.Printf("\n\nTotal Attempts: %d\n", attempts.Read())
		}
		os.Exit(0)
	}()

	nextAttempt := func(done chan dummyType) {
		attempt := attempts.Add()
		defer func() {
			if CanLog(LOG_NORMAL) {
				fmt.Printf("[Nº %d]end ", attempt)
			}
			done <- struct{}{}
		}()
		if CanLog(LOG_NORMAL) {
			fmt.Printf("[Nº %d]begin ", attempt)
		}
		attemptHandle()
	}

	// Chan for ended routines
	done := make(chan dummyType)

	totalShift := 3 * time.Second
	baseDelay := totalShift / time.Duration(maxGoRoutines)
	jitter := baseDelay / 10

	for {
		for activeRoutines < maxGoRoutines {
			if q > 0 && attempts.Read() >= q {
				return nil
			}
			activeRoutines++
			go nextAttempt(done)
			RandDelayWindowed(baseDelay, jitter) // Randomly wait a little
		}
		<-done
		activeRoutines--
	}
}
