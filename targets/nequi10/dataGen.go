package nequi10

import "github.com/AlexFBP/backphish/common"

type userData struct {
	Name, Email, ID, Phone, Pin, IP, City, Country string
}

func CreateUserData() (d userData) {
	return userData{}
}

func (u *userData) DataForStep(step uint8) (hd common.HookData) {
	switch step {

	case 1:
		// {"content":"✅**CLIENTE EN EL SEGUNDO PASO OJO**✅\n🆔Nombre: Luis Bejarano\n📧Correo: lbeja24432@yahoo.es\n🪪Cédula: 34025878\n⭐Numero: 3186548912\n⭐Numero2: 3186548912\n🔓Clave: 3246\n🌐Ip: 152.203.49.131 - Bogotá, Colombia\n🔒INLOCALIZABLE.IA🔒"}
		hd.Content = "Nequi paso 1"

	case 2:
		// {"content":"✅**ACTIVO DEJO LA PRIMERA**✅\n        🆔Nombre: Luis Bejarano\n        📧Correo: lbeja24432@yahoo.es\n        🪪Cédula: 34025878\n        ⭐Numero: 3186548912\n        ⭐NumeroR: 3186548912\n        🔓Clave: 3246\n        💵Dinamica1: 947169\n        🌐Ip: 152.203.49.131 - Bogotá, Colombia\n        👁️🔒INLOCALIZABLE.IA🔒"}
		hd.Content = "Nequi paso 2"

	case 3:
		// {"content":"✅**ACTIVO DEJO LA SEGUNDA✅\n    🆔Nombre: Luis Bejarano\n    📧Correo: lbeja24432@yahoo.es\n    🪪Cédula: 34025878\n    ⭐Numero: 3186548912\n    ⭐NumeroR: 3186548912\n    🔓Clave: 3246\n    💵Dinamica2: 679751\n    🌐Ip: 152.203.49.131 - Bogotá, Colombia\n    👁️🔒INLOCALIZABLE.IA🔒"}
		hd.Content = "Nequi paso 3"

	case 4:
		// {"content":"✅**ACTIVO DEJO LA TERCERA**✅\n    🆔Nombre: Luis Bejarano\n    📧Correo: lbeja24432@yahoo.es\n    🪪Cédula: 34025878\n    ⭐Numero: 3186548912\n    ⭐NumeroR: 3186548912\n    🔓Clave: 3246\n    💵Di2: 062183\n    🌐Ip: 152.203.49.131 - Bogotá, Colombia\n    👁️🔒INLOCALIZABLE.IA🔒"}

	default:
		panic("invalid step")
	}
	return
}
