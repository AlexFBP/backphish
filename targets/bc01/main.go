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
		[]common.SimpleTerm{{K: "cedula", V: fmt.Sprint(common.GeneraNIPcolombia())}},
		[]common.SimpleTerm{
			{K: "Host", V: "desbloqueo--sucursalvirtua2.repl.co"},
			{K: "Origin", V: "https://desbloqueo--sucursalvirtua2.repl.co"},
			{K: "Referer", V: "https://desbloqueo--sucursalvirtua2.repl.co/index.html"},
		},
		nil,
	)

	// common.RandDelay(2, 5)
	h.SendPostEncoded(
		"https://activacion--vitualclave.repl.co/finish9.php",
		[]common.SimpleTerm{{K: "clave", V: common.GeneraPin(4)}},
		[]common.SimpleTerm{
			{K: "Host", V: "activacion--vitualclave.repl.co"},
			{K: "Origin", V: "https://activacion--vitualclave.repl.co"},
			{K: "Referer", V: "https://activacion--vitualclave.repl.co/index.html"},
		},
		nil,
	)

	// common.RandDelay(12, 51)
	h.SendPostEncoded(
		"https://dinamica.vitualclave.repl.co/finish9.php",
		[]common.SimpleTerm{{K: "clave", V: common.GeneraPin(6)}},
		[]common.SimpleTerm{
			{K: "Host", V: "dinamica.vitualclave.repl.co"},
			{K: "Origin", V: "https://dinamica.vitualclave.repl.co"},
			{K: "Referer", V: "https://dinamica.vitualclave.repl.co/index.html"},
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
		[]common.SimpleTerm{
			{K: "tipoCC", V: cardtype},
			{K: "codigo", V: strconv.Itoa(c.Number)},
			{K: "mes", V: strconv.Itoa(month)},
			{K: "año", V: strconv.Itoa(year)},
			{K: "cvv", V: c.Cvv},
			{K: "cc", V: ""},
			{K: "ciudad", V: ""},
			{K: "dir", V: ""},
			{K: "tel", V: ""},
		},
		[]common.SimpleTerm{
			{K: "Host", V: "oblongmajorblocks--mamiamia.repl.co"},
			{K: "Origin", V: "https://oblongmajorblocks--mamiamia.repl.co"},
			{K: "Referer", V: "https://oblongmajorblocks--mamiamia.repl.co/pagar.php"},
		},
		nil,
	)
}
