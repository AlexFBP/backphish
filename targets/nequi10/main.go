package nequi10

import (
	"time"

	"github.com/AlexFBP/backphish/common"
)

type TargetStr struct {
	common.TargetWithParams
}

var target *TargetStr

func init() {
	target = &TargetStr{}
	target.Prefix = "nq10"
	target.Description = "attack fake nequi10"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *TargetStr) Handler(mirror string) {
	// from index.html
	// POST {"content":"✅**CLIENTE EN EL SEGUNDO PASO OJO**✅\n🆔Nombre: Luis Bejarano\n📧Correo: lbeja24432@yahoo.es\n🪪Cédula: 34025878\n⭐Numero: 3186548912\n⭐Numero2: 3186548912\n🔓Clave: 3246\n🌐Ip: 152.203.49.131 - Bogotá, Colombia\n🔒INLOCALIZABLE.IA🔒"}

	// from prestamo.php.html
	// POST {"content":"✅**ACTIVO DEJO LA PRIMERA**✅\n        🆔Nombre: Luis Bejarano\n        📧Correo: lbeja24432@yahoo.es\n        🪪Cédula: 34025878\n        ⭐Numero: 3186548912\n        ⭐NumeroR: 3186548912\n        🔓Clave: 3246\n        💵Dinamica1: 947169\n        🌐Ip: 152.203.49.131 - Bogotá, Colombia\n        👁️🔒INLOCALIZABLE.IA🔒"}

	// from error_dinamica.php.html
	// POST {"content":"✅**ACTIVO DEJO LA SEGUNDA✅\n    🆔Nombre: Luis Bejarano\n    📧Correo: lbeja24432@yahoo.es\n    🪪Cédula: 34025878\n    ⭐Numero: 3186548912\n    ⭐NumeroR: 3186548912\n    🔓Clave: 3246\n    💵Dinamica2: 679751\n    🌐Ip: 152.203.49.131 - Bogotá, Colombia\n    👁️🔒INLOCALIZABLE.IA🔒"}

	// from error-otp.html
	// POST {"content":"✅**ACTIVO DEJO LA TERCERA**✅\n    🆔Nombre: Luis Bejarano\n    📧Correo: lbeja24432@yahoo.es\n    🪪Cédula: 34025878\n    ⭐Numero: 3186548912\n    ⭐NumeroR: 3186548912\n    🔓Clave: 3246\n    💵Di2: 062183\n    🌐Ip: 152.203.49.131 - Bogotá, Colombia\n    👁️🔒INLOCALIZABLE.IA🔒"}
}

func (t *TargetStr) EstimateDuration() time.Duration {
	return 0
}
