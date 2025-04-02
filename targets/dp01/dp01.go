package dp01

import (
	"fmt"
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

type mirror struct {
	common.Telegram
}

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "dp1",
		Description: "attack fake daviplata",
		Mirrors:     getMirrorNames(),
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(base string) {
	// Entrypoint from https://ingressar1davidd.sayo1296.repl.co

	m := mirrData(base)

	tiposDocumento := []string{
		"tarjeta de identidad",
		"cedula de extranjeria",
		"cedula de ciudadania",
	}
	randIp := common.GeneraIP()
	location := "Bogotá, CO"

	m.SendToTelegram(fmt.Sprintf(`DATOS DAVPLAT
TipoDoc: %s
NumDoc: %s
Clave: %s
IP: %s
%s`,
		common.PickRand(tiposDocumento),
		common.GeneraNIPcolombia(),
		common.GeneraPin(4),
		randIp, location))

	common.RandDelayRange(30*time.Second, 60*time.Second)

	m.SendToTelegram(fmt.Sprintf(`DATOS DAVPLAT
Cod1: %s
IP: %s
%s`, common.GeneraPin(6), randIp, location))

	common.RandDelayRange(30*time.Second, 60*time.Second)
}

func getMirrorNames() (names []string) {
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}

func mirrData(name string) (d mirror) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Token, d.Chat = v[1], v[2]
			break
		}
	}
	return
}

var mirrors = [][]string{
	{"ingressar1davidd.sayo1296.repl.co", "6328508246:AAFU2THEk8nvxuzFLR5C6FfqQNuxmsiaFWk", "6332256769"}, // Partially down? still spamable?
}
