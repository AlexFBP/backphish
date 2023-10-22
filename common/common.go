package common

import (
	"fmt"
	"log"
	"math"
	"math/rand" // "crypto/rand"
	"net/url"
	"os"
	"os/signal"
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

func SetMockServer(serverUrl string) error {
	if _, err := url.ParseRequestURI(serverUrl); err != nil && serverUrl != "" {
		return err
	}
	mockServer = serverUrl
	return nil
}

func AttackRunner(attemptHandle func()) error {
	attempts := 0

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println(fmt.Sprintf("\nTotal Attempts: %d", attempts))
		os.Exit(0)
	}()

	for ; ; attempts++ {
		fmt.Print(fmt.Sprintf("\nAttempt Nº %d - ", attempts))
		attemptHandle()
	}
}

func GeneraNIPcolombia() (id int) {
	minOldRange := int(math.Pow10(3))
	oldRange := int(math.Pow10(8)) - minOldRange
	newRange := int(math.Pow10(9))
	universe := oldRange + newRange
	id = rand.Intn(universe)
	if id >= oldRange {
		id += newRange - oldRange
	} else {
		id += minOldRange
	}
	return
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
