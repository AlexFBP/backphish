package unclassified

import (
	"fmt"
	"net/http"

	"github.com/AlexFBP/backphish/common"
)

const (
	file_root = "targets/unclassified/"
	file_up   = file_root + "hostup.yaml"
	file_down = file_root + "hostdown.yaml"
)

type TargetDict map[string](map[string]bool)
type StorageDict map[string]([]string)

func CheckHosts(includeDown bool) {
	allTargets := common.TargetDict{}
	// Read host-up and host-down yml files
	alive := common.StorageDict{}
	err := common.ReadYamlFile(file_up, &alive)
	if err != nil {
		panic(err)
	}
	down := common.StorageDict{}
	err = common.ReadYamlFile(file_down, &down)
	if err != nil {
		panic(err)
	}
	common.MergeDicts(alive, allTargets, true)
	common.MergeDicts(down, allTargets, false)

	// Test all hosts
	for target, scams := range allTargets {
		for scam, state := range scams {
			if !includeDown && !state {
				continue
			}
			change := false
			// TODO: Use goroutines here
			code, httpStat := common.CheckHostState(scam)
			switch code {
			case common.HostUp:
				if !state { // if it was down
					change = true
					allTargets[target][scam] = true
					if common.CanLog(common.LOG_NORMAL) {
						fmt.Printf("Now it seems that '%s' is up\n", scam)
					}
				}
			case common.HostDown:
				if state { // if it was up
					change = true
					allTargets[target][scam] = false
					if common.CanLog(common.LOG_NORMAL) {
						fmt.Printf("Now it seems that '%s' is down\n", scam)
					}
				}
			case common.HostUnknown:
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

func updateHostFiles(allTargets common.TargetDict) {
	var alive, down common.StorageDict
	var err error

	alive, down = common.SplitDicts(allTargets)
	err = common.WriteYamlFile(file_up, alive)
	if err != nil {
		panic(err)
	}
	err = common.WriteYamlFile(file_down, down)
	if err != nil {
		panic(err)
	}
}
