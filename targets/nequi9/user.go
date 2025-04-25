package nequi9

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AlexFBP/backphish/common"
)

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

func (d *userData) DataForStep(step uint8, options map[string]string) (hd common.HookData) {
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
		wrap := ""
		if v, ok := options["wrap"]; ok {
			wrap = v
		}
		if _, ok := options["alt4"]; ok {
			hd.User = "Consulta de Usuario"
			hd.Content = fmt.Sprintf(
				"📛 Nombre: %s\n🆔 Cédula: %s",
				d.FullName, d.Nip)
		} else {
			hd.Content = fmt.Sprintf(
				"Nueva consulta recibida:\n\n%s📛Nombre:📛 %s\n📛Cédula:📛%s %s\n",
				wrap, d.FullName, wrap, d.Nip)
		}

	case 2:
		pre := "Consulta de Usuario"
		if _, ok := options["alt2"]; !ok {
			if wrap, ok := options["wrap-ini"]; ok {
				pre = wrap + pre
			}
			if wrap, ok := options["wrap-end"]; ok {
				pre += wrap
			}
		}
		hd.User = pre
		pre = "Nueva solicitud recibida"
		if _, ok := options["alt2"]; ok {
			pre = "Consulta de Usuario"
			if wrap, ok := options["wrap-ini"]; ok {
				pre = wrap + pre + wrap
			}
		}
		hd.Content = fmt.Sprintf(
			pre+":\n\n**Nombre:** %s\n**Cédula:** %s\n**Número de Usuario:** %s\n**Clave:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass)

	case 3:
		pre := "Confirmación de OTP"
		if _, ok := options["usr1"]; ok {
			pre = "Consulta de Usuario"
		}
		if wrap, ok := options["wrap-ini"]; ok {
			pre = wrap + pre
		}
		if wrap, ok := options["wrap-end"]; ok {
			pre += wrap
		}
		hd.User = pre
		pre = "Nueva confirmación OTP recibida"
		if _, ok := options["cont1"]; ok {
			pre = "✅ACTIVO DEJO LA PRIMERA ✅"
		}
		hd.Content = fmt.Sprintf(
			pre+":\n\n**Nombre:** %s\n**Cédula:** %s\n**Número de Usuario:** %s\n**Clave:** %s\n**Código OTP:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass, common.GeneraPin(6))

	case 4:
		pre := "Confirmación Dinámica 2"
		if wrap, ok := options["wrap"]; ok {
			pre = wrap + pre + wrap
		}
		hd.User = pre
		pre = "Confirmación dinámica 2 recibida"
		if _, ok := options["cont1"]; ok {
			pre = "✅ACTIVO DEJO LA SEGUNDA✅"
		}
		hd.Content = fmt.Sprintf(
			pre+":\n\n**Nombre:** %s\n**Cédula:** %s\n**Número:** %s\n**Clave:** %s\n**Código Dinámico 2:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass, common.GeneraPin(6))

	case 5:
		pre := "Confirmación Dinámica 3"
		if wrap, ok := options["wrap"]; ok {
			pre = wrap + pre + wrap
		}
		hd.User = pre
		pre = "Confirmación dinámica 3 recibida"
		if _, ok := options["cont1"]; ok {
			pre = "✅ACTIVO DEJO LA TERCERA✅"
		}
		hd.Content = fmt.Sprintf(
			pre+":\n\n**Nombre:** %s\n**Cédula:** %s\n**Número:** %s\n**Clave:** %s\n**Código Dinámico 3:** %s",
			d.FullName, d.Nip, d.Phone, d.Pass, common.GeneraPin(6))
	}
	return
}
