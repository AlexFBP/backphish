package bc03

import (
	"fmt"

	// fd "github.com/udistrital/utils_oas/formatdata"
	"github.com/brianvoe/gofakeit"
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

func init() {
	common.MainMenu.Add(menu.CommandOption{
		Command: "bc3", Description: "attack fake bancolombia 3", Function: cmd, // DOWN
	})
}

func cmd(args ...string) error {
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
		[]common.SimpleTerm{
			{K: "registro", V: "undefined"},
			{K: "tok", V: "qwerty0918po22"},
			{K: "t", V: "usuario"},
			{K: "usuario", V: common.RandUserName(p)},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://pineapple21108900.temporary-demo.site"},
			{K: "Referer", V: "https://pineapple21108900.temporary-demo.site/"},
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
		[]common.SimpleTerm{
			{K: "registro", V: fmt.Sprint(ans.Registro)},
			{K: "tok", V: "qwerty0918po22"},
			{K: "t", V: "password"},
			{K: "password", V: common.GeneraPin(4)},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://pineapple21108900.temporary-demo.site"},
			{K: "Referer", V: "https://pineapple21108900.temporary-demo.site/"},
		},
		nil,
	)
}
