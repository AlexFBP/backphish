package common

import (
	"fmt"
)

const (
	LOG_VERBOSE = iota // Verbose and/or debug
	LOG_NORMAL         // Moderate
	LOG_QUIET          // Severe only
)

var logLevels map[int]string = map[int]string{
	LOG_VERBOSE: "Verbose/Debug",
	LOG_NORMAL:  "Normal",
	LOG_QUIET:   "Quiet/Severe",
}

func CanLog(messageLevel int) bool {
	_, ok := logLevels[messageLevel]
	return ok && (messageLevel >= conf.logLevel)
}

func LogLevelsHelp() (s string) {
	s = ""
	for k, v := range logLevels {
		s += fmt.Sprintf(" %d-%s", k, v)
	}
	return
}
