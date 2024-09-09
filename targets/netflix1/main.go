package netflix1

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	// fd "github.com/udistrital/utils_oas/formatdata"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt, common.ArgsHaveTimes(args...))
}

func attempt() {
	h := common.ReqHandler{}
	h.UseJar(true)

	h.SendGet(
		"https://mi-cuentasuscripcionflix.com/",
		map[string]string{}, map[string]string{}, nil,
	)

	// Maybe all the following is not needed
	u, _ := url.Parse("https://mi-cuentasuscripcionflix.com")
	h.PrintCookies(u)

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/captcha/set_captcha_session.php",
		map[string]string{"captcha": "true"},
		map[string]string{
			"Origin":  "https://mi-cuentasuscripcionflix.com",
			"Referer": "https://mi-cuentasuscripcionflix.com/captcha/captcha.php",
		}, nil,
	)

	p := gofakeit.Person()
	l_min := 8
	l_max := 20
	// TODO: Choose randomly either a "strong" or a "dictionary" password
	pass := gofakeit.Password(true, true, true, true, false, rand.Intn(l_max-l_min+1)+l_min)
	// fd.JsonPrint(p)
	// fmt.Println(p)

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/send.php",
		map[string]string{
			"username": p.Contact.Email,
			"password": pass,
		},
		map[string]string{
			"Origin":  "https://mi-cuentasuscripcionflix.com",
			"Referer": "https://mi-cuentasuscripcionflix.com/index.php",
		}, nil,
	)

	h.SendGet(
		"https://mi-cuentasuscripcionflix.com/includes/get_setup.php",
		map[string]string{"auth": "FREELIVE"},
		map[string]string{"Referer": "https://mi-cuentasuscripcionflix.com/billing.php"}, nil,
	)

	n := time.Now()
	bd := gofakeit.DateRange(n.AddDate(-70, 0, 0), n.AddDate(-18, 0, 5))
	cc := strconv.Itoa(p.CreditCard.Number)
	cc = strings.Join([]string{cc[:4], cc[4:8], cc[8:12], cc[12:]}, " ")
	// fmt.Println("cardNum:", cc)
	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/send.php",
		map[string]string{
			"name":      p.FirstName + " " + p.LastName,
			"adresse":   p.Address.Street,
			"ville":     p.Address.City,
			"zip":       p.Address.Zip,
			"tel":       p.Contact.Phone,
			"dob":       bd.Format("01/02/2006"), // Refer https://pkg.go.dev/time#pkg-constants
			"ccc":       cc,
			"exp":       p.CreditCard.Exp,
			"cvc":       p.CreditCard.Cvv,
			"titulaire": p.FirstName + " " + p.LastName,
		},
		map[string]string{
			"Origin":  "https://mi-cuentasuscripcionflix.com",
			"Referer": "https://mi-cuentasuscripcionflix.com/billing.php",
		}, nil,
	)
}
