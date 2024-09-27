package nequi1

import (
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq1",
		Description: "attack fake nequi 1",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
}

func GetAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

// (*1): "no such host" - Down
// (*2): TLS error? Down?
// (*3): "no route to host" - Down?
var mirrors = []string{
	`aplicaparahoy.com`,          // CATCHED (*1)
	`aplicaya-neq.com`,           // NO DNS REPLY (*1)
	`co.nqicolmbia.com`,          // REPORTED (*1)
	`cuztco.com/NEQUI`,           // REPORTED
	`finanzasaturitmo.com`,       // NO DNS REPLY (*1)
	`impuestoscol.com`,           // DOWN BY PROVIDER (*??)
	`impulsatunq.com`,            // REPORTED
	`intelcore.online`,           // REPORTED (*1)
	`n3quionline.com`,            // REPORTED (*1)
	`neq.n3quionline.com`,        // REPORTED (*1)
	`neqwtx.com`,                 // CATCHED
	`newactivalo.com`,            // CATCHED - NO DNS REPLY (*1)
	`nq-col.website`,             // REPORTED
	`nq-colombiaonline.website`,  // NO DNS REPLY (*1)
	`nqi-pr0pls0r.com`,           // CATCHED/ALIVE (*1)
	`nqicolmbia.com/NEQUI`,       // REPORTED
	`nqipr0pulsor.com`,           // NO DNS REPLY (*1)
	`nqprepropulso.com`,          // CATCHED (*1)
	`nqpropulsa.com`,             // CATCHED
	`nqpropulsando.com`,          // CATCHED
	`nuevopropulsor.com`,         // REPORTED (*2)
	`onlineparati.com`,           // REPORTED (*3)
	`parati-nqui.com`,            // NO DNS REPLY (*1)
	`perfectoparti.com`,          // REPORTED (*??)
	`preadelanto.online`,         // NO DNS REPLY (*1)
	`prepropulsandonq.com`,       // CATCHED - NO DNS REPLY (*1)
	`prepropulneq.com`,           // NO DNS REPLY (*1)
	`prepropulnq.com`,            // CATCHED (*1)
	`prestamo-nequi.website`,     // NO DNS REPLY (*1)
	`propulahorrosneq.com`,       // REPORTED
	`propulcolombiano.com`,       // REPORTED (*1)
	`propulideas.com`,            // REPORTED (*1)
	`propulsandoneqpro.com`,      // REPORTED (*1)
	`propulsor-pre.website`,      // NO DNS REPLY (*1)
	`propulsoraprobados.website`, // NO DNS REPLY (*1)
	`siperpropcolombia.com`,      // REPORTED (*2)
	`solicitadesdeya.com`,        // REPORTED (*1)
	`51.107.8.147`,               // DNS CATCHED or DOWN? (*??)
}

func attempt(mirrorPath string) {
	h := common.ReqHandler{}
	h.UseJar(true)

	// h.SendGet("https://"+mirrorPath+"/NEQUI/3d/propulsor/nequi/neq.php", nil, nil, nil)

	// POST (+ Preflight!?) https://yousitesureonlineverification.com/recursos/namekusei.php
	// "namekusei.php" is not responding anymore to those requests...
	// h.SendJSON("https://yousitesureonlineverification.com/recursos/namekusei.php",
	// 	map[string]string{"cedula": common.GeneraNIPcolombia()},
	// 	map[string]string{
	// 		"Origin":  "https://" + mirrorPath + "",
	// 		"Referer": "https://" + mirrorPath + "/",
	// 	}, nil)

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
			"Host":             mirrorPath,
			"Accept":           "*/*",
			"Accept-Language":  "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3",
			"Accept-Encoding":  "gzip, deflate, br, zstd",
			"X-Requested-With": "XMLHttpRequest",
			"Origin":           "https://" + mirrorPath + "",
			"Referer":          "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/neq.php",
		}, nil)
	//     SET-COOKIE:
	// usuario, contrasena, registro, estado
	u, _ := url.Parse("https://" + mirrorPath + "")
	h.PrintCookies(u)

	awaitStatusChange := func() {
		status := ""
		common.TimedCall(3*time.Second, 90*time.Second, func() bool {
			// POST https://{mirrorPath}/NEQUI/3d/process2/estado.php
			//     NOTE: This request returns a "status" number as body, related to consultar_estado()
			// of https://{mirrorPath}/NEQUI/3d/propulsor/nequi/js/functions2.js
			// (NO PAYLOAD! - MUST BE REPEATED WHILE SAME REPLY BODY)
			updatedStatus := ""
			h.SendPostEncoded(
				"https://"+mirrorPath+"/NEQUI/3d/process2/estado.php", nil,
				map[string]string{
					"Host":             mirrorPath,
					"Accept":           "*/*",
					"Accept-Language":  "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3",
					"Accept-Encoding":  "gzip, deflate, br, zstd",
					"X-Requested-With": "XMLHttpRequest",
					"Origin":           "https://" + mirrorPath + "",
					"Referer":          "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/cargando.php",
				}, &updatedStatus)
			if status == "" {
				status = updatedStatus
			} else if status != updatedStatus {
				return false
			}
			return true
		})
	}

	awaitStatusChange()

	h.SendPostEncoded(
		"https://"+mirrorPath+"/NEQUI/3d/process2/pasoOTP.php",
		map[string]string{"otp": common.GeneraPin(6)},
		map[string]string{
			"Host":             mirrorPath,
			"Accept":           "*/*",
			"Accept-Language":  "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3",
			"Accept-Encoding":  "gzip, deflate, br, zstd",
			"X-Requested-With": "XMLHttpRequest",
			"Origin":           "https://" + mirrorPath + "",
			"Referer":          "https://" + mirrorPath + "/NEQUI/3d/propulsor/nequi/otp.php",
		}, nil)
	//     SET-COOKIE:
	// cdinamica
	h.PrintCookies(u)

	awaitStatusChange()
}
