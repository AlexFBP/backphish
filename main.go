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
	"github.com/AlexFBP/backphish/targets/dv1"
	mail47201 "github.com/AlexFBP/backphish/targets/mail472_01"
	"github.com/AlexFBP/backphish/targets/ms01"
	"github.com/AlexFBP/backphish/targets/nequi1"
	"github.com/AlexFBP/backphish/targets/nequi2"
	"github.com/AlexFBP/backphish/targets/nequi3"
	"github.com/AlexFBP/backphish/targets/nequi4"
	"github.com/AlexFBP/backphish/targets/nequi5"
	"github.com/AlexFBP/backphish/targets/netflix1"
)

func main() {
	conf := common.GetConfig()
	commandOptions := common.JoinSlices(
		[]menu.CommandOption{
			// Attacks - Please sort alphabetically by key
			{Command: "472-1", Description: "attack fake 4-72 1", Function: mail47201.Cmd}, // DOWN
			{Command: "bc1", Description: "attack fake bancolombia 1", Function: bc01.Cmd}, // Half down? still spamable?
			{Command: "bc2", Description: "attack fake bancolombia 2", Function: bc02.Cmd}, // DOWN
			{Command: "bc3", Description: "attack fake bancolombia 3", Function: bc03.Cmd}, // DOWN
			{Command: "dp1", Description: "attack fake daviplata 1", Function: dp01.Cmd1},  // Partially down? still spamable?
		},
		dv1.GetAllCmds(),
		[]menu.CommandOption{
			{Command: "ms1", Description: "attack fake MS login 1", Function: ms01.Cmd}, // DOWN
		},
		nequi1.GetAllCmds(),
		nequi2.GetAllCmds(), // DOWN
		nequi3.GetAllCmds(),
		[]menu.CommandOption{
			{Command: "nq4", Description: "attack fake nequi4", Function: nequi4.Cmd}, // DOWN
			{Command: "nq5", Description: "attack fake nequi5", Function: nequi5.Cmd},
			{Command: "nf1", Description: "attack fake netflix 1", Function: netflix1.Cmd}, // DOWN? (Domain still alive)
			// Playground - Please keep this one at the end
			{Command: "test", Description: "playground (not a real attack)", Function: playground.Cmd},
		},
	)

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
