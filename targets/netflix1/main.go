package netflix1

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit"
	// fd "github.com/udistrital/utils_oas/formatdata"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	h := common.ReqHandler{}
	h.UseJar(true)

	h.SendGet(
		"https://mi-cuentasuscripcionflix.com/",
		[]common.SimpleTerm{}, []common.SimpleTerm{}, nil,
	)

	// Maybe all the following is not needed
	u, _ := url.Parse("https://mi-cuentasuscripcionflix.com")
	h.PrintCookies(u)

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/captcha/set_captcha_session.php",
		[]common.SimpleTerm{{K: "captcha", V: "true"}},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://mi-cuentasuscripcionflix.com"},
			{K: "Referer", V: "https://mi-cuentasuscripcionflix.com/captcha/captcha.php"},
		}, nil,
	)

	p := gofakeit.Person()
	l_min := 8
	l_max := 20
	// TODO: Choose randomly either a "strong" or a "dictionary" password
	pass := gofakeit.Password(true, true, true, true, false, rand.Intn(l_max-l_min+1)+l_min)
	if common.CanLog(common.LOG_VERBOSE) {
		// fd.JsonPrint(p)
		fmt.Println(p)
	}

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/send.php",
		[]common.SimpleTerm{
			{K: "username", V: p.Contact.Email},
			{K: "password", V: pass},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://mi-cuentasuscripcionflix.com"},
			{K: "Referer", V: "https://mi-cuentasuscripcionflix.com/index.php"},
		}, nil,
	)

	h.SendGet(
		"https://mi-cuentasuscripcionflix.com/includes/get_setup.php",
		[]common.SimpleTerm{{K: "auth", V: "FREELIVE"}},
		[]common.SimpleTerm{{K: "Referer", V: "https://mi-cuentasuscripcionflix.com/billing.php"}}, nil,
	)

	n := time.Now()
	bd := gofakeit.DateRange(n.AddDate(-70, 0, 0), n.AddDate(-18, 0, 5))
	cc := strconv.Itoa(p.CreditCard.Number)
	cc = common.AddSeparator(cc, 1, " ")
	if common.CanLog(common.LOG_VERBOSE) {
		fmt.Println("cardNum:", cc)
	}
	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/send.php",
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
			{K: "Origin", V: "https://mi-cuentasuscripcionflix.com"},
			{K: "Referer", V: "https://mi-cuentasuscripcionflix.com/billing.php"},
		}, nil,
	)
}
