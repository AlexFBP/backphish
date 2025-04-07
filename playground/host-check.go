package playground

import "github.com/AlexFBP/backphish/common"

const (
	// Host states
	HostUp = iota
	HostDown
	HostUnknown
)

type TargetDict map[string](map[string]common.DummyType)

func checkHosts() {
	// Read host-up and host-down yml files
	file_up := "playground/hostup.yml"
	alive := TargetDict{}
	err := common.ReadYamlFile(file_up, &alive)
	if err != nil {
		panic(err)
	}
	file_down := "playground/hostdown.yml"
	down := TargetDict{}
	err = common.ReadYamlFile(file_down, &down)
	if err != nil {
		panic(err)
	}

	// Test all hosts
	for target, scams := range alive {
		for scam := range scams {
			switch hostState(scam) {
			case HostUp:
				// do nothing
			case HostDown:
				// move to host-down
				// remove from host-up
			case HostUnknown:
				// warn
			}
		}
	}
	for target, scams := range down {
		for scam := range scams {
			switch hostState(scam) {
			case HostUp:
				// move to host-up
				// remove from host-down
			case HostDown:
				// do nothing
			case HostUnknown:
				// warn
			}
		}
	}

	// Update the host-up and host-down files
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

func hostState(url string) int {
	// Test the host with a timeout
	// If the host is down, move it to the host-down file
	// If the host is up, move it to the host-up file
	// If the host is not reachable, retry after a delay
	// If the host is still not reachable, move it to the host-down file
	// If the host is reachable, move it to the host-up file
	return HostUnknown
}
