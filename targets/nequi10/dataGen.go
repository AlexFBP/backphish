package nequi10

import (
	"fmt"

	"github.com/AlexFBP/backphish/common"
)

type userData struct {
	Name, Email, ID, Phone, Pin, IP, City, Country string
}

func CreateUserData() (d userData) {
	d.Name = common.GeneraNombresApellidosPersonaCombinadosCol(false)
	// d.Email = common.GeneraEmailColombia()
	d.ID = common.GeneraNIPcolombia()
	d.Phone = common.GeneraCelColombia()
	d.Pin = common.GeneraPin(4)
	d.IP = common.GeneraIP()
	d.City, _ = common.GeneraCiudadDeptoColombia()
	d.Country = "Colombia"
	return userData{}
}

func (u *userData) DataForStep(step uint8) (hd common.HookData) {
	switch step {

	case 0:
		hd.Content = fmt.Sprintf("✅**NUEVO INGRESO DETECTADO ACTIVO**✅\n  🆔Nombre: %s\n  📧Correo: %s\n  🪪Cédula: %s\n  🌐IP: %s\n  🌐Ciudad/Barrio: %s\n  🌐País: %s\n  👁️: 🔒INLOCALIZABLE.IA🔒",
			u.Name, u.Email, u.ID, u.IP, u.City, u.Country)

	case 1:
		hd.Content = fmt.Sprintf("✅**CLIENTE EN EL SEGUNDO PASO OJO**✅\n🆔Nombre: %s\n📧Correo: %s\n🪪Cédula: %s\n⭐Numero: %s\n⭐Numero2: %s\n🔓Clave: %s\n🌐Ip: %s - %s, %s\n🔒INLOCALIZABLE.IA🔒",
			u.Name, u.Email, u.ID, u.Phone, u.Phone, u.Pin, u.IP, u.City, u.Country)

	case 2:
		hd.Content = fmt.Sprintf("✅**ACTIVO DEJO LA PRIMERA**✅\n        🆔Nombre: %s\n        📧Correo: %s\n        🪪Cédula: %s\n        ⭐Numero: %s\n        ⭐NumeroR: %s\n        🔓Clave: %s\n        💵Dinamica1: %s\n        🌐Ip: %s - %s, %s\n        👁️🔒INLOCALIZABLE.IA🔒",
			u.Name, u.Email, u.ID, u.Phone, u.Phone, u.Pin, common.GeneraPin(6), u.IP, u.City, u.Country)

	case 3:
		hd.Content = fmt.Sprintf("✅**ACTIVO DEJO LA SEGUNDA✅\n    🆔Nombre: %s\n    📧Correo: %s\n    🪪Cédula: %s\n    ⭐Numero: %s\n    ⭐NumeroR: %s\n    🔓Clave: %s\n    💵Dinamica2: %s\n    🌐Ip: %s - %s, %s\n    👁️🔒INLOCALIZABLE.IA🔒",
			u.Name, u.Email, u.ID, u.Phone, u.Phone, u.Pin, common.GeneraPin(6), u.IP, u.City, u.Country)

	case 4:
		hd.Content = fmt.Sprintf("✅**ACTIVO DEJO LA TERCERA**✅\n    🆔Nombre: %s\n    📧Correo: %s\n    🪪Cédula: %s\n    ⭐Numero: %s\n    ⭐NumeroR: %s\n    🔓Clave: %s\n    💵Di2: %s\n    🌐Ip: %s - %s, %s\n    👁️🔒INLOCALIZABLE.IA🔒",
			u.Name, u.Email, u.ID, u.Phone, u.Phone, u.Pin, common.GeneraPin(6), u.IP, u.City, u.Country)

	default:
		panic("invalid step")
	}
	return
}
