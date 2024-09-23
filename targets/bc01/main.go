package bc01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	h := common.ReqHandler{}
	// common.RandDelay(3, 10)
	h.SendPostEncoded(
		"https://desbloqueo--sucursalvirtua2.repl.co/finish9.php",
		map[string]string{"cedula": fmt.Sprint(common.GeneraNIPcolombia())},
		map[string]string{
			"Host":    "desbloqueo--sucursalvirtua2.repl.co",
			"Origin":  "https://desbloqueo--sucursalvirtua2.repl.co",
			"Referer": "https://desbloqueo--sucursalvirtua2.repl.co/index.html",
		},
		nil,
	)

	// common.RandDelay(2, 5)
	h.SendPostEncoded(
		"https://activacion--vitualclave.repl.co/finish9.php",
		map[string]string{"clave": fmt.Sprintf("%04d", common.GeneraPin(4))},
		map[string]string{
			"Host":    "activacion--vitualclave.repl.co",
			"Origin":  "https://activacion--vitualclave.repl.co",
			"Referer": "https://activacion--vitualclave.repl.co/index.html",
		},
		nil,
	)

	// common.RandDelay(12, 51)
	h.SendPostEncoded(
		"https://dinamica.vitualclave.repl.co/finish9.php",
		map[string]string{"clave": fmt.Sprintf("%06d", common.GeneraPin(6))},
		map[string]string{
			"Host":    "dinamica.vitualclave.repl.co",
			"Origin":  "https://dinamica.vitualclave.repl.co",
			"Referer": "https://dinamica.vitualclave.repl.co/index.html",
		},
		nil,
	)

	// common.RandDelay(30, 80)
	var c *gofakeit.CreditCardInfo
	var month, year, attempts int
	var cardtype string
	allowed := map[string]string{
		"Visa":             "visa",
		"MasterCard":       "master",
		"American Express": "amex",
	}
	for ; ; attempts++ {
		c = gofakeit.CreditCard()
		if v, ok := allowed[c.Type]; ok {
			cardtype = v
			break
		}
	}
	exp := strings.Split(c.Exp, "/")
	month, _ = strconv.Atoi(exp[0])
	year, _ = strconv.Atoi(exp[1])
	year += 2000
	h.SendPostEncoded(
		"https://oblongmajorblocks--mamiamia.repl.co/finish.php",
		map[string]string{
			"tipoCC": cardtype,
			"codigo": strconv.Itoa(c.Number),
			"mes":    strconv.Itoa(month),
			"año":    strconv.Itoa(year),
			"cvv":    c.Cvv,
			"cc":     "",
			"ciudad": "",
			"dir":    "",
			"tel":    "",
		},
		map[string]string{
			"Host":    "oblongmajorblocks--mamiamia.repl.co",
			"Origin":  "https://oblongmajorblocks--mamiamia.repl.co",
			"Referer": "https://oblongmajorblocks--mamiamia.repl.co/pagar.php",
		},
		nil,
	)
}
