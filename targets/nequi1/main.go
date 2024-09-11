package nequi1

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

func GetAllCmds(args ...string) (opts []menu.CommandOption) {
	for k, v := range mirrors {
		opts = append(opts, menu.CommandOption{
			Command:     fmt.Sprintf("nq1-%d", k+1),
			Description: fmt.Sprintf("attack fake nequi 1, mirror %d (%s)", k+1, v),
			Function:    getCmd(k, args...),
		})
	}
	return
}

func getCmd(k int, argss ...string) (f func(args ...string) error) {
	return func(args ...string) error {
		return common.AttackRunner(mirrorAttempt(k), common.ArgsHaveTimes(argss...))
	}
}

var mirrors = []string{
	`aplicaparahoy.com`,
	`aplicaya-neq.com`,
	`finanzasaturitmo.com`,
	`neq.n3quionline.com`,
	`n3quionline.com`,
	`nqipr0pulsor.com`,
	`nqprepropulso.com`, // DNS CATCHED?
	`nqpropulsa.com`,
	`nuevopropulsor.com`,
	`onlineparati.com`,
	`parati-nqui.com`,
	`prepropulneq.com`,
	`prestamo-nequi.website`,
	`siperpropcolombia.com`,
	`51.107.8.147`, // DNS CATCHED?
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
		map[string]string{"cedula": common.GeneraNIPcolombia()},
		map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/",
		}, nil)

	cel := common.GeneraCelColombia()
	opts := []string{"iPhone", "iPod", "PC"}
	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/pasousuario.php",
		map[string]string{
			"pass":  common.GeneraPin(4),
			"user":  strings.Join([]string{cel[:3], cel[3:6], cel[6:]}, " "),
			"dis":   opts[rand.Intn(len(opts))],
			"banco": "NEQUI",
		}, map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/neq.php",
		}, nil)
	//     SET-COOKIE:
	// usuario, contrasena, registro, estado
	u, _ := url.Parse("https://" + mirrorPath + "")
	h.PrintCookies(u)

	awaitStatusChange := func() {
		poll_secs := 3
		max_secs := 90
		max_attempts := max_secs / poll_secs
		status := ""
		for attempts := 0; attempts < max_attempts; attempts++ {
			// POST https://{mirrorPath}/NEQUI/3d/process2/estado.php
			//     NOTE: This request returns a "status" number as body, related to consultar_estado()
			// of https://{mirrorPath}/NEQUI/3d/propulsor/nequi/js/functions2.js
			// (NO PAYLOAD! - MUST BE REPEATED WHILE REPLY BODY = "3" - should break with 12)
			updatedStatus := ""
			h.SendPostEncoded(
				"https://"+mirrorPath+"/NEQUI/3d/process2/estado.php", nil,
				map[string]string{
					"Origin":  "https://" + mirrorPath + "",
					"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/cargando.php",
				}, &updatedStatus)
			if status == "" {
				status = updatedStatus
			} else if status != updatedStatus {
				break
			}
			time.Sleep(time.Duration(poll_secs) * time.Second)
		}
	}

	awaitStatusChange()

	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/pasoOTP.php",
		map[string]string{"otp": common.GeneraPin(6)},
		map[string]string{
			"Origin":  "https://" + mirrorPath + "",
			"Referer": "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/otp.php",
		}, nil)
	//     SET-COOKIE:
	// cdinamica
	h.PrintCookies(u)

	awaitStatusChange()
}
