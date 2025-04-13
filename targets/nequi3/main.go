package nequi3

import (
	"fmt"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "nq3"
	target.Description = "attack fake nequi3"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

// Mirrors. The comment depending on last checked state:
//   - (*1): "no such host" - Down
var mirrors = []string{
	`ahorrosnq.com`,         // REPORTED (*1)
	`creatusahorros.com`,    // REPORTED (*1)
	`impulsadorahorros.com`, // REPORTED (*1)
	`neqpropulsaya.com`,     // CATCHED (*1)
	`neqpronet.online`,      // (*1)
	`prestadecision.com`,    // REPORTED (*1)
	`propuladmin.com`,       // REPORTED (*1)
	`tucredirapido.com`,     // REPORTED (*1)
}

func (t *mirrorTarget) Handler(base string) {
	h := common.ReqHandler{}
	h.UseJar(true)

	cel := common.GeneraCelColombia()
	h.SendPostEncoded(
		"https://"+base+"/process/pasousuario.php",
		[]common.SimpleTerm{
			{K: "pass", V: fmt.Sprintf("%s - %sOTP %s", // "316 546 4981 - 3795OTP 524844"
				common.AddSeparator(cel, 0, " "),
				common.GeneraPin(4), common.GeneraPin(6))},
			{K: "dis", V: "Android"},
			{K: "banco", V: "bancolombia"},
		},
		[]common.SimpleTerm{
			{K: "Host", V: base},
			{K: "Accept", V: "*/*"},
			{K: "Accept-Language", V: "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3"},
			{K: "Accept-Encoding", V: "gzip, deflate, br, zstd"},
			{K: "X-Requested-With", V: "XMLHttpRequest"},
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/bdigital/otp.php?p"},
		},
		nil,
	)
}
