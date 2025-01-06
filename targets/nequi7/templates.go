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
	d.Phone = common.GeneraCelColombia()
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

func (d *usrData) template2(num uint8) (msg string) {
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
