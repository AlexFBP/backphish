package playground

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/AlexFBP/backphish/common"
)

const (
	// Host states
	HostUp = iota
	HostDown
	HostUnknown

	file_up   = "playground/hostup.yml"
	file_down = "playground/hostdown.yml"
)

type TargetDict map[string](map[string]bool)
type StorageDict map[string]([]string)

func checkHosts(includeDown bool) {
	allTargets := TargetDict{}
	// Read host-up and host-down yml files
	alive := StorageDict{}
	err := common.ReadYamlFile(file_up, &alive)
	if err != nil {
		panic(err)
	}
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
			if !includeDown && !state {
				continue
			}
			change := false
			// TODO: Use goroutines here
			code, httpStat := checkHostState(scam)
			switch code {
			case HostUp:
				if !state { // if it was down
					change = true
					allTargets[target][scam] = true
					if common.CanLog(common.LOG_NORMAL) {
						fmt.Printf("Now it seems that '%s' is up\n", scam)
					}
				}
			case HostDown:
				if state { // if it was up
					change = true
					allTargets[target][scam] = false
					if common.CanLog(common.LOG_NORMAL) {
						fmt.Printf("Now it seems that '%s' is down\n", scam)
					}
				}
			case HostUnknown:
				// warn
				if common.CanLog(common.LOG_NORMAL) {
					fmt.Printf("Unknown state for host '%s' - HTTP Stat:%d - %s\n",
						scam, httpStat, http.StatusText(httpStat))
				}
			}
			if change {
				updateHostFiles(allTargets)
			}
		}
	}
}

func updateHostFiles(allTargets TargetDict) {
	var alive, down StorageDict
	var err error

	alive, down = splitDicts(allTargets)
	err = common.WriteYamlFile(file_up, alive)
	if err != nil {
		panic(err)
	}
	err = common.WriteYamlFile(file_down, down)
	if err != nil {
		panic(err)
	}
}

func checkHostState(url string) (state, httpCode int) {
	err := common.CheckURL(url)
	state = HostUnknown
	if err != nil {
		return
	}
	h := common.ReqHandler{}
	var step uint8
	err, step = h.SendGet(url, nil, nil, nil)
	if err != nil {
		if common.CanLog(common.LOG_VERBOSE) {
			fmt.Printf("Host %s seems to be down - lastStep: %d - err: %s\n", url, step, err)
		}
		return HostDown, 0
	}

	httpCode = h.Response.StatusCode
	switch httpCode {
	case http.StatusUnavailableForLegalReasons:
		state = HostDown
	default:
		if httpCode >= 100 && httpCode < 300 {
			if common.CanLog(common.LOG_VERBOSE) {
				fmt.Printf("Host %s seems to be up - lastStep: %d\n", url, step)
			}
			state = HostUp
		}
	}
	return
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
	sortMirrors(alive)
	sortMirrors(down)
	return
}

func sortMirrors(stDict StorageDict) {
	cmp := func(a, b string) int {
		a_parsed, err_a := url.Parse(a)
		b_parsed, err_b := url.Parse(b)

		// Compare validness of URLs
		if err_a != nil && err_b != nil {
			return 0
		} else if err_a != nil {
			return -1
		} else if err_b != nil {
			return 1
		}

		// Compare hostnames
		if sub := compareSubdomains(a_parsed.Hostname(), b_parsed.Hostname()); sub != 0 {
			return sub
		}

		// Compare paths
		a_path := strings.ToLower(a_parsed.Path)
		b_path := strings.ToLower(b_parsed.Path)
		if a_path != b_path {
			return strings.Compare(a_path, b_path)
		}

		// Default to comparing the full URL
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	}
	for target, mirrors := range stDict {
		slices.SortFunc(mirrors, cmp)
		stDict[target] = mirrors
	}
}

func compareSubdomains(a, b string) int {
	a_parts := strings.Split(a, ".")
	b_parts := strings.Split(b, ".")
	a_len := len(a_parts)
	b_len := len(b_parts)
	common := minInt(a_len, b_len)

	// Compare each equivalent part
	for i := 1; i <= common; i++ {
		test_a := strings.ToLower(a_parts[a_len-i])
		test_b := strings.ToLower(b_parts[b_len-i])
		if test_a != test_b {
			return strings.Compare(test_a, test_b)
		}
	}
	return 0
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
