package main

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/playground"
	"github.com/AlexFBP/backphish/targets/bc01"
	"github.com/AlexFBP/backphish/targets/dp01"
	mail47201 "github.com/AlexFBP/backphish/targets/mail472_01"
)

func main() {
	commandOptions := []menu.CommandOption{
		{Command: "dp1", Description: "attack fake daviplata 1", Function: dp01.Cmd1},
		{Command: "bc1", Description: "attack fake bancolombia 1", Function: bc01.Cmd},
		{Command: "472-1", Description: "attack fake 4-72 1", Function: mail47201.Cmd},
		{Command: "test", Description: "playground (not a real attack)", Function: playground.Cmd},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help > ", 0)
	menu := menu.NewMenu(commandOptions, menuOptions)
	menu.Start()
}
