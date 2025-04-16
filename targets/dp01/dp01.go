package dp01

import (
	"fmt"
	"time"

	"github.com/AlexFBP/backphish/common"
)

type mirrorManager struct {
	common.TargetWithParams
}

var target *mirrorManager

type mirror struct {
	common.Telegram
}

func init() {
	target = &mirrorManager{}
	target.Prefix = "dp1"
	target.Description = "attack fake daviplata"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *mirrorManager) Handler(base string) {
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

func mirrData(name string) (d mirror) {
	v := target.GetMirrorParams(name)
	if len(v) < 2 {
		return
	}
	d.Token, d.Chat = v[0], v[1]
	return
}

var mirrors = [][]string{
	{"ingressar1davidd.sayo1296.repl.co", "6328508246:AAFU2THEk8nvxuzFLR5C6FfqQNuxmsiaFWk", "6332256769"}, // Partially down? still spamable?
}

func (t *mirrorManager) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/")
	return state
}
