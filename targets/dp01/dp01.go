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

var m mirror

func init() {
	m.Token = "6328508246:AAFU2THEk8nvxuzFLR5C6FfqQNuxmsiaFWk"
	m.Chat = "6332256769"
	common.MainMenu.Add(menu.CommandOption{
		Command: "dp1", Description: "attack fake daviplata 1", Function: cmd, // Partially down? still spamable?
	})
}

func cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	// Entrypoint from https://ingressar1davidd.sayo1296.repl.co

	h := common.ReqHandler{}

	tiposDocumento := []string{
		"tarjeta de identidad",
		"cedula de extranjeria",
		"cedula de ciudadania",
	}
	randIp := common.GeneraIP()
	location := "Bogotá, CO"

	m.SendToTelegram(&h, fmt.Sprintf(`DATOS DAVPLAT
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

	m.SendToTelegram(&h, fmt.Sprintf(`DATOS DAVPLAT
Cod1: %s
IP: %s
%s`, common.GeneraPin(6), randIp, location))

	common.RandDelayRange(30*time.Second, 60*time.Second)
}
