package nequi7

import (
	"fmt"
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq7",
		Description: "attack fake nequi7",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(GetAllCmds())
}

func GetAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

var mirrors = []string{
	`cops.blob.core.windows.net/nequitepresta`,            // ALIVE
	`nequicop.blob.core.windows.net/listo`,                // ALIVE
	`nequitepresta.blob.core.windows.net/nequisalvavidas`, // ALIVE
	`propulsor.blob.core.windows.net/nequicop`,            // Reported
	`propulsornequi.blob.core.windows.net/prestamonequi`,  // ALIVE
	`propulsornequi.blob.core.windows.net/nequi10`,        // ALIVE
	`tunequi.blob.core.windows.net/propulsor`,             // ALIVE
	`prestamosnequi.blob.core.windows.net/personas10`,     // ALIVE
	`prestamosnequi.blob.core.windows.net/personas11`,     // ALIVE
	`nequisoluciones.blob.core.windows.net/prestamo2`,     // ALIVE
	`prestamosnequi.blob.core.windows.net/personas4`,      // ALIVE
	`nequisoluciones.blob.core.windows.net/prestamo3`,     // ALIVE
	`prestamosnequi.blob.core.windows.net/personas8`,      // ALIVE
	`prestamosnequi.blob.core.windows.net/personas2`,      // ALIVE
	`nequisolventa.blob.core.windows.net/nequipass3`,      // ALIVE
	`beneficiosnequi.blob.core.windows.net/personas`,      // ALIVE
	`prestamosnequi.blob.core.windows.net/personas7`,      // ALIVE
	`prestamosnequi.blob.core.windows.net/personas6`,      // ALIVE
	`nequisolventa.blob.core.windows.net/nequipass1`,      // ALIVE
	`nequionline.blob.core.windows.net/prestamo`,          // ALIVE
	`nequisolventa.blob.core.windows.net/nequipass`,       // ALIVE
	`beneficiosnequi.blob.core.windows.net/person`,        // ALIVE
	`nequisoluciones.blob.core.windows.net/nequi`,         // ALIVE
	`prestamosnequi.blob.core.windows.net/personas9`,      // ALIVE
	// ``, // ALIVE
}

var mirrorDetails = [][3]string{
	{"-1002484241220", "7709093911:AAHlo3XrnsIV6hKcAHhijHLFdSnnLVkXIsY", "short"},
	{"-1002177027036", "7974637279:AAE9xaNbLyWeaXFViW7aav77EmUhFJr0r9s", "short"},
	{"-1002282106001", "7200967568:AAFJtTq9WBZ2Ayu0IAQ0BaXjsEH1heXTW00", "long"},
	{"-1002383703541", "7914478130:AAE2pCUs3ww9M9pKb1xAWfkQdXa8y7VFTSI", "short"},
	{"-1002403437567", "7795679390:AAEO8dqypaHRLPfzOjr091RVum5UV1mcQKA", "short"},
	{"-1002383251748", "8047026966:AAH2-41mYGwDCJLuOY5wpOOofKtj8coK0zI", "long"},
	{"-1002337684537", "7755303993:AAFjlXcO1uXA3gOoIW8p6Js_lTh3SM6wDss", "short"},
	{"-1002186144829", "7677367624:AAHekmyn7jqt_MqZH18k-Uh8AxXIrJnl-nQ", "alt1"},
	{"-1002152992408", "7287252560:AAGi6eBp509QR-PWMdPx3hkojoifzw6ddpU", "long"},
	{"-1002408055533", "8023442730:AAG2MzlRmACy5ufoZAL7-bUeJLhIs_kqhMc", "long"},
	{"-1002471720604", "7574903007:AAG9msllZ75x73VD6_IRiBRaF0noTplGCZI", "alt2"},
	{"-1002451015783", "7607065641:AAHTc_jql9Au_GjT27Nf2MAUPd_tPq4oYi4", "long"},
	{"-1002445779361", "7531945755:AAFSjCYMEx8LMgEUsUiSj3QTzo0ZGp09wAY", "alt2"},
	{"-1002365728544", "8042146105:AAHOji-OFpFK1QTDwV_ZubK-ppsbHU5NFFg", "alt2"},
	{"-1002467189534", "8112748636:AAH3SOP4de3e6wAEAlcN62VGKyewdYSCEEY", "alt2"},
	{"-1002367126378", "8008206371:AAGRMTLYt9iwwKNnmLB_2pwZzGM50m23RjA", "alt2"},
	{"-1002444541599", "7685508863:AAEl3vrdrgcLRtpZTc1RkZ3W6eb-w9PY5Pg", "alt2"},
	{"-1002314729656", "8053719754:AAHksyrsJsKpH9fA7CxFN1grD85sWyAIAjk", "alt2"},
	{"-1002355362792", "7849311285:AAFX9q0Eqmrh1VHCjiJm2KK-mGk16XEqwd0", "alt2"},
	{"-1002264456092", "7660727941:AAGtaApnHXMijI56fINjBx4n5uj_ykh6LJM", "alt2"},
	{"-1002405841116", "7752567905:AAFP9cAFxbx8n537K-MZOW30KVbADBrceac", "alt2"},
	{"-1002410326093", "7783373587:AAF0x_KnI8JXczJq_83kkIBU9zxx4Hck7do", "alt2"},
	{"-1002001733552", "6837930538:AAGXkOZtNHOdY_t6NunPKL_Ch-YgwMnObTw", "alt2"},
	{"-1002362805954", "7851607767:AAEuAHYLfRoxNi8WYn3d1ICzmbPFmL5UuGU", "alt2"},
	// {"", "", ""},
}

