package mail47201

import (
	"fmt"

	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt, common.ArgsHaveTimes(args...))
	// For single attack, comment line above and uncomment this:
	// attempt()
	// return nil
}

func attempt() {
	// GET https://guianacional4-72.com/inicio.php

	// POST https://guianacional4-72.com/consultar.php

	// POST https://guianacional4-72.com/actualizar_datos.php

	pers := gofakeit.Person()
	// common.RandDelay(60, 60*5)
	common.SendPostEncoded(
		"https://guianacional4-72.com/informacion_pago.php",
		map[string]string{
			"nombre":           pers.FirstName,
			"apellido":         pers.LastName,
			"cedula":           fmt.Sprint(common.GeneraNIPcolombia()),
			"telefono":         pers.Contact.Phone,
			"email":            pers.Contact.Email,
			"ciudad":           pers.Address.City,
			"direccion":        pers.Address.Address,
			"actualizar_datos": "Continuar",
		},
		map[string]string{
			"Origin":  "https://guianacional4-72.com",
			"Referer": "https://guianacional4-72.com/actualizar_datos.php",
		},
		nil,
	)

	// common.RandDelay(30, 80)
	common.SendPostEncoded(
		"https://guianacional4-72.com/comprobando.php",
		map[string]string{
			"codigo": fmt.Sprint(pers.CreditCard.Number),
			"fecha":  pers.CreditCard.Exp,
			"cvv":    pers.CreditCard.Cvv,
			"enviar": "Pagar Servicio",
		},
		map[string]string{
			"Origin":  "https://guianacional4-72.com",
			"Referer": "https://guianacional4-72.com/informacion_pago.php",
		},
		nil,
	)
}
