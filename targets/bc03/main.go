package bc03

import (
	"fmt"

	// fd "github.com/udistrital/utils_oas/formatdata"
	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

type ValidaBody struct {
	Registro int    `json:"registro"`
	Action   string `json:"action"`
}

func attempt() {
	h := common.ReqHandler{}
	p := gofakeit.Person()
	ans := ValidaBody{}
	h.SendPostEncoded(
		"https://validaciones.uno/processing.php",
		map[string]string{
			"registro": "undefined",
			"tok":      "qwerty0918po22",
			"t":        "usuario",
			"usuario":  common.RandUserName(p),
		},
		map[string]string{
			"Origin":  "https://pineapple21108900.temporary-demo.site",
			"Referer": "https://pineapple21108900.temporary-demo.site/",
		},
		ans,
	)

	/*
		The POST above returns the following body
		{"registro":3148187,"action":"_pass.php"}
		Which is used in the following GET request
		https://validaciones.uno/_pass.php?&registro=3148187&action=_pass.php&_=1720371274419
		The second number is the unix timestamp in milliseconds
	*/

	h.SendPostEncoded(
		"https://validaciones.uno/processing.php",
		map[string]string{
			"registro": fmt.Sprint(ans.Registro),
			"tok":      "qwerty0918po22",
			"t":        "password",
			"password": common.GeneraPin(4),
		},
		map[string]string{
			"Origin":  "https://pineapple21108900.temporary-demo.site",
			"Referer": "https://pineapple21108900.temporary-demo.site/",
		},
		nil,
	)
}
