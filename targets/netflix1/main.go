package netflix1

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	// fd "github.com/udistrital/utils_oas/formatdata"
	"github.com/brianvoe/gofakeit"
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nf1",
		Description: "attack fake netflix 1",
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
	h.UseJar(true)

	h.SendGet(
		"https://"+base+"/",
		[]common.SimpleTerm{}, []common.SimpleTerm{}, nil,
	)

	// Maybe all the following is not needed
	u, _ := url.Parse("https://" + base)
	h.PrintCookies(u)

	h.SendPostEncoded(
		"https://"+base+"/captcha/set_captcha_session.php",
		[]common.SimpleTerm{{K: "captcha", V: "true"}},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/captcha/captcha.php"},
		}, nil,
	)

	p := gofakeit.Person()
	l_min := 8
	l_max := 20
	// TODO: Choose randomly either a "strong" or a "dictionary" password
	pass := common.RandPass1(l_min, l_max)
	if common.CanLog(common.LOG_VERBOSE) {
		// fd.JsonPrint(p)
		fmt.Println(p)
	}

	h.SendPostEncoded(
		"https://"+base+"/send.php",
		[]common.SimpleTerm{
			{K: "username", V: p.Contact.Email},
			{K: "password", V: pass},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/index.php"},
		}, nil,
	)

	h.SendGet(
		"https://"+base+"/includes/get_setup.php",
		[]common.SimpleTerm{{K: "auth", V: "FREELIVE"}},
		[]common.SimpleTerm{{K: "Referer", V: "https://" + base + "/billing.php"}}, nil,
	)

	n := time.Now()
	bd := gofakeit.DateRange(n.AddDate(-70, 0, 0), n.AddDate(-18, 0, 5))
	cc := strconv.Itoa(p.CreditCard.Number)
	cc = common.AddSeparator(cc, 1, " ")
	if common.CanLog(common.LOG_VERBOSE) {
		fmt.Println("cardNum:", cc)
	}
	h.SendPostEncoded(
		"https://"+base+"/send.php",
		[]common.SimpleTerm{
			{K: "name", V: p.FirstName + " " + p.LastName},
			{K: "adresse", V: p.Address.Street},
			{K: "ville", V: p.Address.City},
			{K: "zip", V: p.Address.Zip},
			{K: "tel", V: p.Contact.Phone},
			{K: "dob", V: bd.Format("01/02/2006")}, // Refer https://pkg.go.dev/time#pkg-constants
			{K: "ccc", V: cc},
			{K: "exp", V: p.CreditCard.Exp},
			{K: "cvc", V: p.CreditCard.Cvv},
			{K: "titulaire", V: p.FirstName + " " + p.LastName},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/billing.php"},
		}, nil,
	)
}

var mirrors = []string{
	"mi-cuentasuscripcionflix.com", // DOWN? (Domain still alive)
}
