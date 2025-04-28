package nequi2

import (
	"time"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "nq2"
	target.Description = "attack fake nequi2"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

// (*1): "no such host" - Down
// (*2): Apparently down (by provider)
var mirrors = []string{
	"activatupropulsor.store",       // (*1)
	"aprobados-neq.site",            // (*1)
	"propulsores-aprobados.website", // (*1)
	"prpulsorpreeaprobado.website",  // (*1)
	"prospereee.website",            // (*1)
	"prospereeee.website",           // (*1)
	"sacatuprepulsor.site",          // (*2)
}

func (t *mirrorTarget) Handler(base string) {

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

func (t *mirrorTarget) EstimateDuration() time.Duration {
	return 35 * time.Second
}

func (t *mirrorTarget) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/rastrear/informacion/login.php")
	return state
}
