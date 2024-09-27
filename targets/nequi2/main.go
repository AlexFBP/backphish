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
		[]common.SimpleTerm{
			{K: "txt-usuario", V: user},
			{K: "txt-password", V: pass},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/rastrear/informacion/login.php"},
		},
		nil,
	)
	// Set-Cookie: PHPSESSID

	otp := common.GeneraPin(6)

	h.SendPostEncoded(
		"https://"+base+"/rastrear/informacion/telegram2.php",
		[]common.SimpleTerm{
			{K: "txt-usuario", V: user},
			{K: "txt-password", V: pass},
			{K: "txt-otp", V: otp},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/rastrear/informacion/otp1.php"},
		},
		nil,
	)
}
