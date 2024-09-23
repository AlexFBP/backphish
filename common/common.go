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

	"github.com/brianvoe/gofakeit"
)

var (
	mockServer string
)

func init() {
	rand.Seed(time.Now().UnixMilli())
	SetMockServer("")
	iniciaIndices()
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
				fmt.Printf("Attempt Nº %d finished\n", attempt)
			}
			done <- struct{}{}
		}()
		if CanLog(LOG_NORMAL) {
			fmt.Printf("Attempt Nº %d - ", attempt)
		}
		attemptHandle()
	}

	// Chan for ended routines
	done := make(chan dummyType)

	for {
		for activeRoutines < maxGoRoutines {
			if q > 0 && attempts.Read() >= q {
				return nil
			}
			activeRoutines++
			go nextAttempt(done)
		}
		<-done
		activeRoutines--
	}
}

func GeneraNIPcolombia() string {
	minOldRange := int(1e3)
	oldRange := int(1e8) - minOldRange
	newRange := int(1e9)
	universe := oldRange + newRange
	id := rand.Intn(universe)
	if id >= oldRange {
		id += newRange - oldRange
	} else {
		id += minOldRange
	}
	return strconv.Itoa(id)
}

// From https://wiki.openstreetmap.org/wiki/ES:Colombia/Gu%C3%ADa_para_mapear/n%C3%BAmeros_telef%C3%B3nicos
var colombiaRangosCelulares = [][]int{
	{300, 2000000, 9399999},
	{301, 2000000, 7999999},
	{302, 2000000, 4699999},
	{303, 2000000, 6849999},
	{304, 2000000, 3899999},
	{324, 2000000, 6999999},
	{305, 7000000, 9599999},
	{310, 2000000, 9999999},
	{311, 2000000, 9999999},
	{312, 2000000, 9999999},
	{313, 2000000, 9999999},
	{314, 2000000, 9999999},
	{320, 2000000, 9999999},
	{321, 2000000, 9999999},
	{322, 2000000, 9999999},
	{323, 2000000, 9999999},
	{315, 2000000, 9999999},
	{316, 2000000, 9999999},
	{317, 2000000, 9999999},
	{318, 2000000, 9999999},
	{319, 2000000, 9999999},
	{350, 2000000, 9399999},
	{351, 2000000, 9399999},
	{302, 4700000, 8699999},
	// {323,6000000,9999999}, // Range already covered
	{324, 1000000, 1999999},
	{324, 7000000, 9999999},
	{333, 0, 499999},
	{333, 6000000, 6999999},
}

var indicesRangosCelulares []int

func iniciaIndices() {
	if len(indicesRangosCelulares) == 0 {
		indicesRangosCelulares = make([]int, len(colombiaRangosCelulares))
		universo := 0 // Numbers that can be generated
		for k, v := range colombiaRangosCelulares {
			numeros := v[2] - v[1] + 1 // max - min + 1
			universo += numeros
			indicesRangosCelulares[k] = universo
		}
	}
}

func generaCelValido(pos int) (cel int, e error) {
	max := indicesRangosCelulares[len(indicesRangosCelulares)-1]
	if pos > max {
		return -1, fmt.Errorf("%d over %d", pos, max)
	} else if pos < 0 {
		return -1, fmt.Errorf("must be >= 0 - got: %d", pos)
	}
	for k, v := range indicesRangosCelulares {
		if pos > v {
			continue
		}
		rule := colombiaRangosCelulares[k]
		cel = rule[0] * int(1e7)
		cel += rule[1]
		if k > 0 {
			cel += pos - indicesRangosCelulares[k-1]
		} else {
			cel += pos
		}
		return cel, nil
	}
	return -1, fmt.Errorf("unhandled value: %d", pos)
}

func GeneraCelColombia() string {
	u := rand.Intn(indicesRangosCelulares[len(indicesRangosCelulares)-1])
	n, _ := generaCelValido(u)
	return strconv.Itoa(n)
}

func GeneraPin(digitos int) string {
	pin := rand.Intn(int(math.Pow10(digitos)))
	return fmt.Sprintf("%0"+strconv.Itoa(digitos)+"d", pin)
}

func GeneraIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Random delay in a range of time Duration. If min > max, values will be considered swapped!
//
// Example:
//
//	RandDelayRange(3*time.Second, 5*time.Second)
//
// is a random delay ranging from 3 to 5 seconds
//
// Previous example is equivalent to:
//
//	RandDelayWindowed(4*time.Second, 2*time.Second)
func RandDelayRange(min, max time.Duration) {
	if min > max {
		min, max = max, min
	}
	time.Sleep(min + time.Duration(rand.Intn(int(max-min))))
}

// Delay within a random window of duration. The 'duration' is the middle/average delay
// of a 'window' of allowed deviations against the specified duration.
// If half of the window (the peak) exceeds the duration, the window will be trimmed
// to twice the duration
//
// Example:
//
//	RandDelayWindowed(10*time.Second, 2*time.Second)
//
// would be equivalent to a delay of 10 seconds, in average.
// The 2 seconds window, will give a range of +/- 1 second
// in which the final delay will be randomly ranging
//
// Previous example, in other words, is equivalent to:
//
//	RandDelayRange(9*time.Second, 11*time.Second)
func RandDelayWindowed(duration, window time.Duration) {
	peak := window / 2
	if peak > duration {
		// Trim peak and fit window
		peak, window = duration, 2*duration
	}
	min := duration - peak
	RandDelayRange(min, min+window)
}

func RandUserName(p *gofakeit.PersonInfo) string {
	u := ""
	switch rand.Intn(4) {
	default:
	case 0:
		u = p.LastName + p.FirstName
	case 1:
		u = p.FirstName + p.LastName
	case 2:
		u = p.FirstName
	case 3:
		u = p.LastName
	}
	return u
}
