package nequi7

import (
	"fmt"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq7",
		Description: "attack fake nequi7",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(GetAllCmds())
}

func GetAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

var mirrors = []string{
	`cops.blob.core.windows.net/nequitepresta`,            //
	`nequicop.blob.core.windows.net/listo`,                //
	`nequitepresta.blob.core.windows.net/nequisalvavidas`, //
	`propulsor.blob.core.windows.net/nequicop`,            // Reported
	`propulsornequi.blob.core.windows.net/prestamonequi`,  //
	`propulsornequi.blob.core.windows.net/nequi10`,        //
	`tunequi.blob.core.windows.net/propulsor`,             //
}

var mirrorDetails = [][2]string{
	{"", ""},
	{"", ""},
	{"", ""},
	{"", ""},
	{"-1002403437567", "7795679390:AAEO8dqypaHRLPfzOjr091RVum5UV1mcQKA"},
	{"", ""},
	{"-1002337684537", "7755303993:AAFjlXcO1uXA3gOoIW8p6Js_lTh3SM6wDss"},
}

func attempt(mirror string) {
	h := common.ReqHandler{}

	det := mirrorDetails[common.FindPos(mirror, mirrors)]
	cel := common.GeneraCelColombia()
	pin := common.GeneraPin(4)
	ip := common.GeneraIP()

	h.SendJSON("https://api.telegram.org/bot"+det[1]+"/sendMessage", common.TgMsg{
		ChatID: det[0],
		Text: fmt.Sprintf("Nequi_Meta_Infinito\nNúmero: %s\nClave: %s\nIP: %s\nCiudad: %s, País: %s",
			common.AddSeparator(cel, 0, " "), pin, ip, "Bogotá", "Colombia"),
	}, nil, nil)

}
