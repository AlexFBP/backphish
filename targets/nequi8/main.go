package nequi8

import (
	"time"

	"github.com/AlexFBP/backphish/common"
)

type mirrorManager struct {
	common.TargetWithParams
}

var target *mirrorManager

func init() {
	target = &mirrorManager{}
	target.Prefix = "nq8"
	target.Description = "attack fake nequi8"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *mirrorManager) Handler(mirror string) {
	h := common.ReqHandler{}
	h.UseJar(true)
	d := newUser()

	// ip data comes originally from https://ipapi.co/json/

	m := mirrData(mirror)

	// (step 1) POST back1 (from step_1/step_1.html)
	// {"documentNumber":"78215248","fullName":"Camilo Urrutia","userIP":"191.104.154.101","city":"Bogotá","country":"Colombia"}
	m.send(&h, d.DataForStep(1, m.HasOption("step1")), true)
	common.RandDelayRange(10*time.Second, 30*time.Second)

	// (step 2) POST back2 (from step_1/step_1_loader --> step_1/neq.html)
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
	m.send(&h, d.DataForStep(2, false), false)
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// (step 3) POST back2 (from otp)
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
	m.send(&h, d.DataForStep(3, false), false)
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// (step 4) POST back2 (from loading --> dinamica)
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
	m.send(&h, d.DataForStep(4, !m.HasOption("no-step")), false)
	common.RandDelayRange(35*time.Second, 50*time.Second)

	// (step 5) POST back2 (from load --> dinamica2)
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
	m.send(&h, d.DataForStep(5, !m.HasOption("no-step")), false)
	common.RandDelayRange(35*time.Second, 50*time.Second)
}

func (t *mirrorManager) EstimateDuration() time.Duration {
	return 120 * time.Second
}

func (t *mirrorManager) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/prestamo.html")
	return state
}
