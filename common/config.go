package common

import (
	"flag"
	"fmt"
)

type Config struct {
	target    string
	times     int
	mock      bool
	threadQty int
}

var conf Config

func init() {
	conf = Config{}
	conf.ParseFlags()
}

func GetConfig() *Config {
	return &conf
}

func (c *Config) GetTarget() string {
	return c.target
}
func (c *Config) GetTimes() int {
	return c.times
}

const (
	DEFAULT_TIMES     = 4
	DEFAULT_MOCK      = "http://localhost:1080"
	DEFAULT_PROCESSES = 1
)

func (c *Config) ClearFlags() {
	c.target = ""
	c.times = DEFAULT_TIMES
	c.mock = false
	c.threadQty = DEFAULT_PROCESSES
}

func (c *Config) ParseFlags() {
	c.ClearFlags()

	targetPtr := flag.String("a", "", "Attack name")
	timesPtr := flag.Int("n", DEFAULT_TIMES, fmt.Sprintf("Times. Default:%d times", DEFAULT_TIMES))
	mockPtr := flag.Bool("m", false, "Use mock server in HTTP requests. Default: false")
	threadQtyPtr := flag.Int("p", DEFAULT_PROCESSES, fmt.Sprint("Quantity of simultaneous processes. Default:", DEFAULT_PROCESSES))

	flag.Parse()
	c.target = *targetPtr
	c.times = *timesPtr
	c.mock = *mockPtr
	c.threadQty = *threadQtyPtr

	if c.mock {
		SetMockServer(DEFAULT_MOCK)
	}
}
