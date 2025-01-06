package nequi7

import (
	"fmt"
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

var workMirrors []string

func init() {
	workMirrors = getMirrorNames()

	target = &common.Target{
		Prefix:      "nq7",
		Description: "attack fake nequi7",
		Mirrors:     workMirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(mirror string) {
	h := common.ReqHandler{}

	det := mirrors[common.FindPos(mirror, workMirrors)]
	cel := common.GeneraCelColombia()
	pin := common.GeneraPin(4)
	ip := common.GeneraIP()

	city, _ := common.GeneraCiudadDeptoColombia()
	country := "Colombia"
	ced := common.GeneraNIPcolombia()
	nom := common.GeneraNombresApellidosPersonaCombinadosCol(false)

	long := ""
	if det[3] == "long" || det[3] == "alt1" || det[3] == "alt2" {
		full := ""
		if det[3] == "alt1" || det[3] == "alt2" {
			long = fmt.Sprintf("🆔Nombres: %s\n🪪Cedula: %s\n", nom, ced)
			full = fmt.Sprintf("👁Nequi Paso 1👁\n%s🌏IP: %s\n🏙Ciudad: %s\n🇨🇴País: %s",
				long, ip, city, country)
		} else {
			long = fmt.Sprintf("Nombres: %s\nCedula: %s\n", nom, ced)
			full = fmt.Sprintf("Nequi_Paso_1\n%sIP: %s\nCiudad: %s, País: %s",
				long, ip, city, country)
		}

		h.SendJSON("https://api.telegram.org/bot"+det[2]+"/sendMessage", common.TgMsg{
			ChatID: det[1],
			Text:   full,
		}, nil, nil)

		common.RandDelayRange(4*time.Second, 20*time.Second)
	}

	tfa := ""
	emoji1 := ""
	loc := ""
	if det[3] == "alt1" {
		tfa = fmt.Sprintf("⭐️Dinamica1: %s\n", common.GeneraPin(6))
		loc = fmt.Sprintf("🇨🇴Ubicación: %s", city)
		emoji1 = "🤠"
	} else if det[3] == "alt2" {
		emoji1 = "👤"
		loc = fmt.Sprintf("🇨🇴Ciudad: %s, País: %s", city, country)
	}

	if det[3] == "alt1" || det[3] == "alt2" {
		h.SendJSON("https://api.telegram.org/bot"+det[2]+"/sendMessage", common.TgMsg{
			ChatID: det[1],
			Text: fmt.Sprintf(`%sNequi_Meta_Infinito%s
  🆔Nombres: %s
  🪪Cedula: %s
  #️⃣Número: %s
  🔐Clave: %s
  %s🌏IP: %s
  %s`,
				emoji1, emoji1, nom, ced, common.AddSeparator(cel, 0, " "), pin, tfa, ip, loc),
		}, nil, nil)

	} else if det[3] == "long" || det[3] == "short" {
		h.SendJSON("https://api.telegram.org/bot"+det[2]+"/sendMessage", common.TgMsg{
			ChatID: det[1],
			Text: fmt.Sprintf("Nequi_Meta_Infinito\n%sNúmero: %s\nClave: %s\nIP: %s\nCiudad: %s, País: %s",
				long, common.AddSeparator(cel, 0, " "), pin, ip, city, country),
		}, nil, nil)
	}
	common.RandDelayRange(1*time.Second, 2*time.Second)

}
