package main

import (
	"fmt"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
	"github.com/AlexFBP/backphish/playground"
	_ "github.com/AlexFBP/backphish/targets/bc01"
	_ "github.com/AlexFBP/backphish/targets/bc03"
	_ "github.com/AlexFBP/backphish/targets/bdv1"
	_ "github.com/AlexFBP/backphish/targets/dp01"
	_ "github.com/AlexFBP/backphish/targets/dv1"
	_ "github.com/AlexFBP/backphish/targets/mail472_01"
	_ "github.com/AlexFBP/backphish/targets/ms01"
	_ "github.com/AlexFBP/backphish/targets/nequi1"
	_ "github.com/AlexFBP/backphish/targets/nequi2"
	_ "github.com/AlexFBP/backphish/targets/nequi3"
	_ "github.com/AlexFBP/backphish/targets/nequi4"
	_ "github.com/AlexFBP/backphish/targets/nequi5"
	_ "github.com/AlexFBP/backphish/targets/nequi6"
	_ "github.com/AlexFBP/backphish/targets/nequi7"
	_ "github.com/AlexFBP/backphish/targets/nequi8"
	_ "github.com/AlexFBP/backphish/targets/nequi9"
	_ "github.com/AlexFBP/backphish/targets/netflix1"
)

func main() {
	conf := common.GetConfig()
	commandOptions := getCommandOptions()

	target, times, threads := conf.GetTarget(), conf.GetTimes(), conf.GetThreads()
	if target == "" {
		menuOptions := menu.NewMenuOptions("'menu' for help > ", 0)
		menu := menu.NewMenu(commandOptions, menuOptions)
		menu.Start()
	} else {
		executeTarget(commandOptions, target, times, threads)
	}
}

// getCommandOptions returns the list of command options.
func getCommandOptions() []menu.CommandOption {
	return common.JoinSlices(
		common.MainMenu.GetAll(),
		[]menu.CommandOption{
			// Playground - Please keep this one at the end
			{Command: "test", Description: "playground (not a real attack)", Function: playground.Cmd},
		},
	)
}

// executeTarget executes the specified target with the given parameters.
func executeTarget(commandOptions []menu.CommandOption, target string, times, threads int) {
	var command *menu.CommandOption
	for _, v := range commandOptions {
		if v.Command == target {
			command = &v
			break
		}
	}
	if command != nil {
		if common.CanLog(common.LOG_NORMAL) {
			fmt.Printf("Using target \"%s\"", target)
			if times > 0 {
				fmt.Printf(", only \"%d\" times", times)
			} else {
				fmt.Print(", no limit")
			}
			fmt.Printf(" in %d threads", threads)
			fmt.Println()
		}
		command.Function()
	} else {
		if common.CanLog(common.LOG_QUIET) {
			fmt.Printf("\"%s\" is not a valid target\n", target)
		}
	}
}
