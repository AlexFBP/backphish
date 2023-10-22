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
	// For single attack, comment line above and uncomment this:
	// attempt()
	// return nil
}

func attempt() {
	common.RandDelay(20, 60)
	domains := []string{
		"hotmail.com",
		"live.com",
		"outlook.com",
	}
	em := gofakeit.Email()
	em = fmt.Sprintf("%s@%s", strings.Split(em, "@")[0], domains[rand.Intn(len(domains))])
	form := map[string]string{
		"argaml": em,
		"argapw": gofakeit.Password(gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Bool(), gofakeit.Number(5, 20)),
	}
	pin := fmt.Sprintf("%04d", common.GeneraPin(4))
	form["pn1"] = pin
	form["pn2"] = pin
	if gofakeit.Bool() {
		form["KMSI"] = "on"
	}
	common.SendPostEncoded("http://recuperacion004.0hi.me/next.php",
		form,
		map[string]string{
			"Host":    "recuperacion004.0hi.me",
			"Origin":  "http://recuperacion004.0hi.me",
			"Referer": "http://recuperacion004.0hi.me/Iniciar%20Sesi%C3%B3n.html",
		},
	)
}
