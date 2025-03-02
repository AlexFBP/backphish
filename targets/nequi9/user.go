package nequi9

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AlexFBP/backphish/common"
)

type hookData struct {
	User    string `json:"username"`
	Content string `json:"content"`
}

type userData struct {
	Nip      string
	FullName string
	Phone    string
	Pass     string
	IP       string
}

func newUser() (d userData) {
	d.Nip = common.GeneraNIPcolombia()
	d.FullName = common.GeneraNombresApellidosPersonaCombinadosCol(false)
	d.Phone = common.AddSeparator(common.GeneraCelColombia(), 0, " ")
	d.Pass = common.GeneraPin(4)
	d.IP = common.GeneraIP()
	return
}

func (d *userData) DataForStep(step uint8) (hd hookData) {
	switch step {

	case 0:
		// should be the same returned in JS by `new Date().toLocaleString()`
		// as example, "27/2/2025, 19:51:00"
		format := "02/01/2006, 3:04:05"
		// Assume a probability of user having 24h clock
		Probability24h := float32(0.2)
		if rand.Float32() > Probability24h {
			format += " PM"
		}
		// see https://pkg.go.dev/time#pkg-constants
		date := time.Now().Format(format)

		hd.Content = fmt.Sprintf(
			"🔔 *NUEVA VISITA NEQUI* 🔔\n\n📅 *Fecha:* %s\n🌍 *IP:* %s\n🚧 *Proveedor de Red:* No disponible",
			date, d.IP)

	case 1:
		hd.User = "Consulta de Crédito"
		hd.Content = fmt.Sprintf(
			"Nueva consulta recibida:\n\n📛Nombre:📛 %s\n📛Cédula:📛 %s\n",
			d.FullName, d.Nip)

	case 2:
		hd.User = "Consulta de Usuario"
		hd.Content = fmt.Sprintf(
			"Nueva solicitud recibida:\n\n**Nombre:** %s\n**Cédula:** %s\n**Número de Usuario:** %s\n**Clave:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass)

	case 3:
		hd.User = "Confirmación de OTP"
		hd.Content = fmt.Sprintf(
			"Nueva confirmación OTP recibida:\n\n**Nombre:** %s\n**Cédula:** %s\n**Número de Usuario:** %s\n**Clave:** %s\n**Código OTP:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass, common.GeneraPin(6))

	case 4:
		hd.User = "Confirmación Dinámica 2"
		hd.Content = fmt.Sprintf(
			"Confirmación dinámica 2 recibida:\n\n**Nombre:** %s\n**Cédula:** %s\n**Número:** %s\n**Clave:** %s\n**Código Dinámico 2:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass, common.GeneraPin(6))

	case 5:
		hd.User = "Confirmación Dinámica 3"
		hd.Content = fmt.Sprintf(
			"Confirmación dinámica 3 recibida:\n\n**Nombre:** %s\n**Cédula:** %s\n**Número:** %s\n**Clave:** %s\n**Código Dinámico 3:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass, common.GeneraPin(6))
	}
	return
}
