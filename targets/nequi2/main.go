package nequi2

import (
	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	const base = "propulsorpreaprobado.website"

	h := common.ReqHandler{}
	h.UseJar(true)

	user := common.GeneraCelColombia()
	pass := common.GeneraPin(4)

	h.SendPostEncoded(
		"https://"+base+"/rastrear/informacion/telegram.php",
		map[string]string{
			"txt-usuario":  user,
			"txt-password": pass,
		},
		map[string]string{
			"Origin":  "https://" + base,
			"Referer": "https://" + base + "/rastrear/informacion/login.php",
		},
		nil,
	)
	// Set-Cookie: PHPSESSID

	otp := common.GeneraPin(6)

	h.SendPostEncoded(
		"https://"+base+"/rastrear/informacion/telegram2.php",
		map[string]string{
			"txt-usuario":  user,
			"txt-password": pass,
			"txt-otp":      otp,
		},
		map[string]string{
			"Origin":  "https://" + base,
			"Referer": "https://" + base + "/rastrear/informacion/otp1.php",
		},
		nil,
	)
}
