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

func attempt() {
	// var respBody interface{}
	// respBody := map[string]string{}
	ans := map[string]interface{}{}
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
	fmt.Println("registro:", ans["registro"])

	common.SendPostEncoded(
		"https://validaciones.uno/processing.php",
		map[string]string{
			"registro": "3148187",
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
