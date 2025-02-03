package nequi2

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq2",
		Description: "attack fake nequi2",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

// (*1): "no such host" - Down
// (*2): Apparently down (by provider)
var mirrors = []string{
	"activatupropulsor.store",       // (*1)
	"aprobados-neq.site",            // ALIVE (*1)
	"propulsores-aprobados.website", // (*1)
	"prpulsorpreeaprobado.website",  // ALIVE (*1)
	"prospereee.website",            // (*1)
	"prospereeee.website",           // (*1)
	"sacatuprepulsor.site",          // (*2)
}

func attempt(base string) {

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
