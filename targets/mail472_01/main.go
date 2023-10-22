package mail47201

import "github.com/AlexFBP/backphish/common"

func Cmd(args ...string) error {
	attempt()
	return nil
}

func attempt() {
	// GET https://guianacional4-72.com/inicio.php

	// POST https://guianacional4-72.com/consultar.php

	// POST https://guianacional4-72.com/actualizar_datos.php

	common.SendPostEncoded(
		// "http://localhost:1080",
		"https://guianacional4-72.com/informacion_pago.php",
		map[string]string{
			"nombre":           "Juan",
			"apellido":         "Perez",
			"cedula":           "1000101010",
			"telefono":         "3131313131",
			"email":            "a@t.com",
			"ciudad":           "Cali",
			"direccion":        "Cali",
			"actualizar_datos": "Continuar",
		},
		map[string]string{
			"Origin":  "https://guianacional4-72.com",
			"Referer": "https://guianacional4-72.com/actualizar_datos.php",
		},
	)

	common.SendPostEncoded(
		// "http://localhost:1080",
		"https://guianacional4-72.com/comprobando.php",
		map[string]string{
			"codigo": "4573207634320360",
			"fecha":  "07/25",
			"cvv":    "509",
			"enviar": "Pagar Servicio",
		},
		map[string]string{
			"Origin":  "https://guianacional4-72.com",
			"Referer": "https://guianacional4-72.com/informacion_pago.php",
		},
	)
}
