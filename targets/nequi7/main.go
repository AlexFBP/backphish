package nequi7

import (
	"fmt"
	"time"

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
	`cops.blob.core.windows.net/nequitepresta`,            // ALIVE
	`nequicop.blob.core.windows.net/listo`,                // ALIVE
	`nequitepresta.blob.core.windows.net/nequisalvavidas`, // ALIVE
	`propulsor.blob.core.windows.net/nequicop`,            // Reported
	`propulsornequi.blob.core.windows.net/prestamonequi`,  // ALIVE
	`propulsornequi.blob.core.windows.net/nequi10`,        // ALIVE
	`tunequi.blob.core.windows.net/propulsor`,             // ALIVE
}

var mirrorDetails = [][3]string{
	{"-1002484241220", "7709093911:AAHlo3XrnsIV6hKcAHhijHLFdSnnLVkXIsY", "short"},
	{"-1002177027036", "7974637279:AAE9xaNbLyWeaXFViW7aav77EmUhFJr0r9s", "short"},
	{"-1002282106001", "7200967568:AAFJtTq9WBZ2Ayu0IAQ0BaXjsEH1heXTW00", "long"},
	{"-1002383703541", "7914478130:AAE2pCUs3ww9M9pKb1xAWfkQdXa8y7VFTSI", "short"},
	{"-1002403437567", "7795679390:AAEO8dqypaHRLPfzOjr091RVum5UV1mcQKA", "short"},
	{"-1002383251748", "8047026966:AAH2-41mYGwDCJLuOY5wpOOofKtj8coK0zI", "long"},
	{"-1002337684537", "7755303993:AAFjlXcO1uXA3gOoIW8p6Js_lTh3SM6wDss", "short"},
}

func attempt(mirror string) {
	h := common.ReqHandler{}

	det := mirrorDetails[common.FindPos(mirror, mirrors)]
	cel := common.GeneraCelColombia()
	pin := common.GeneraPin(4)
	ip := common.GeneraIP()
	city := "Bogotá"
	country := "Colombia"

	long := ""
	if det[2] == "long" {
		ced := common.GeneraNIPcolombia()
		long = fmt.Sprintf("Nombres: %s\nCedula: %s\n", "", ced)

		h.SendJSON("https://api.telegram.org/bot"+det[1]+"/sendMessage", common.TgMsg{
			ChatID: det[0],
			Text: fmt.Sprintf("Nequi_Paso_1\n%sIP: %s\nCiudad: %s, País: %s",
				long, ip, city, country),
		}, nil, nil)

		common.RandDelayRange(4*time.Second, 20*time.Second)
	}

	h.SendJSON("https://api.telegram.org/bot"+det[1]+"/sendMessage", common.TgMsg{
		ChatID: det[0],
		Text: fmt.Sprintf("Nequi_Meta_Infinito\n%sNúmero: %s\nClave: %s\nIP: %s\nCiudad: %s, País: %s",
			long, common.AddSeparator(cel, 0, " "), pin, ip, city, country),
	}, nil, nil)

	common.RandDelayRange(1*time.Second, 2*time.Second)

}
