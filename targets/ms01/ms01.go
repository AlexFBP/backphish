package ms01

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	h := common.ReqHandler{}
	// common.RandDelay(20, 60)
	domains := []string{
		"hotmail.com",
		"live.com",
		"outlook.com",
	}
	em := gofakeit.Email()
	em = fmt.Sprintf("%s@%s", strings.Split(em, "@")[0], domains[rand.Intn(len(domains))])
	form := []common.SimpleTerm{
		{K: "argaml", V: em},
		{K: "argapw", V: gofakeit.Password(gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Number(5, 20))},
	}
	pin := common.GeneraPin(4)
	form = append(form, []common.SimpleTerm{
		{K: "pn1", V: pin},
		{K: "pn2", V: pin},
	}...)
	if gofakeit.Bool() {
		form = append(form, common.SimpleTerm{K: "KMSI", V: "on"})
	}
	h.SendPostEncoded("http://recuperacion004.0hi.me/next.php",
		form,
		[]common.SimpleTerm{
			{K: "Host", V: "recuperacion004.0hi.me"},
			{K: "Origin", V: "http://recuperacion004.0hi.me"},
			{K: "Referer", V: "http://recuperacion004.0hi.me/Iniciar%20Sesi%C3%B3n.html"},
		},
		nil,
	)
}
