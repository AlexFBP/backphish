package nequi1

import (
	"net/url"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt, common.ArgsHaveTimes(args...))
}

func attempt() {
	h := common.ReqHandler{}
	h.UseJar(true)

	// POST (+ Preflight!?) https://yousitesureonlineverification.com/recursos/namekusei.php
	h.SendJSON("https://yousitesureonlineverification.com/recursos/namekusei.php",
		map[string]string{"cedula": "23234654"},
		map[string]string{
			"Origin":  "https://finanzasaturitmo.com",
			"Referer": "https://finanzasaturitmo.com/",
		}, nil)

	h.SendPostEncoded(
		"https://finanzasaturitmo.com/NEQUI/3d/process2/pasousuario.php",
		map[string]string{
			"pass":  "3075",
			"user":  "310 350 5050",
			"dis":   "PC",
			"banco": "NEQUI",
		}, map[string]string{
			"Origin":  "https://finanzasaturitmo.com",
			"Referer": "https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/neq.php",
		}, nil)
	//     SET-COOKIE:
	// usuario, contrasena, registro, estado
	u, _ := url.Parse("https://finanzasaturitmo.com")
	h.PrintCookies(u)

	status := ""
	// POST https://finanzasaturitmo.com/NEQUI/3d/process2/estado.php
	//     NOTE: This request returns a "status" number as body, related to consultar_estado()
	// of https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/js/functions2.js
	// (NO PAYLOAD! - MUST BE REPEATED WHILE REPLY BODY = "1" - should break with 2)
	h.SendPostEncoded(
		"https://finanzasaturitmo.com/NEQUI/3d/process2/estado.php", nil,
		map[string]string{
			"Origin":  "https://finanzasaturitmo.com",
			"Referer": "https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/cargando.php",
		}, &status)

	h.SendPostEncoded(
		"https://finanzasaturitmo.com/NEQUI/3d/process2/pasoOTP.php",
		map[string]string{"otp": "726548"},
		map[string]string{
			"Origin":  "https://finanzasaturitmo.com",
			"Referer": "https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/otp.php",
		}, nil)
	//     SET-COOKIE:
	// cdinamica
	h.PrintCookies(u)

	// POST https://finanzasaturitmo.com/NEQUI/3d/process2/estado.php
	//     NOTE: This request returns a "status" number as body, related to consultar_estado()
	// of https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/js/functions2.js
	// (NO PAYLOAD! - MUST BE REPEATED WHILE REPLY BODY = "3" - should break with 12)
	h.SendPostEncoded(
		"https://finanzasaturitmo.com/NEQUI/3d/process2/estado.php", nil,
		map[string]string{
			"Origin":  "https://finanzasaturitmo.com",
			"Referer": "https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/cargando.php",
		}, &status)
}
