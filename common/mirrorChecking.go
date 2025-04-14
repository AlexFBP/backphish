package common

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

const (
	// Host states
	HostUp = iota
	HostDown
	HostUnknown
)

type TargetDict map[string](map[string]bool)
type StorageDict map[string]([]string)

func CheckHostState(url string) (state, httpCode int) {
	err := CheckURL(url)
	state = HostUnknown
	if err != nil {
		return
	}
	h := ReqHandler{}
	var step uint8
	err, step = h.SendGet(url, nil, nil, nil)
	if err != nil {
		if CanLog(LOG_VERBOSE) {
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
			if CanLog(LOG_VERBOSE) {
				fmt.Printf("Host %s seems to be up - lastStep: %d\n", url, step)
			}
			state = HostUp
		}
	}
	return
}

// mergeDicts merges the source dictionary into the target dictionary, with a preset state for each entry.
func MergeDicts(source StorageDict, target TargetDict, state bool) {
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
func SplitDicts(dict TargetDict) (alive, down StorageDict) {
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
	SortMirrors(alive)
	SortMirrors(down)
	return
}

func SortMirrors(stDict StorageDict) {
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
		if sub := CompareSubdomains(a_parsed.Hostname(), b_parsed.Hostname()); sub != 0 {
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

func CompareSubdomains(a, b string) int {
	a_parts := strings.Split(a, ".")
	b_parts := strings.Split(b, ".")
	a_len := len(a_parts)
	b_len := len(b_parts)
	common := MinInt(a_len, b_len)

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

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
