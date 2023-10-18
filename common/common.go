package common

import (
	"fmt"
	"log"
	"math"
	"math/rand" // "crypto/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func AttackRunner(attemptHandle func()) error {
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

func GeneraNIPcolombia() int {
	return rand.Intn(100000000) + 1000000000
}

func GeneraPin(digitos int) int {
	return rand.Intn(int(math.Pow10(digitos)))
}

func GeneraIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandDelay(minSeconds, maxSeconds int) {
	time.Sleep(time.Second * time.Duration(minSeconds+rand.Intn(maxSeconds-minSeconds)))
}
