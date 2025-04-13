package mail47201

import (
	"fmt"

	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "472"
	target.Description = "attack fake 472"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *mirrorTarget) Handler(base string) {
	h := common.ReqHandler{}
	// GET https://guianacional4-72.com/inicio.php

	// POST https://guianacional4-72.com/consultar.php

	// POST https://guianacional4-72.com/actualizar_datos.php

	pers := gofakeit.Person()
	// common.RandDelay(60, 60*5)
	h.SendPostEncoded(
		"https://"+base+"/informacion_pago.php",
		[]common.SimpleTerm{
			{K: "nombre", V: pers.FirstName},
			{K: "apellido", V: pers.LastName},
			{K: "cedula", V: fmt.Sprint(common.GeneraNIPcolombia())},
			{K: "telefono", V: pers.Contact.Phone},
			{K: "email", V: pers.Contact.Email},
			{K: "ciudad", V: pers.Address.City},
			{K: "direccion", V: pers.Address.Address},
			{K: "actualizar_datos", V: "Continuar"},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/actualizar_datos.php"},
		},
		nil,
	)

	// common.RandDelay(30, 80)
	h.SendPostEncoded(
		"https://"+base+"/comprobando.php",
		[]common.SimpleTerm{
			{K: "codigo", V: fmt.Sprint(pers.CreditCard.Number)},
			{K: "fecha", V: pers.CreditCard.Exp},
			{K: "cvv", V: pers.CreditCard.Cvv},
			{K: "enviar", V: "Pagar Servicio"},
		},
		[]common.SimpleTerm{
			{K: "Origin", V: "https://" + base},
			{K: "Referer", V: "https://" + base + "/informacion_pago.php"},
		},
		nil,
	)
}

var mirrors = []string{
	"guianacional4-72.com", // DOWN
}