func attempt(mirror string) {
	h := common.ReqHandler{}

	det := mirrorDetails[common.FindPos(mirror, mirrors)]
	cel := common.GeneraCelColombia()
	pin := common.GeneraPin(4)
	ip := common.GeneraIP()

	city, _ := common.GeneraCiudadDeptoColombia()
	country := "Colombia"
	ced := common.GeneraNIPcolombia()
	nom := common.GeneraNombresApellidosPersonaCombinadosCol(false)

	long := ""
	if det[2] == "long" || det[2] == "alt1" || det[2] == "alt2" {
		full := ""
		if det[2] == "alt1" || det[2] == "alt2" {
			long = fmt.Sprintf("🆔Nombres: %s\n🪪Cedula: %s\n", nom, ced)
			full = fmt.Sprintf("👁Nequi Paso 1👁\n%s🌏IP: %s\n🏙Ciudad: %s\n🇨🇴País: %s",
				long, ip, city, country)
		} else {
			long = fmt.Sprintf("Nombres: %s\nCedula: %s\n", nom, ced)
			full = fmt.Sprintf("Nequi_Paso_1\n%sIP: %s\nCiudad: %s, País: %s",
				long, ip, city, country)
		}

		h.SendJSON("https://api.telegram.org/bot"+det[1]+"/sendMessage", common.TgMsg{
			ChatID: det[0],
			Text:   full,
		}, nil, nil)

		common.RandDelayRange(4*time.Second, 20*time.Second)
	}

	tfa := ""
	emoji1 := ""
	loc := ""
	if det[2] == "alt1" {
		tfa = common.GeneraPin(6)
		tfa = fmt.Sprintf("⭐️Dinamica1: %s\n", common.GeneraPin(6))
		loc = fmt.Sprintf("🇨🇴Ubicación: %s", city)
		emoji1 = "🤠"
	} else if det[2] == "alt2" {
		emoji1 = "👤"
		loc = fmt.Sprintf("🇨🇴Ciudad: %s, País: %s", city, country)
	}

	if det[2] == "alt1" || det[2] == "alt2" {
		h.SendJSON("https://api.telegram.org/bot"+det[1]+"/sendMessage", common.TgMsg{
			ChatID: det[0],
			Text: fmt.Sprintf(`%sNequi_Meta_Infinito%s
  🆔Nombres: %s
  🪪Cedula: %s
  #️⃣Número: %s
  🔐Clave: %s
  %s🌏IP: %s
  %s`,
				emoji1, emoji1, nom, ced, common.AddSeparator(cel, 0, " "), pin, tfa, ip, loc),
		}, nil, nil)

	} else if det[2] == "long" || det[2] == "short" {
		h.SendJSON("https://api.telegram.org/bot"+det[1]+"/sendMessage", common.TgMsg{
			ChatID: det[0],
			Text: fmt.Sprintf("Nequi_Meta_Infinito\n%sNúmero: %s\nClave: %s\nIP: %s\nCiudad: %s, País: %s",
				long, common.AddSeparator(cel, 0, " "), pin, ip, city, country),
		}, nil, nil)
	}
	common.RandDelayRange(1*time.Second, 2*time.Second)

}
