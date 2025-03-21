package ms01

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit"
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "ms1",
		Description: "attack fake MS login",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(base string) {
	h := common.ReqHandler{}
	// common.RandDelay(20, 60)
	domains := []string{
		"hotmail.com",
		"live.com",
		"outlook.com",
	}
	em := gofakeit.Email()
	em = fmt.Sprintf("%s@%s", strings.Split(em, "@")[0], common.PickRand(domains))
	form := []common.SimpleTerm{
		{K: "argaml", V: em},
		{K: "argapw", V: common.RandPass1(5, 20)},
	}
	pin := common.GeneraPin(4)
	form = append(form, []common.SimpleTerm{
		{K: "pn1", V: pin},
		{K: "pn2", V: pin},
	}...)
	if gofakeit.Bool() {
		form = append(form, common.SimpleTerm{K: "KMSI", V: "on"})
	}
	h.SendPostEncoded("http://"+base+"/next.php",
		form,
		[]common.SimpleTerm{
			{K: "Host", V: "" + base + ""},
			{K: "Origin", V: "http://" + base + ""},
			{K: "Referer", V: "http://" + base + "/Iniciar%20Sesi%C3%B3n.html"},
		},
		nil,
	)
}

var mirrors = []string{
	"recuperacion004.0hi.me", // DOWN
}
