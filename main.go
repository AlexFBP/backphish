package main

import (
	"fmt"

	"github.com/turret-io/go-menu/menu"

	// "github.com/AlexFBP/backphish/common"
	"github.com/AlexFBP/backphish/playground"
	"github.com/AlexFBP/backphish/targets/bc01"
	"github.com/AlexFBP/backphish/targets/bc02"
	"github.com/AlexFBP/backphish/targets/bc03"
	"github.com/AlexFBP/backphish/targets/dp01"
	mail47201 "github.com/AlexFBP/backphish/targets/mail472_01"
	"github.com/AlexFBP/backphish/targets/ms01"
)

func main() {
	// common.SetMockServer("http://localhost:1080")
	parseFlags()
	commandOptions := []menu.CommandOption{
		// Attacks - Please sort alphabetically by key
		{Command: "472-1", Description: "attack fake 4-72 1", Function: mail47201.Cmd},
		{Command: "bc1", Description: "attack fake bancolombia 1", Function: bc01.Cmd},
		{Command: "bc2", Description: "attack fake bancolombia 2", Function: bc02.Cmd},
		{Command: "bc3", Description: "attack fake bancolombia 3", Function: bc03.Cmd},
		{Command: "dp1", Description: "attack fake daviplata 1", Function: dp01.Cmd1},
		{Command: "ms1", Description: "attack fake MS login 1", Function: ms01.Cmd},
		// Playground - Please let this one at the end
		{Command: "test", Description: "playground (not a real attack)", Function: playground.Cmd},
	}
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
			fmt.Printf("Using target \"%s\"", target)
			if times > 0 {
				fmt.Printf(", only \"%d\" times", times)
			} else {
				fmt.Print(", no limit")
			}
			fmt.Println()
			command.Function(fmt.Sprintf("%d", times))
		} else {
			fmt.Printf("\"%s\" is not a valid target\n", target)
		}
	}
}
