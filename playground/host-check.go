package playground

import (
	"fmt"

	"github.com/AlexFBP/backphish/common"
)

const (
	// Host states
	HostUp = iota
	HostDown
	HostUnknown
)

type TargetDict map[string](map[string]bool)
type StorageDict map[string]([]string)

func checkHosts() {
	allTargets := TargetDict{}
	// Read host-up and host-down yml files
	file_up := "playground/hostup.yml"
	alive := StorageDict{}
	err := common.ReadYamlFile(file_up, &alive)
	if err != nil {
		panic(err)
	}
	file_down := "playground/hostdown.yml"
	down := StorageDict{}
	err = common.ReadYamlFile(file_down, &down)
	if err != nil {
		panic(err)
	}
	mergeDicts(alive, allTargets, true)
	mergeDicts(down, allTargets, false)

	// Test all hosts
	for target, scams := range allTargets {
		for scam, state := range scams {
			switch checkHostState(scam) {
			case HostUp:
				if !state { // if it was down
					allTargets[target][scam] = true
					if common.CanLog(common.LOG_NORMAL) {
						fmt.Printf("Now it seems that '%s' is up\n", scam)
					}
				}
			case HostDown:
				if state { // if it was up
					allTargets[target][scam] = false
					if common.CanLog(common.LOG_NORMAL) {
						fmt.Printf("Now it seems that '%s' is down\n", scam)
					}
				}
			case HostUnknown:
				// warn
				if common.CanLog(common.LOG_NORMAL) {
					fmt.Printf("Unknown state for host '%s'\n", scam)
				}
			}
		}
	}

	// Update the host-up and host-down files
	alive, down = splitDicts(allTargets)
	err = common.WriteYamlFile(file_up, alive)
	if err != nil {
		panic(err)
	}
	err = common.WriteYamlFile(file_down, down)
	if err != nil {
		panic(err)
	}
	// Print the results
}

func checkHostState(url string) int {
	err := common.CheckURL(url)
	if err != nil {
		return HostUnknown
	}
	h := common.ReqHandler{}
	var step uint8
	err, step = h.SendGet(url, nil, nil, nil)
	if err != nil {
		if common.CanLog(common.LOG_VERBOSE) {
			fmt.Printf("Host %s seems to be down - lastStep: %d - err: %s\n", url, step, err)
		}
		return HostDown
	}
	if h.Response.StatusCode >= 100 && h.Response.StatusCode < 300 {
		if common.CanLog(common.LOG_VERBOSE) {
			fmt.Printf("Host %s seems to be up - lastStep: %d\n", url, step)
		}
		return HostUp
	}
	return HostUnknown
}

// mergeDicts merges the source dictionary into the target dictionary, with a preset state for each entry.
func mergeDicts(source StorageDict, target TargetDict, state bool) {
	for rawTarget, rawMirrors := range source {
		for _, rawMirror := range rawMirrors {
			if target[rawTarget] == nil {
				target[rawTarget] = map[string]bool{}
			}
			target[rawTarget][rawMirror] = state
		}
	}
}

// splitDicts splits the target dictionary into two separate dictionaries: one for alive hosts and one for down hosts.
func splitDicts(dict TargetDict) (alive, down StorageDict) {
	alive = StorageDict{}
	down = StorageDict{}
	for target, scams := range dict {
		for scam, isAlive := range scams {
			if isAlive {
				if alive[target] == nil {
					alive[target] = []string{}
				}
				alive[target] = append(alive[target], scam)
			} else {
				if down[target] == nil {
					down[target] = []string{}
				}
				down[target] = append(down[target], scam)
			}
		}
	}
	return
}
