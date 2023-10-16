package common

import (
	"log"
	"math/rand" // "crypto/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func AttackRunner(attemptHandle func()) error {
	rand.Seed(time.Now().UnixMilli())
	attempts := 0

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("Total Attempts: ")
		log.Println(attempts)
		os.Exit(0)
	}()

	for ; ; attempts++ {
		attemptHandle()
	}
}
