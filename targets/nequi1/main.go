package nequi1

import (
	"log"
	"net/url"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(mirrorAttempt(1), common.ArgsHaveTimes(args...))
}

var mirrors = []string{
	`aplicaya-neq.com`,
	`finanzasaturitmo.com`,
	`nqprepropulso.com`,
	`nqpropulsa.com`,
	`nuevopropulsor.com`,
	`onlineparati.com`,
	`prepropulneq.com`,
	`siperpropcolombia.com`,
}

func mirrorAttempt(n int) common.AttemptHander {
	L := len(mirrors)
	if 0 <= n && n < L {
		return func() {
			attempt(mirrors[n])
		}
	}
	return func() {
		log.Fatalf("[FATAL] Wrong mirror number, max:%d - got:%d", L, n)
	}
}

func attempt(mirrorPath string) {
	h := common.ReqHandler{}
	h.UseJar(true)

	// POST (+ Preflight!?) https://yousitesureonlineverification.com/recursos/namekusei.php
	h.SendJSON("https://yousitesureonlineverification.com/recursos/namekusei.php",
		map[string]string{"cedula": "23234654"},
		map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/",
		}, nil)

	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/pasousuario.php",
		map[string]string{
			"pass":  "3075",
			"user":  "310 350 5050",
			"dis":   "PC",
			"banco": "NEQUI",
		}, map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/neq.php",
		}, nil)
	//     SET-COOKIE:
	// usuario, contrasena, registro, estado
	u, _ := url.Parse("https://" + mirrorPath + "")
	h.PrintCookies(u)

	status := ""
	// POST https://{mirrorPath}/NEQUI/3d/process2/estado.php
	//     NOTE: This request returns a "status" number as body, related to consultar_estado()
	// of https://{mirrorPath}/NEQUI/3d/propulsor/nequi/js/functions2.js
	// (NO PAYLOAD! - MUST BE REPEATED WHILE REPLY BODY = "1" - should break with 2)
	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/estado.php", nil,
		map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/cargando.php",
		}, &status)

	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/pasoOTP.php",
		map[string]string{"otp": "726548"},
		map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/otp.php",
		}, nil)
	//     SET-COOKIE:
	// cdinamica
	h.PrintCookies(u)

	// POST https://{mirrorPath}/NEQUI/3d/process2/estado.php
	//     NOTE: This request returns a "status" number as body, related to consultar_estado()
	// of https://{mirrorPath}/NEQUI/3d/propulsor/nequi/js/functions2.js
	// (NO PAYLOAD! - MUST BE REPEATED WHILE REPLY BODY = "3" - should break with 12)
	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/estado.php", nil,
		map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/cargando.php",
		}, &status)
}
