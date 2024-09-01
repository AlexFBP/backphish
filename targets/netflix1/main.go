package netflix1

func Cmd(args ...string) error {
	// return common.AttackRunner(attempt, common.ArgsHaveTimes(args...))
	// For single attack, comment line above and uncomment this:
	attempt()
	return nil
}

func attempt() {
	// GET https://mi-cuentasuscripcionflix.com/
	// to catch in response headers:
	// set-cookie - PHPSESSID=...; path=/
	// ----> To be used in subsequent requests in Cookie header

	// POST https://mi-cuentasuscripcionflix.com/captcha/set_captcha_session.php
	// (origin: https://mi-cuentasuscripcionflix.com)
	// (referer: https://mi-cuentasuscripcionflix.com/captcha/captcha.php)
	// form (content-type: application/x-www-form-urlencoded; charset=UTF-8) with values:
	// captcha: true

	// POST https://mi-cuentasuscripcionflix.com/send.php
	// (origin: https://mi-cuentasuscripcionflix.com)
	// (referer: https://mi-cuentasuscripcionflix.com/index.php)
	// form with FAKE values:
	// username: care@hotmail.com
	// password: datsacoolpassword

	// POST https://mi-cuentasuscripcionflix.com/send.php
	// (origin: https://mi-cuentasuscripcionflix.com)
	// (referer: https://mi-cuentasuscripcionflix.com/billing.php)
	// form with FAKE values:
	// name: Kaarlo Negassi
	// adresse: 800 Howard Baker Jr. Ave
	// ville: Knoxville
	// zip: 37915
	// tel: 8652157450
	// dob: 08/02/1989
	// ccc: 4342 4437 0984 1952
	// exp: 10/29
	// cvc: 671
	// titulaire: Emin Devari
}
