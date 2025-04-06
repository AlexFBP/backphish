package common

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	target    string
	times     int
	threadQty int
	logLevel  int
}

var conf Config

func init() {
	conf = Config{}
	if !isTestMode() { // Evita procesar banderas si estamos en modo de prueba
		conf.ParseFlags()
	}
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
func (c *Config) GetThreads() int {
	return c.threadQty
}

const (
	DEFAULT_TIMES     = 4
	DEFAULT_MOCK      = "http://localhost:1080"
	DEFAULT_PROCESSES = 1
	DEFAULT_LOGLEVEL  = LOG_NORMAL
)

func (c *Config) ClearFlags() {
	c.target = ""
	c.times = DEFAULT_TIMES
	c.threadQty = DEFAULT_PROCESSES
}

func (c *Config) ParseFlags() {
	c.ClearFlags()

	targetPtr := flag.String("n", "", "Attack name")
	timesPtr := flag.Int("q", DEFAULT_TIMES, "Total quantity of attacks. Set 0 for unlimited")
	flag.Func("m", "Use mock server in HTTP requests. Must be a valid URL or empty string \"\" to take default. (\""+DEFAULT_MOCK+"\")", SetMockServer)
	threadQtyPtr := flag.Int("p", DEFAULT_PROCESSES, "Simultaneous processes/threads")
	logLevelPtr := flag.Int("l", DEFAULT_LOGLEVEL, fmt.Sprintf("Log Level. Values:%s", LogLevelsHelp()))

	flag.Parse()
	c.target = *targetPtr
	c.times = *timesPtr
	c.threadQty = *threadQtyPtr
	c.logLevel = *logLevelPtr
}

// isTestMode checks if the program is running in test mode
func isTestMode() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") {
			return true
		}
	}
	return false
}
