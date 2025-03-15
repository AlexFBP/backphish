package nequi7

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	common.Telegram
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Chat, d.Token = v[1], v[2]
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
	{`nequisolucion.blob.core.windows.net/admins`, // (*1)
		"-1002436596397", "7970385783:AAHXBPnt6RzDi0Hhf4W7XtlG1VZuuMMfVyw"},
	{`nequi.blob.core.windows.net/propulsores`, // (*1)
		"-1002313093398", "7353581823:AAEoXWONcGo4Afr7eXDn8cDQHwiXukNmbMo"},
	{`nequi.blob.core.windows.net/nequipropulsores`, // (*1)
		"-1002152992408", "7287252560:AAGi6eBp509QR-PWMdPx3hkojoifzw6ddpU"},
	{`nequi.blob.core.windows.net/nequi`, // (*1)
		"-1002314729656", "8053719754:AAHksyrsJsKpH9fA7CxFN1grD85sWyAIAjk"},
	{`nequiaplicacion.blob.core.windows.net/prestamo7`, // (*1)
		"-1002305342766", "8189972648:AAEbAXhsouCatLUDAA3q316If4eDExR6n6E"},
	{`nequiaplicacion.blob.core.windows.net/prestamo8`, // (*1)
		"-1002393796118", "7884031460:AAFDXx6nn8YQ8ZPAjcuH89qJNkLS-EcnJBQ"},
	{`equiponequi.blob.core.windows.net/cops`, // (*1)
		"-1002348169138", "6377789175:AAFvwG062q6raYYagisTTk3yV34ahzkzz_A"}, // "Encrypted"... sure hahaha
	{`nequipropulsor.blob.core.windows.net/cops`, // (*1)
		"-1002426206600", "6656456134:AAGx7giD07Ier3ZqyVcVp8PQq8Q89e8qNZA"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/nequi`, // (*1)
		"-1002480614709", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/nequis`, // (*1)
		"-1002393280382", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequi.blob.core.windows.net/actividades`, // (*1)
		"-1002264456092", "7660727941:AAGtaApnHXMijI56fINjBx4n5uj_ykh6LJM"},
	{`nequisolucion.blob.core.windows.net/admin`, // (*1)
		"-1002374516758", "8102808415:AAGyZ_bwz6ohRwyqqIDZIlkx9vLejHMY5Vw"},
	{`nequipropulsores.blob.core.windows.net/colo`, // (*1)
		"-1002475294234", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/cops`, // (*1)
		"-1002427872991", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequisolucion.blob.core.windows.net/newapp`, // (*1)
		"-1002429872948", "7767432535:AAFCVM1Gy2fa1zl-1ehPVUoVsleNdjb5flg"},
	{`nequipropulsores.blob.core.windows.net/col`, // (*1)
		"-1002365397330", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`nequipropulsor.blob.core.windows.net/nequi`, // (*1)
		"-1002306266613", "6634297645:AAEKwmfXidaya-E-v7aJtfbpUotJ2VlCu4s"},
	{`nequipropulsor.blob.core.windows.net/cop`, // (*1)
		"-1002398860222", "6656456134:AAGx7giD07Ier3ZqyVcVp8PQq8Q89e8qNZA"}, // "Encrypted"... sure hahaha
	{`nequisolucion.blob.core.windows.net/app4`, // (*1)
		"-1002357449097", "7901233529:AAFUXg-7j3p_WayRZIUxbzhhR8gdYkOKizo"},
	{`nequiaplicacion.blob.core.windows.net/promo`, // (*1)
		"-1002397456533", "7001646923:AAE7nNoBS-TpDEDX_Fbhdh0tOvLTjyMyFC0"}, // "Encrypted"... sure hahaha
	{`nequisolucion.blob.core.windows.net/prestamonequi`, // (*1)
		"-1002389802747", "8149293119:AAGdu6a1GZSolzrpReEgoQFp6jnU0jn9Ep0"},
	{`nequipropulsor.blob.core.windows.net/company`, // (*1)
		"-1002367929542", "6656456134:AAGx7giD07Ier3ZqyVcVp8PQq8Q89e8qNZA"}, // "Encrypted"... sure hahaha
	{`nequiaplicacion.blob.core.windows.net/nequicolombia`, // (*1)
		"-1002293172819", "7229207986:AAH1_DUTvhUdJiEC0w9AJlQ8-yCqymt1oUU"}, // "Encrypted"... sure hahaha
	{`nequipropulsores.blob.core.windows.net/nequiss`, // (*1)
		"-1002381013958", "7069977751:AAHd9fcmrgN4_GBORSQTsrXSAa1xv9U0nXs"}, // "Encrypted"... sure hahaha
	{`equiponequi.blob.core.windows.net/nequis`, // (*1)
		"-1002439624246", "7525007091:AAGRCYm-_lz0xhV4N_MGSKx3FfDmHGowwbc"}, // "Encrypted"... sure hahaha
	{`nequiapp.blob.core.windows.net/salvavidas2`, // (*1)
		"-1002430384625", "7874725327:AAG32fTFa-89-zCClta2930n55LlYM1xWEA"}, // "Encrypted"... sure hahaha
	{`nequipropulsors.blob.core.windows.net/nequiplus`, // (*1)
		"-1002304685402", "8141847441:AAGCTjvs3qgEgikFa1laUTOX4eqR6Il1EhE"}, // "Encrypted"... sure hahaha
	{`nequipropulsors.blob.core.windows.net/nequipres`, // (*1)
		"-1002434807569", "8141847441:AAGCTjvs3qgEgikFa1laUTOX4eqR6Il1EhE"}, // "Encrypted"... sure hahaha
	{`nequiapp.blob.core.windows.net/neq5`, // (*1)
		"-1002230908332", "7702964113:AAHKU83E6Gjj4714YGNhNozD0AoZbJDLvIQ"}, // "Encrypted"... sure hahaha
	{`nequiapp.blob.core.windows.net/neq`, // (*1)
		"-1002449941642", "7702964113:AAHKU83E6Gjj4714YGNhNozD0AoZbJDLvIQ"}, // "Encrypted"... sure hahaha
	{`nequiapp.blob.core.windows.net/neq3`, // (*1)
		"-1002485074024", "7702964113:AAHKU83E6Gjj4714YGNhNozD0AoZbJDLvIQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps1`, // (*1)
		"-1002437725741", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps11`, // (*1)
		"-1002316491522", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps6`, // (*1)
		"-1002479593161", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps8`, // (*1)
		"-1002291237067", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps2`, // (*1)
		"-1002322450297", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps12`, // (*1)
		"-1002420827524", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`tuappnequi.blob.core.windows.net/apps18`, // (*1)
		"-1002337338043", "7349424091:AAFJc5GIQapRF6Ar4AfrfLVGsmlAPvWu4qQ"}, // "Encrypted"... sure hahaha
	{`nequiprepropulsor.blob.core.windows.net/neq`, // (*1)
		"-1002455200156", "8020269665:AAGGCs-wc9L91g6w24_pUgvfxOrdYEPvv-c"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/nqui`, // Reported
		"-4622242335", "7595290252:AAHrBUFb83YYa2PqxN4bKm4Jobp2L77wDzg"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/neq`, // (*1)
		"-4698512807", "7538925344:AAFVcN0B0elQ_W3BEWEYIcrpX8laRxghEs4"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/neq1`, // (*1)
		"-4668697183", "7118131600:AAG0-5w2v5ii7x_uPVLjokMCtJvBujqyKKA"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/colombia`, // Reported
		"-4602927457", "7625581943:AAEcP4xg_kUtxp2_v0BcLKAatvj71qA133c"}, // "Ultra Encrypted"... sure hahaha
	{`prestamopropulsor.blob.core.windows.net/salvavida`, // (*1)
		"-1002302069556", "7874725327:AAG32fTFa-89-zCClta2930n55LlYM1xWEA"}, // "Encrypted"... sure hahaha
	{`nequcolmbia.blob.core.windows.net/adquierelo`, // Reported
		"-4682972807", "8034307210:AAHvI_YoL2K60sUzTz7vhg-Nuk40gqAkork"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/nqui`, // Reported
		"-4784878857", "7711970244:AAFGt3jgMJuRx8bEMXj0PfCs2uDj8H88KvQ"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/colmbia`, // Reported
		"-4088574570", "6881434228:AAEgBp4Z6mDAvS-gM09Pw6AOfqLFamZ4cg0"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/app1`, // Reported
		"-4603941167", "7818061361:AAHooeHbmULNtqbo6NkPrqLf3bpijPIeNYQ"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/app2`, // Reported
		"-4604118641", "7584387742:AAHwnxFRLbGLrKlB1XEdT6fAqbAdBhndhlE"}, // "Ultra Encrypted"... sure hahaha
	{`nequcolmbia.blob.core.windows.net/hazlotuyo`, // Reported
		"-4722879433", "7911604690:AAHMh8BZcTMG9RCU1sJGeClYTfoeBa-vj-0"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/online`, // Reported
		"-4630723922", "7504782701:AAFIe24PdmMuihd5CTiTUq3Rcq53Q-qg7_0"}, // "Ultra Encrypted"... sure hahaha
	{`activalo.blob.core.windows.net/ahora`, // Reported
		"-4727218549", "7815234358:AAGsf5LdPN7TGa5m8tRvFt3L4rNnRnyUaDU"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/solictaahora`, // Reported
		"-4720837131", "6881434228:AAEgBp4Z6mDAvS-gM09Pw6AOfqLFamZ4cg0"}, // "Ultra Encrypted"... sure hahaha
	{`nequiprepropulsor.blob.core.windows.net/neq2`, // Reported
		"-4704093414", "7653580677:AAHpfKX_uvUBJ8FbJkOfqAzVE3lhAe9KGAo"}, // "Ultra Encrypted"... sure hahaha
	{`solicita.blob.core.windows.net/propul`, // (*1)
		"-4696337158", "7483762073:AAFg7eXymySTgb2Ja7AtSYU5IX2QAyUlXyg"}, // "Ultra Encrypted"... sure hahaha
	{`nequlcolmbia.blob.core.windows.net/neqcolmbia`, // Reported
		"-4714570940", "7900216272:AAHI3D_S0Lnv1OZ9YoBCXkyxsNt0beKEXR0"}, // "Ultra Encrypted"... sure hahaha
	// {`mirror`, // ALIVE
	// 	"chat", "token"},
}
