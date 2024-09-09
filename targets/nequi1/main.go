package nequi1

import (
	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt, common.ArgsHaveTimes(args...))
}

func attempt() {
	// POST (+ Preflight!?) https://yousitesureonlineverification.com/recursos/namekusei.php
	// content-type: application/json
	// origin: https://finanzasaturitmo.com
	// referer: https://finanzasaturitmo.com/
	// {"cedula":"23234654"}

	// POST https://finanzasaturitmo.com/NEQUI/3d/process2/pasousuario.php
	//     HEADERS:
	// content-type: application/x-www-form-urlencoded; charset=UTF-8
	// origin: https://finanzasaturitmo.com
	// referer: https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/neq.php
	//     VALUES:
	// pass: 3075
	// user: 310 350 5050
	// dis: PC
	// banco: NEQUI
	//     SET-COOKIE:
	// usuario, contrasena, registro, estado

	// POST https://finanzasaturitmo.com/NEQUI/3d/process2/estado.php
	//     HEADERS:
	// Cookie: usuario=310%20350%205050; contrasena=3075; registro=516; estado=1
	// content-type: application/json
	// origin: https://finanzasaturitmo.com
	// referer: https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/neq.php
	//     NOTE: This request returns a "status" number as body, related to consultar_estado()
	// of https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/js/functions2.js
	// (NO PAYLOAD! - MUST BE REPEATED WHILE REPLY BODY = "1" - should break with 2)

	// POST https://finanzasaturitmo.com/NEQUI/3d/process2/pasoOTP.php
	//     HEADERS:
	// content-type: application/x-www-form-urlencoded; charset=UTF-8
	// Referer:	https://finanzasaturitmo.com/NEQUI/3d/propulsor/nequi/otp.php
	// 			VALUES:
	// otp=726548
}
