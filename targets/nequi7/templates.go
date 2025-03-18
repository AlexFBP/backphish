package nequi7

import (
	"fmt"

	"github.com/AlexFBP/backphish/common"
)

type usrData struct {
	AllNames, ID, IP, City, Country, Phone, Pin string
}

func newUser() (d usrData) {
	d.AllNames = common.GeneraNombresApellidosPersonaCombinadosCol(false)
	d.ID = common.GeneraNIPcolombia()
	d.IP = common.GeneraIP()
	d.City, _ = common.GeneraCiudadDeptoColombia()
	d.Country = "Colombia"
	d.Phone = common.AddSeparator(common.GeneraCelColombia(), 0, " ")
	d.Pin = common.GeneraPin(4)
	return
}

func (d *usrData) template1() string {
	return fmt.Sprintf(`
👁Nequi Paso 1👁
🆔Nombres: %s
🪪Cedula: %s
🌏IP: %s
🏙Ciudad: %s
🇨🇴País: %s`,
		d.AllNames, d.ID, d.IP, d.City, d.Country)
}

func (d *usrData) template1b(num uint8) (msg string) {
	begin, dyn, end := "", "", ""

	if num == 0 {
		begin = `
👤Nequi_Meta_Infinito👤`
		end = fmt.Sprintf(`🇨🇴Ciudad: %s, País: %s`, d.City, d.Country)
	} else { // > 0
		begin = `🤠Nequi_Meta_Infinito🤠`
		dyn = fmt.Sprintf(`⭐️Dinamica%d: %s
`, num, common.GeneraPin(6))
		end = `🇨🇴Ubicación: ` + d.City
	}

	return fmt.Sprintf(`%s
🆔Nombres: %s
🪪Cedula: %s
#️⃣Número: %s
🔐Clave: %s
%s🌏IP: %s
%s`,
		begin, d.AllNames, d.ID, d.Phone, d.Pin, dyn, d.IP, end)
}

func (d *usrData) template2(step uint8) (msg string) {
	switch step {

	case 1:
		// ⭐⭐Nequi Paso 1⭐⭐\n\n🪪Cedula: 54890215\n👤Nombres: Julian Cardenas
		return fmt.Sprintf("⭐⭐Nequi Paso 1⭐⭐\n\n🪪Cedula: %s\n👤Nombres: %s",
			d.ID, d.AllNames)

	case 2:
		// ⭐⭐Nequi paso 2⭐⭐\n\n🪪Cedula: 54890215\n👤Nombres: Julian Cardenas\n📱Número: 315 468 2158\n 🔒Clave: 2547
		return fmt.Sprintf("⭐⭐Nequi paso 2⭐⭐\n\n🪪Cedula: %s\n👤Nombres: %s\n📱Número: %s\n 🔒Clave: %s",
			d.ID, d.AllNames, d.Phone, d.Pin)

	case 3:
		// ⭐️⭐️Dinamica 1⭐️⭐️\n\n🪪Cedula: 54890215\n👤Nombres: Julian Cardenas\n📱Número: 315 468 2158\n 🔒Clave: 2547\n 🔒Clave Dinamina 1: 431377
		return fmt.Sprintf("⭐️⭐️Dinamica 1⭐️⭐️\n\n🪪Cedula: %s\n👤Nombres: %s\n📱Número: %s\n 🔒Clave: %s\n 🔒Clave Dinamina 1: %s",
			d.ID, d.AllNames, d.Phone, d.Pin, common.GeneraPin(6))

	case 4:
		// ⭐️⭐️Dinamica 2⭐️⭐️\n\n🪪Cedula: 54890215\n👤Nombres: Julian Cardenas\n📱Número: 315 468 2158\n 🔒Clave: 2547\n 🔒Clave Dinamina 2: 971968
		return fmt.Sprintf("⭐️⭐️Dinamica 2⭐️⭐️\n\n🪪Cedula: %s\n👤Nombres: %s\n📱Número: %s\n 🔒Clave: %s\n 🔒Clave Dinamina 2: %s",
			d.ID, d.AllNames, d.Phone, d.Pin, common.GeneraPin(6))

	case 5:
		// ⭐️⭐️Dinamica 3⭐️⭐️\n\n🪪Cedula: 54890215\n👤Nombres: Julian Cardenas\n📱Número: 315 468 2158\n 🔒Clave: 2547\n 🔒Clave Dinamina 3: 394826
		return fmt.Sprintf("⭐️⭐️Dinamica 3⭐️⭐️\n\n🪪Cedula: %s\n👤Nombres: %s\n📱Número: %s\n 🔒Clave: %s\n 🔒Clave Dinamina 3: %s",
			d.ID, d.AllNames, d.Phone, d.Pin, common.GeneraPin(6))
	}

	return
}
