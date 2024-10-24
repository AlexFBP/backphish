package common

import (
	"github.com/turret-io/go-menu/menu"
)

type MenuManager_t struct {
	menu []menu.CommandOption
}

func (mm *MenuManager_t) Add(entry menu.CommandOption) {
	name := entry.Command
	if mm.Created(name) {
		panic(`entry "` + name + `" was previously added`)
	}
	mm.menu = append(mm.menu, entry)
}

func (mm *MenuManager_t) AddMany(entries []menu.CommandOption) {
	for _, entry := range entries {
		mm.Add(entry)
	}
}

func (mm *MenuManager_t) Created(name string) bool {
	for _, entry := range mm.menu {
		if entry.Command == name || entry.Command == "test" {
			return true
		}
	}
	return false
}

func (mm *MenuManager_t) GetAll() []menu.CommandOption {
	// TODO: Sort before returning
	return mm.menu
}

var MainMenu MenuManager_t
