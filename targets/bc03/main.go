package bc03

import (
	"fmt"

	// fd "github.com/udistrital/utils_oas/formatdata"
	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "bc3"
	target.Description = "attack fake bancolombia 3"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

type ValidaBody struct {
	Registro int    `json:"registro"`
	Action   string `json:"action"`
}

func (t *mirrorTarget) Handler(base string) {
	h := common.ReqHandler{}
	p := gofakeit.Person()
	const BACK = "https://validaciones.uno/processing.php"
	ans := ValidaBody{}
	h.SendPostEncoded(
		BACK,
		[]common.SimpleTerm{
			{K: "registro", V: "undefined"},
			{K: "tok", V: "qwerty0918po22"},
			{K: "t", V: "usuario"},
			{K: "usuario", V: common.RandUserName(p)},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/"},
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
		BACK,
		[]common.SimpleTerm{
			{K: "registro", V: fmt.Sprint(ans.Registro)},
			{K: "tok", V: "qwerty0918po22"},
			{K: "t", V: "password"},
			{K: "password", V: common.GeneraPin(4)},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/"},
		},
		nil,
	)
}

var mirrors = []string{
	"pineapple21108900.temporary-demo.site", // DOWN
}
