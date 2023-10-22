package ms01

import "github.com/AlexFBP/backphish/common"

func Cmd() {
	attempt()
}

func attempt() {
	common.SendPostEncoded("http://recuperacion004.0hi.me/next.php",
		map[string]string{
			"argaml": "metanseundedo@inutiles.com",
			"argapw": "carevergas",
			"pn1":    "3698",
			"pn2":    "3698",
			"KMSI":   "on",
		},
		map[string]string{
			"Host":    "recuperacion004.0hi.me",
			"Origin":  "http://recuperacion004.0hi.me",
			"Referer": "http://recuperacion004.0hi.me/Iniciar%20Sesi%C3%B3n.html",
		},
	)
}
