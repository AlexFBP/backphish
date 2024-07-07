package common

import (
	"fmt"
	"log"
	"math"
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

func SetMockServer(serverUrl string) error {
	if _, err := url.ParseRequestURI(serverUrl); err != nil && serverUrl != "" {
		return err
	}
	mockServer = serverUrl
	return nil
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

func AttackRunner(attemptHandle func(), q int) error {
	attempts := 1

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("\n\nTotal Attempts: %d\n", attempts)
		os.Exit(0)
	}()

	for ; ; attempts++ {
		fmt.Printf("\nAttempt Nº %d - ", attempts)
		attemptHandle()

		if q > 0 && attempts >= q {
			break
		}
	}
	return nil
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
