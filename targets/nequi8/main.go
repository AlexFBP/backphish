package nequi8

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq7",
		Description: "attack fake nequi7",
		Mirrors:     getMirrorNames(),
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(mirror string) {
	h := common.ReqHandler{}
	h.UseJar(true)
	var d msg

	// ip data comes originally from https://ipapi.co/json/

	m := mirrData(mirror)

	// POST back1 (from step_1/step_1.html)
	// {"documentNumber":"78215248","fullName":"Camilo Urrutia","userIP":"191.104.154.101","city":"Bogotá","country":"Colombia"}
	m.send(&h, d, true)

	// POST back2 (from step_1/step_1_loader --> step_1/neq.html)
	/*
		{
				documentNumber: cedula,
				fullName: nombreApellido,
				username: usuario,
				password: clave,
				userIP: ip,
				city: ciudad,
				country: pais,
		}
	*/
	m.send(&h, d, false)

	// POST back2 (from otp)
	/*
		{
				documentNumber: cedula,
				fullName: nombreApellido,
				username: numero,
				password: clave,
				userIP: ip,
				city: ipData.city || 'Desconocida',
				country: ipData.country_name || 'Desconocido',
				otpCode: otpCode
		}
	*/
	m.send(&h, d, false)

	// POST back2 (from loading --> dinamica)
	/*
		{
				documentNumber: cedula,
				fullName: nombreApellido,
				username: numero,
				password: clave,
				userIP: ip,
				city: data.city || 'Desconocida',
				country: data.country_name || 'Desconocido',
				dynamicCode: dynamicCode,
				step: 'dinamica2' // mirr1 does NOT have this field
		}
	*/
	m.send(&h, d, false)

	// POST back2 (from load --> dinamica2)
	/*
		{
				documentNumber: cedula,
				fullName: nombreApellido,
				username: numero,
				password: clave,
				userIP: ip,
				city: data.city || 'Desconocida',
				country: data.country_name || 'Desconocido',
				dynamicCode: dynamicCode,
				step: 'dinamica3' // mirr1 does NOT have this
		}
	*/
	m.send(&h, d, false)
}
