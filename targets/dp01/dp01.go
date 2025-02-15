package dp01

import (
	"fmt"
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

func init() {
	common.MainMenu.Add(menu.CommandOption{
		Command: "dp1", Description: "attack fake daviplata 1", Function: cmd, // Partially down? still spamable?
	})
}

func cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	h := common.ReqHandler{}
	posturl := "https://api.telegram.org/bot6328508246:AAFU2THEk8nvxuzFLR5C6FfqQNuxmsiaFWk/sendMessage"
	headers := []common.SimpleTerm{
		{K: "Origin", V: "https://ingressar1davidd.sayo1296.repl.co"},
		{K: "Referer", V: "https://ingressar1davidd.sayo1296.repl.co/"},
		// Opens https://replit.com/@sayo1296/ingressar1davidd?v=1#index.html
		// Uses tg bot & chat id: 6328508246:AAFU2THEk8nvxuzFLR5C6FfqQNuxmsiaFWk 6332256769
	}

	chat_id := "6332256769"
	tiposDocumento := []string{
		"tarjeta de identidad",
		"cedula de extranjeria",
		"cedula de ciudadania",
	}
	randIp := common.GeneraIP()
	location := "Bogotá, CO"

	h.SendJSON(posturl, common.TgMsg{
		ChatID: chat_id,
		Text: fmt.Sprintf(`DATOS DAVPLAT
TipoDoc: %s
NumDoc: %s
Clave: %s
IP: %s
%s`,
			common.PickRand(tiposDocumento),
			common.GeneraNIPcolombia(),
			common.GeneraPin(4),
			randIp, location),
		// DATOS DAVPLAT
		// TipoDoc: cedula de ciudadania
		// NumDoc: 1010111010
		// Clave: 2365
		// IP: 123.123.123.123
		// Bogotá, CO
	}, headers, nil)

	common.RandDelayRange(30*time.Second, 60*time.Second)

	h.SendJSON(posturl, common.TgMsg{
		ChatID: chat_id,
		Text: fmt.Sprintf(`DATOS DAVPLAT
Cod1: %s
IP: %s
%s`, common.GeneraPin(6), randIp, location),
		// `DATOS DAVPLAT
		// Cod1: 258415
		// IP: 123.123.123.123
		// Bogotá, CO`,
	}, headers, nil)

	common.RandDelayRange(30*time.Second, 60*time.Second)
}
