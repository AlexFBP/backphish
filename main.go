package main

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/dp01"
)

func main() {
	commandOptions := []menu.CommandOption{
		menu.CommandOption{"dp1", "attack fake daviplata 1", dp01.Cmd1},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help > ", 0)
	menu := menu.NewMenu(commandOptions, menuOptions)
	menu.Start()
}
