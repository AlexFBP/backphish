package netflix1

import (
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
	// u, _ := url.Parse("https://mi-cuentasuscripcionflix.com")
	// const COOKIE_NAME = "PHPSESSID"
	// sess := ""
	// for _, cookie := range h.Jar.Cookies(u) {
	// 	fmt.Printf("%s\t%s\n", cookie.Name, cookie.Value)
	// 	if cookie.Name == COOKIE_NAME {
	// 		sess = cookie.Value
	// 	}
	// }

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/captcha/set_captcha_session.php",
		map[string]string{"captcha": "true"},
		map[string]string{
			"Origin":  "https://mi-cuentasuscripcionflix.com",
			"Referer": "https://mi-cuentasuscripcionflix.com/captcha/captcha.php",
		}, nil,
	)

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/send.php",
		map[string]string{
			"username": "care@hotmail.com",
			"password": "datsacoolpassword",
		},
		map[string]string{
			"Origin":  "https://mi-cuentasuscripcionflix.com",
			"Referer": "https://mi-cuentasuscripcionflix.com/index.php",
		}, nil,
	)

	h.SendPostEncoded(
		"https://mi-cuentasuscripcionflix.com/send.php",
		map[string]string{
			"name":      "Kaarlo Negassi",
			"adresse":   "800 Howard Baker Jr. Ave",
			"ville":     "Knoxville",
			"zip":       "37915",
			"tel":       "8652157450",
			"dob":       "08/02/1989",
			"ccc":       "4342 4437 0984 1952",
			"exp":       "10/29",
			"cvc":       "671",
			"titulaire": "Emin Devari",
		},
		map[string]string{
			"Origin":  "https://mi-cuentasuscripcionflix.com",
			"Referer": "https://mi-cuentasuscripcionflix.com/billing.php",
		}, nil,
	)
}
