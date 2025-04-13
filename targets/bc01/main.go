package bc01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetBase
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "bc1"
	target.Description = "attack fake bancolombia 1"
	common.MainMenu.Register(target)
}

func (t *mirrorTarget) GetMirrors() (names []string) {
	names = make([]string, 0)
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}

var mirrors = [][]string{
	{
		"branch1",                             // Half down? still spamable?
		"desbloqueo--sucursalvirtua2.repl.co", // DOWN?
		"activacion--vitualclave.repl.co",     // Opens web editor https://replit.com/@vitualclave/dinamica#index.html - Found tg token & chat! 6527691794:AAF599l03u2LDqzEGvLJxLPmTn5hOel0eY0 5786097476
		"dinamica.vitualclave.repl.co",        // Opens web editor https://replit.com/@vitualclave/dinamica#index.html - Found tg token & chat! 6527691794:AAF599l03u2LDqzEGvLJxLPmTn5hOel0eY0 5786097476
		"oblongmajorblocks--mamiamia.repl.co", // DOWN?
	},
}

func getMirrorData(name string) (steps []string) {
	for _, v := range mirrors {
		if name == v[0] {
			steps = v[1:len(mirrors)]
			break
		}
	}
	return
}

func (t *mirrorTarget) Handler(branch string) {
	h := common.ReqHandler{}
	m := getMirrorData(branch)
	// common.RandDelay(3, 10)
	h.SendPostEncoded(
		"https://"+m[0]+"/finish9.php",
		[]common.SimpleTerm{{K: "cedula", V: fmt.Sprint(common.GeneraNIPcolombia())}},
		[]common.SimpleTerm{
			{K: "Host", V: "" + m[0]},
			{K: "Origin", V: "https://" + m[0]},
			{K: "Referer", V: "https://" + m[0] + "/index.html"},
		},
		nil,
	)

	// common.RandDelay(2, 5)
	h.SendPostEncoded(
		"https://"+m[1]+"/finish9.php",
		[]common.SimpleTerm{{K: "clave", V: common.GeneraPin(4)}},
		[]common.SimpleTerm{
			{K: "Host", V: "" + m[1]},
			{K: "Origin", V: "https://" + m[1]},
			{K: "Referer", V: "https://" + m[1] + "/index.html"},
		},
		nil,
	)

	// common.RandDelay(12, 51)
	h.SendPostEncoded(
		"https://"+m[2]+"/finish9.php",
		[]common.SimpleTerm{{K: "clave", V: common.GeneraPin(6)}},
		[]common.SimpleTerm{
			{K: "Host", V: "" + m[2]},
			{K: "Origin", V: "https://" + m[2]},
			{K: "Referer", V: "https://" + m[2] + "/index.html"},
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
		"https://"+m[3]+"/finish.php",
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
			{K: "Host", V: "" + m[3]},
			{K: "Origin", V: "https://" + m[3]},
			{K: "Referer", V: "https://" + m[3] + "/pagar.php"},
		},
		nil,
	)
}
