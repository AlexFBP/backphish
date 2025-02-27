package nequi9

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq9",
		Description: "attack fake nequi9",
		Mirrors:     getMirrorNames(),
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(mirror string) {
	// h := common.ReqHandler{}
	// d := newUser()
	// m := mirrData(mirror)

	// (step 1) POST application/json webhook (from ini.html)
	/*
		{
			username: "Consulta de Crédito",
			content: `Nueva consulta recibida:\n\n📛Nombre:📛 ${nombreApellido}\n📛Cédula:📛 ${cedula}\n`
		}
	*/

	// (step 2) POST webhook (from neq.html)
	/*
		{
			username: "Consulta de Usuario",
			content: `Nueva solicitud recibida:\n\n` +
											`**Nombre:** ${nombreApellido}\n` +
											`**Cédula:** ${cedula}\n` +
											`**Número de Usuario:** ${usuario}\n` +
											`**Clave:** ${clave}`
		}
	*/

	// (step 3) POST webhook (from otp.html)
	/*
		{
			username: "Confirmación de OTP",
			content: `Nueva confirmación OTP recibida:\n\n` +
									`**Nombre:** ${nombreApellido}\n` +
									`**Cédula:** ${cedula}\n` +
									`**Número de Usuario:** ${numero}\n` +
									`**Clave:** ${clave}\n` +
									`**Código OTP:** ${otpCode}`
		}
	*/

	// (step 4) POST back2 (from loading --> dinamica)
	/*
		{
			username: "Confirmación Dinámica 2",
			content: `Confirmación dinámica 2 recibida:\n\n` +
									`**Nombre:** ${nombreApellido || "No especificado"}\n` +
									`**Cédula:** ${cedula || "No especificado"}\n` +
									`**Número:** ${numero || "No especificado"}\n` +
									`**Clave:** ${clave || "No especificado"}\n` +
									`**Código Dinámico 2:** ${otpCode2}`
		}
	*/

	// (step 5) POST back2 (from load --> dinamica2)
	/*
		{
			username: "Confirmación Dinámica 3",
			content: `Confirmación dinámica 3 recibida:\n\n` +
									`**Nombre:** ${nombreApellido || "No especificado"}\n` +
									`**Cédula:** ${cedula || "No especificado"}\n` +
									`**Número:** ${numero || "No especificado"}\n` +
									`**Clave:** ${clave || "No especificado"}\n` +
									`**Código Dinámico 3:** ${otpCode3}`
		}
	*/
}
