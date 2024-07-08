package bc03

import (
	"fmt"

	// fd "github.com/udistrital/utils_oas/formatdata"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	// return common.AttackRunner(attempt)
	// For single attack, comment line above and uncomment this:
	attempt()
	return nil
}

type ValidaBody struct {
	Registro int    `json:"registro"`
	Action   string `json:"action"`
}

func attempt() {
	var ans ValidaBody
	common.SendPostEncoded(
		"https://validaciones.uno/processing.php",
		map[string]string{
			"registro": "undefined",
			"tok":      "qwerty0918po22",
			"t":        "usuario",
			"usuario":  "gonorreas",
		},
		map[string]string{
			"Origin":  "https://pineapple21108900.temporary-demo.site",
			"Referer": "https://pineapple21108900.temporary-demo.site/",
		},
		&ans,
	)

	/*
		The POST above returns the following body
		{"registro":3148187,"action":"_pass.php"}
		Which is used in the following GET request
		https://validaciones.uno/_pass.php?&registro=3148187&action=_pass.php&_=1720371274419
		The second number is the unix timestamp in milliseconds
	*/

	// fd.JsonPrint(ans)
	fmt.Println("registro:", ans.Registro)

	common.SendPostEncoded(
		"https://validaciones.uno/processing.php",
		map[string]string{
			"registro": fmt.Sprint(ans.Registro),
			"tok":      "qwerty0918po22",
			"t":        "password",
			"password": "6666",
		},
		map[string]string{
			"Origin":  "https://pineapple21108900.temporary-demo.site",
			"Referer": "https://pineapple21108900.temporary-demo.site/",
		},
		nil,
	)
}
