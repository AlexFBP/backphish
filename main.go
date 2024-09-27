package main

import (
	"fmt"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
	"github.com/AlexFBP/backphish/playground"
	"github.com/AlexFBP/backphish/targets/bc01"
	"github.com/AlexFBP/backphish/targets/bc02"
	"github.com/AlexFBP/backphish/targets/bc03"
	"github.com/AlexFBP/backphish/targets/dp01"
	mail47201 "github.com/AlexFBP/backphish/targets/mail472_01"
	"github.com/AlexFBP/backphish/targets/ms01"
	"github.com/AlexFBP/backphish/targets/nequi1"
	"github.com/AlexFBP/backphish/targets/nequi2"
	"github.com/AlexFBP/backphish/targets/nequi3"
	"github.com/AlexFBP/backphish/targets/netflix1"
)

func main() {
	conf := common.GetConfig()
	commandOptions := []menu.CommandOption{
		// Attacks - Please sort alphabetically by key
		{Command: "472-1", Description: "attack fake 4-72 1", Function: mail47201.Cmd},
		{Command: "bc1", Description: "attack fake bancolombia 1", Function: bc01.Cmd},
		{Command: "bc2", Description: "attack fake bancolombia 2", Function: bc02.Cmd},
		{Command: "bc3", Description: "attack fake bancolombia 3", Function: bc03.Cmd},
		{Command: "dp1", Description: "attack fake daviplata 1", Function: dp01.Cmd1},
		{Command: "ms1", Description: "attack fake MS login 1", Function: ms01.Cmd},
		{Command: "nf1", Description: "attack fake netflix 1", Function: netflix1.Cmd},
	}
	commandOptions = append(commandOptions, nequi1.GetAllCmds()...)
	commandOptions = append(commandOptions, []menu.CommandOption{
		{Command: "nq2", Description: "attack fake nequi2", Function: nequi2.Cmd},
		{Command: "nq3", Description: "attack fake nequi3", Function: nequi3.Cmd},
		// Playground - Please keep this one at the end
		{Command: "test", Description: "playground (not a real attack)", Function: playground.Cmd},
	}...)

	target, times, threads := conf.GetTarget(), conf.GetTimes(), conf.GetThreads()
	if target == "" {
		menuOptions := menu.NewMenuOptions("'menu' for help > ", 0)
		menu := menu.NewMenu(commandOptions, menuOptions)
		menu.Start()
	} else {
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
}
