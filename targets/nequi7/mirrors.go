package nequi7

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	Chat, Tok string
}

func (m *mirrorData) sendToTelegram(h *common.ReqHandler, msg string) {
	h.SendJSON("https://api.telegram.org/bot"+m.Tok+"/sendMessage", common.TgMsg{
		ChatID: m.Chat,
		Text:   msg,
	}, nil, nil)
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Chat, d.Tok = v[1], v[2]
			break
		}
	}
	return
}

func getMirrorNames() (names []string) {
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}

// Mirrors. The comment depending on last checked state:
//   - (*1): Apparently Down by host provider
var mirrors = [][]string{
	{`cops.blob.core.windows.net/nequitepresta`, // (*1??)
		"-1002484241220", "7709093911:AAHlo3XrnsIV6hKcAHhijHLFdSnnLVkXIsY"},
	{`nequicop.blob.core.windows.net/listo`, // (*1??)
		"-1002177027036", "7974637279:AAE9xaNbLyWeaXFViW7aav77EmUhFJr0r9s"},
	{`nequitepresta.blob.core.windows.net/nequisalvavidas`, // (*1??)
		"-1002282106001", "7200967568:AAFJtTq9WBZ2Ayu0IAQ0BaXjsEH1heXTW00"},
	{`propulsor.blob.core.windows.net/nequicop`, // Reported (*1??)
		"-1002383703541", "7914478130:AAE2pCUs3ww9M9pKb1xAWfkQdXa8y7VFTSI"},
	{`propulsornequi.blob.core.windows.net/prestamonequi`, // (*1??)
		"-1002403437567", "7795679390:AAEO8dqypaHRLPfzOjr091RVum5UV1mcQKA"},
	{`propulsornequi.blob.core.windows.net/nequi10`, // (*1??)
		"-1002383251748", "8047026966:AAH2-41mYGwDCJLuOY5wpOOofKtj8coK0zI"},
	{`tunequi.blob.core.windows.net/propulsor`, // (*1??)
		"-1002337684537", "7755303993:AAFjlXcO1uXA3gOoIW8p6Js_lTh3SM6wDss"},
	{`prestamosnequi.blob.core.windows.net/personas10`, // (*1)
		"-1002186144829", "7677367624:AAHekmyn7jqt_MqZH18k-Uh8AxXIrJnl-nQ"},
	{`prestamosnequi.blob.core.windows.net/personas11`, // (*1)
		"-1002152992408", "7287252560:AAGi6eBp509QR-PWMdPx3hkojoifzw6ddpU"},
	{`nequisoluciones.blob.core.windows.net/prestamo2`, // (*1)
		"-1002408055533", "8023442730:AAG2MzlRmACy5ufoZAL7-bUeJLhIs_kqhMc"},
	{`prestamosnequi.blob.core.windows.net/personas4`, // (*1)
		"-1002471720604", "7574903007:AAG9msllZ75x73VD6_IRiBRaF0noTplGCZI"},
	{`nequisoluciones.blob.core.windows.net/prestamo3`, // (*1)
		"-1002451015783", "7607065641:AAHTc_jql9Au_GjT27Nf2MAUPd_tPq4oYi4"},
	{`prestamosnequi.blob.core.windows.net/personas8`, // (*1)
		"-1002445779361", "7531945755:AAFSjCYMEx8LMgEUsUiSj3QTzo0ZGp09wAY"},
	{`prestamosnequi.blob.core.windows.net/personas2`, // (*1)
		"-1002365728544", "8042146105:AAHOji-OFpFK1QTDwV_ZubK-ppsbHU5NFFg"},
	{`nequisolventa.blob.core.windows.net/nequipass3`, // (*1)
		"-1002467189534", "8112748636:AAH3SOP4de3e6wAEAlcN62VGKyewdYSCEEY"},
	{`beneficiosnequi.blob.core.windows.net/personas`, // (*1)
		"-1002367126378", "8008206371:AAGRMTLYt9iwwKNnmLB_2pwZzGM50m23RjA"},
	{`prestamosnequi.blob.core.windows.net/personas7`, // (*1)
		"-1002444541599", "7685508863:AAEl3vrdrgcLRtpZTc1RkZ3W6eb-w9PY5Pg"},
	{`prestamosnequi.blob.core.windows.net/personas6`, // (*1)
		"-1002314729656", "8053719754:AAHksyrsJsKpH9fA7CxFN1grD85sWyAIAjk"},
	{`nequisolventa.blob.core.windows.net/nequipass1`, // (*1)
		"-1002355362792", "7849311285:AAFX9q0Eqmrh1VHCjiJm2KK-mGk16XEqwd0"},
	{`nequionline.blob.core.windows.net/prestamo`, // (*1)
		"-1002264456092", "7660727941:AAGtaApnHXMijI56fINjBx4n5uj_ykh6LJM"},
	{`nequisolventa.blob.core.windows.net/nequipass`, // (*1)
		"-1002405841116", "7752567905:AAFP9cAFxbx8n537K-MZOW30KVbADBrceac"},
	{`beneficiosnequi.blob.core.windows.net/person`, // (*1)
		"-1002410326093", "7783373587:AAF0x_KnI8JXczJq_83kkIBU9zxx4Hck7do"},
	{`nequisoluciones.blob.core.windows.net/nequi`, // (*1)
		"-1002001733552", "6837930538:AAGXkOZtNHOdY_t6NunPKL_Ch-YgwMnObTw"},
	{`prestamosnequi.blob.core.windows.net/personas9`, // (*1)
		"-1002362805954", "7851607767:AAEuAHYLfRoxNi8WYn3d1ICzmbPFmL5UuGU"},
	{`nequisolucion.blob.core.windows.net/admins`, // Reported
		"-1002436596397", "7970385783:AAHXBPnt6RzDi0Hhf4W7XtlG1VZuuMMfVyw"},
	{`nequi.blob.core.windows.net/propulsores`, // Reported
		"-1002313093398", "7353581823:AAEoXWONcGo4Afr7eXDn8cDQHwiXukNmbMo"},
	{`nequi.blob.core.windows.net/nequipropulsores`, // Reported
		"-1002152992408", "7287252560:AAGi6eBp509QR-PWMdPx3hkojoifzw6ddpU"},
	{`nequi.blob.core.windows.net/nequi`, // Reported
		"-1002314729656", "8053719754:AAHksyrsJsKpH9fA7CxFN1grD85sWyAIAjk"},
	{`nequiaplicacion.blob.core.windows.net/prestamo7`, // Reported
		"-1002305342766", "8189972648:AAEbAXhsouCatLUDAA3q316If4eDExR6n6E"},
	{`nequiaplicacion.blob.core.windows.net/prestamo8`, // Reported
		"-1002393796118", "7884031460:AAFDXx6nn8YQ8ZPAjcuH89qJNkLS-EcnJBQ"},
	{`equiponequi.blob.core.windows.net/cops`, // Reported
		"-1002348169138", "6377789175:AAFvwG062q6raYYagisTTk3yV34ahzkzz_A"}, // "Encrypted"... sure hahaha
	{`nequipropulsor.blob.core.windows.net/cops`, // Reported
		"-1002426206600", "6656456134:AAGx7giD07Ier3ZqyVcVp8PQq8Q89e8qNZA"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/nequi`, // Reported
		"-1002480614709", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/nequis`, // Reported
		"-1002393280382", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequi.blob.core.windows.net/actividades`, // Reported
		"-1002264456092", "7660727941:AAGtaApnHXMijI56fINjBx4n5uj_ykh6LJM"},
	{`nequisolucion.blob.core.windows.net/admin`, // Reported
		"-1002374516758", "8102808415:AAGyZ_bwz6ohRwyqqIDZIlkx9vLejHMY5Vw"},
	{`nequipropulsores.blob.core.windows.net/colo`, // Reported
		"-1002475294234", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/cops`, // Reported
		"-1002427872991", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequisolucion.blob.core.windows.net/newapp`, // Reported
		"-1002429872948", "7767432535:AAFCVM1Gy2fa1zl-1ehPVUoVsleNdjb5flg"},
	{`nequipropulsores.blob.core.windows.net/col`, // Reported
		"-1002365397330", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequipropulsor.blob.core.windows.net/nequi`, // Reported
		"-1002306266613", "6634297645:AAEKwmfXidaya-E-v7aJtfbpUotJ2VlCu4s"},
	{`nequipropulsor.blob.core.windows.net/cop`, // Reported
		"-1002398860222", "6656456134:AAGx7giD07Ier3ZqyVcVp8PQq8Q89e8qNZA"}, // "Encrypted"... sure hahaha
	{`nequisolucion.blob.core.windows.net/app4`, // Reported
		"-1002357449097", "7901233529:AAFUXg-7j3p_WayRZIUxbzhhR8gdYkOKizo"},
	{`nequiaplicacion.blob.core.windows.net/promo`, // Reported
		"-1002397456533", "7001646923:AAE7nNoBS-TpDEDX_Fbhdh0tOvLTjyMyFC0"}, // "Encrypted"... sure hahaha
	{`nequisolucion.blob.core.windows.net/prestamonequi`, // Reported
		"-1002389802747", "8149293119:AAGdu6a1GZSolzrpReEgoQFp6jnU0jn9Ep0"},
	{`nequipropulsor.blob.core.windows.net/company`, // Reported
		"-1002367929542", "6656456134:AAGx7giD07Ier3ZqyVcVp8PQq8Q89e8qNZA"}, // "Encrypted"... sure hahaha
	{`nequiaplicacion.blob.core.windows.net/nequicolombia`, // Reported
		"-1002293172819", "7229207986:AAH1_DUTvhUdJiEC0w9AJlQ8-yCqymt1oUU"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/nequiss`, // Reported
		"-1002381013958", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	// {``, // ALIVE
	// 	"", ""},
}
