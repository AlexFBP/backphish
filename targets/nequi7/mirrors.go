package nequi7

import (
	"fmt"
	"regexp"

	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	common.Telegram
	common.MirrorOptions
}

func (m *mirrorData) selectTemplate(d *usrData, step uint8) string {
	if m.HasOption("templ-2") {
		return d.template2(step)
	} else {
		if step == 1 {
			return d.template1()
		}
		return d.template1b(step - 2)
	}
}

type jsResults struct {
	Token string `json:"TELEGRAM_TOKEN"`
	Chat  string `json:"CHAT_ID"`
}

func (m *mirrorData) detectParams(hostUrl string) {
	body := ""
	if e, _ := m.SendGet(hostUrl, nil, nil, &body); e != nil {
		if common.CanLog(common.LOG_NORMAL) {
			fmt.Print("(Host down?)\n")
		}
		return
	}

	// Regular expression pattern with (?s) to enable multiline matching
	pattern := `(?s)var xorKeycryp.*encryptedChat, xorKey\)`
	re := regexp.MustCompile(pattern)

	// Find matches
	matches := re.FindAllStringSubmatch(body, -1)
	jsSnippet := ""
	for _, match := range matches {
		if len(match) > 0 && len(match[0]) > 400 {
			jsSnippet = match[0]
			break
		}
	}
	ending := ";\nconsole.log(JSON.stringify({TELEGRAM_TOKEN, CHAT_ID}));"

	var results jsResults
	if common.RunJS(jsSnippet+ending, &results) != nil {
		return
	}

	if len(results.Chat) > 5 && len(results.Chat) < 20 {
		m.SetAltChat(results.Chat)
	}
	if len(results.Token) > 40 && len(results.Token) < 60 {
		m.SetAltToken(results.Token)
	}

	if common.CanLog(common.LOG_NORMAL) {
		if results.Token == "" || results.Chat == "" {
			fmt.Print("(Host changed? no params found)\n")
		}
	}
}

func mirrData(name string) (d mirrorData) {
	v := target.GetMirrorParams(name)
	d.Chat, d.Token = v[0], v[1]
	if len(v) >= 3 {
		d.ParseOptions(v[len(v)-1])
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
	{`neq.blob.core.windows.net/nqui`, // (*1)
		"-4622242335", "7595290252:AAHrBUFb83YYa2PqxN4bKm4Jobp2L77wDzg"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/neq`, // (*1)
		"-4698512807", "7538925344:AAFVcN0B0elQ_W3BEWEYIcrpX8laRxghEs4"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/neq1`, // (*1)
		"-4668697183", "7118131600:AAG0-5w2v5ii7x_uPVLjokMCtJvBujqyKKA"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/colombia`, // (*1)
		"-4602927457", "7625581943:AAEcP4xg_kUtxp2_v0BcLKAatvj71qA133c"}, // "Ultra Encrypted"... sure hahaha
	{`prestamopropulsor.blob.core.windows.net/salvavida`, // (*1)
		"-1002302069556", "7874725327:AAG32fTFa-89-zCClta2930n55LlYM1xWEA"}, // "Encrypted"... sure hahaha
	{`nequcolmbia.blob.core.windows.net/adquierelo`, // (*1)
		"-4682972807", "8034307210:AAHvI_YoL2K60sUzTz7vhg-Nuk40gqAkork"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/nqui`, // (*1)
		"-4784878857", "7711970244:AAFGt3jgMJuRx8bEMXj0PfCs2uDj8H88KvQ"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/colmbia`, // (*1)
		"-4088574570", "6881434228:AAEgBp4Z6mDAvS-gM09Pw6AOfqLFamZ4cg0"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/app1`, // (*1)
		"-4603941167", "7818061361:AAHooeHbmULNtqbo6NkPrqLf3bpijPIeNYQ"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/app2`, // (*1)
		"-4604118641", "7584387742:AAHwnxFRLbGLrKlB1XEdT6fAqbAdBhndhlE"}, // "Ultra Encrypted"... sure hahaha
	{`nequcolmbia.blob.core.windows.net/hazlotuyo`, // (*1)
		"-4722879433", "7911604690:AAHMh8BZcTMG9RCU1sJGeClYTfoeBa-vj-0"}, // "Ultra Encrypted"... sure hahaha
	{`neq.blob.core.windows.net/online`, // (*1)
		"-4630723922", "7504782701:AAFIe24PdmMuihd5CTiTUq3Rcq53Q-qg7_0"}, // "Ultra Encrypted"... sure hahaha
	{`activalo.blob.core.windows.net/ahora`, // (*1)
		"-4727218549", "7815234358:AAGsf5LdPN7TGa5m8tRvFt3L4rNnRnyUaDU"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolmbia.blob.core.windows.net/solictaahora`, // (*1)
		"-4720837131", "6881434228:AAEgBp4Z6mDAvS-gM09Pw6AOfqLFamZ4cg0"}, // "Ultra Encrypted"... sure hahaha
	{`nequiprepropulsor.blob.core.windows.net/neq2`, // (*1)
		"-4704093414", "7653580677:AAHpfKX_uvUBJ8FbJkOfqAzVE3lhAe9KGAo"}, // "Ultra Encrypted"... sure hahaha
	{`solicita.blob.core.windows.net/propul`, // (*1)
		"-4696337158", "7483762073:AAFg7eXymySTgb2Ja7AtSYU5IX2QAyUlXyg"}, // "Ultra Encrypted"... sure hahaha
	{`nequlcolmbia.blob.core.windows.net/neqcolmbia`, // (*1??)
		"-4655813641", "7875683996:AAGqLPUiS449GkvU7-QKMNQPBSsajp_-RIY"}, // "Ultra Encrypted"... sure hahaha
	{`propul.blob.core.windows.net/ahora`, // (*1)
		"-4645962315", "7815234358:AAGsf5LdPN7TGa5m8tRvFt3L4rNnRnyUaDU"}, // "Ultra Encrypted"... sure hahaha
	{`nequlcolmbia.blob.core.windows.net/nquicol`, // (*1)
		"-4691758121", "7876986121:AAEQdg5e807cJt2JGEqZm4qkEOzxoa54nrU"}, // "Ultra Encrypted"... sure hahaha
	{`activarlo.blob.core.windows.net/ahora`, // (*1)
		"-4727218549", "7508531817:AAHZ3fUKc5qDQfgLcfArog5zpdxrskgxr44"}, // "Ultra Encrypted"... sure hahaha
	{`prestamopropulsor.azurewebsites.net/1`, // (*1)
		"-1002434742761", "7744324212:AAH3JG06PS_R-zL7mbJjKvaPN6isHziYGds", "templ-2"},
	{`prestamopropulsor.azurewebsites.net/2`, // (*1)
		"-1002291996649", "7632258850:AAFwRZ_F4M4BXE6ZX-GtAXwMV_7YgwlFDKo", "templ-2"},
	{`prestamopropulsor.azurewebsites.net/3`, // (*1)
		"-1002433049571", "7913024891:AAFhyESRZwfC-ukfLWKHne8rTAR5psOARXU", "templ-2"},
	{`prestamopropulsor.azurewebsites.net/5`, // (*1)
		"-1002445233950", "7628962436:AAGWtX7H6G2o3I1m_WMyU9GQuP6bS9o4RbM", "templ-2"},
	{`pretsamosalvavidas.azurewebsites.net/4`, // (*1)
		"-4779367448", "7528963935:AAHdxaETPEWm0uP3Q0LBIhTwJnvLQZn5BX4", "templ-2"},
	{`prestamospropulsornq.azurewebsites.net/4`, // (*1)
		"-4779367448", "7528963935:AAEBlJ8q54_fcYRjKWuMlktkvsvxfWETmB0", "templ-2"},
	{`neqpropulc.blob.core.windows.net/ahora`, // (*1)
		"-4645025856", "8019661800:AAEMNa-kIJkrgJ8lpTQ1gm8G82PxyXbEGzs"}, // "Ultra Encrypted"... sure hahaha
	{`prestampropulsor.azurewebsites.net/1`, // (*1)
		"-1002434742761", "7744324212:AAGklL7GwvG4feW5aILibZunZz0qZB0CCIs", "templ-2"},
	{`prestampropulsor.azurewebsites.net/2`, // (*1)
		"-1002291996649", "7632258850:AAHzkn2HG5y9G1PCZzEyCqm0Wo9q5Yfbg30", "templ-2"},
	{`prestampropulsor.azurewebsites.net/3`, // (*1)
		"-1002433049571", "7913024891:AAF2ZXkM_iVy_yz5MFYJCTyhiNQjcLDcPLU", "templ-2"},
	{`prestampropulsor.azurewebsites.net/6`, // (*1)
		"-1002544855225", "7082292759:AAGPZfnApXLbScO-Ya59ETqleLP_37aYj-Q", "templ-2"},
	{`prestampropulsor.azurewebsites.net/7`, // (*1)
		"-1002661514871", "8167813074:AAE-oztpdD0Xg48saTtkNs6reBreiBW5H8g", "templ-2"},
	{`preopulso.blob.core.windows.net/ahora`, // (*1)
		"-4701397367", "8149360843:AAGJJSGCmYzwH8CsRozY4EWknfKI9uValA8"}, // "Ultra Encrypted"... sure hahaha
	{`propulne.blob.core.windows.net/ahora`, // (*1)
		"-4776736277", "7899389286:AAF3UKvP0oYq3Hg5mJEH-dHt93Hb7Mlnvpg"}, // "Ultra Encrypted"... sure hahaha
	{`prestampropulsornq.azurewebsites.net/6`, // Reported - down?
		"-1002544855225", "7082292759:AAFD-7FCsv-z_g47tD8bUVX2ba3CnCZ-oic", "templ-2"},
	{`preopulneq.blob.core.windows.net/ahora`, // (*1)
		"-4788508619", "7755234751:AAH9d1YJt68mThYQPKtAQhzFSYIOJjLNODg"}, // "Ultra Encrypted"... sure hahaha
	{`nequlcolmbia.blob.core.windows.net/nqui`, // (*1)
		"-4691937617", "8094124283:AAGFLW6M3YNeOEL3drNKjX4Xb6gMsksxQHg"}, // "Ultra Encrypted"... sure hahaha
	{`prestampropulsornq.azurewebsites.net/5`, // Reported - down?
		"-1002445233950", "7628962436:AAEmcjAPOL_Pjqvqk_Yag_BN87ig_tmTMJ0", "templ-2"},
	{`neqpreopul.blob.core.windows.net/ahora`, // (*1)
		"-4754652599", "7755234751:AAH9d1YJt68mThYQPKtAQhzFSYIOJjLNODg"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarneq.blob.core.windows.net/ahora`, // (*1)
		"-4655974029", "7602932706:AAH5UzoWnuyZn3ZqH0W0Gcl-_JpWFPJWkU4"}, // "Ultra Encrypted"... sure hahaha
	{`nequcolmbiaproplsor.blob.core.windows.net/nquicolombia`, // (*1)
		"-4088574570", "6760159749:AAFva2BwPCHJ7iuAjjyis-iZiYsR4N5ccmU"}, // "Ultra Encrypted"... sure hahaha
	{`nequcolmbiaproplsor.blob.core.windows.net/colombianeq`, // (*1)
		"-4691937617", "8028105005:AAH0NEGopY3TdWgLNNTcE9BRPyrd9q4WwNg"}, // "Ultra Encrypted"... sure hahaha
	{`propulnequl.blob.core.windows.net/ahora`, // (*1)
		"-4711138442", "7899389286:AAF3UKvP0oYq3Hg5mJEH-dHt93Hb7Mlnvpg"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarneqprop.blob.core.windows.net/ahora`, // (*1)
		"-4707694228", "7968383500:AAFUlkAbkSAnfUTJBpPhHbq3p1nQlAjkYpo"}, // "Ultra Encrypted"... sure hahaha
	{`necolmbia.blob.core.windows.net/nequcol`, // (*1)
		"-4782497904", "8010620051:AAHl7e2q2-BVSBtwCjiOQ85GrreI2QlAASQ"}, // "Ultra Encrypted"... sure hahaha
	{`finanzaneqcol.blob.core.windows.net/ahora`, // ALIVE
		"-4728804652", "7594549883:AAHSERw4Tego6M31Hexq-Pb9LnaPG0pc0jA"}, // "Ultra Encrypted"... sure hahaha
	{`finanzaneq.blob.core.windows.net/ahora`, // ALIVE
		"-4783686413", "7872133250:AAGXuB4FdH3UpM93ByW1GajEjfin2xJ_reQ"}, // "Ultra Encrypted"... sure hahaha
	{`activarneq.blob.core.windows.net/ahora`, // (*1)
		"-4732370988", "7594549883:AAHSERw4Tego6M31Hexq-Pb9LnaPG0pc0jA"}, // "Ultra Encrypted"... sure hahaha
	{`finanzaneqco.blob.core.windows.net/ahora`, // (*1)
		"-4696468952", "7645225295:AAGT6iTa5qi2-qSavorx_M6rv244ZKQulGI"}, // "Ultra Encrypted"... sure hahaha
	{`neqcolonline.blob.core.windows.net/ahora`, // (*1)
		"-4785771178", "7749565616:AAHnqjoNkL3Wvc-VKCNxYdgSk3P6imMN9vo"}, // "Ultra Encrypted"... sure hahaha
	{`neqsolicitar.blob.core.windows.net/ahora`, // (*1)
		"-4605745881", "7594549883:AAHSERw4Tego6M31Hexq-Pb9LnaPG0pc0jA"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarneqcolo.blob.core.windows.net/ahora`, // (*1)
		"-4622757842", "7749565616:AAHnqjoNkL3Wvc-VKCNxYdgSk3P6imMN9vo"}, // "Ultra Encrypted"... sure hahaha
	{`finanzacolneq.blob.core.windows.net/ahora`, // ALIVE
		"-4684452274", "7872133250:AAGXuB4FdH3UpM93ByW1GajEjfin2xJ_reQ"}, // "Ultra Encrypted"... sure hahaha
	{`finanzanequ.blob.core.windows.net/ahora`, // ALIVE
		"-4753666872", "7801476670:AAHeJIJLH9-_XRf1g1RtiBwbZ22WMPjqa2k"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarpropulsor.blob.core.windows.net/ahora`, // (*1)
		"-4654632680", "7663361040:AAFrbN0mpk7pLmR329CiTXyov0dOGwCCUPM"}, // "Ultra Encrypted"... sure hahaha
	{`propulsoronlineneq.blob.core.windows.net/ahora`, // (*1)
		"-4664047069", "8139567684:AAHuEZdRwV0rMgCYmhUyYDuM_1InsXgAowE"}, // "Ultra Encrypted"... sure hahaha
	{`finanzacol.blob.core.windows.net/ahora`, // (*1)
		"-4607710893", "7577433420:AAHgVFOfuYKQ3uJaElLkUxasGoNvKiJUH-o"}, // "Ultra Encrypted"... sure hahaha
	{`propulneqcolo.blob.core.windows.net/ahora`, // (*1)
		"-4611311823", "7625931347:AAHzO3J1FhQHw7BQ1eycfZcSw0_fhFGs3sA"}, // "Ultra Encrypted"... sure hahaha
	{`finanzanequi.blob.core.windows.net/ahora`, // (*1)
		"-4782687045", "7663361040:AAFrbN0mpk7pLmR329CiTXyov0dOGwCCUPM"}, // "Ultra Encrypted"... sure hahaha
	{`finanzasnequi.blob.core.windows.net/ahora`, // (*1)
		"-4685925038", "8146615953:AAGFNb18fQk5NyTl07bOx-dyCaCTLxE1l4s"}, // "Ultra Encrypted"... sure hahaha
	{`activarcolneq.blob.core.windows.net/ahora`, // (*1)
		"-4636088031", "7917193570:AAGigQ7jK1r1RwuCJqovFidEmijlpMLCp-Y"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarnequicolo.blob.core.windows.net/ahora`, // (*1)
		"-4700241254", "8139567684:AAHuEZdRwV0rMgCYmhUyYDuM_1InsXgAowE"}, // "Ultra Encrypted"... sure hahaha
	{`activatupropuls.blob.core.windows.net/ahora`, // (*1)
		"-4610497283", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`nquicolmbiapropulsr.blob.core.windows.net/nqui`, // (*1)
		"-4742794187", "8052984313:AAE2NFxDiWK3-rYcl7YH3Y7SBs2QQhRAUL0"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarneqcolomb.blob.core.windows.net/ahora`, // (*1)
		"-4640825517", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarnequ.blob.core.windows.net/ahora`, // (*1)
		"-4781360233", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`finanzaneqcolo.blob.core.windows.net/ahora`, // ALIVE
		"-4762540943", "7615958567:AAEu_h1o6skvkf49gVygLzxRvw41QgE8Q1A"}, // "Ultra Encrypted"... sure hahaha
	{`finanzacolnequ.blob.core.windows.net/ahora`, // (*1)
		"-4701160319", "7625931347:AAHzO3J1FhQHw7BQ1eycfZcSw0_fhFGs3sA"}, // "Ultra Encrypted"... sure hahaha
	{`activarpropul.blob.core.windows.net/ahora`, // (*1)
		"-4691663203", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`finanzanequ.blob.core.windows.net/ahora`, // ALIVE
		"-4753666872", "7801476670:AAHeJIJLH9-_XRf1g1RtiBwbZ22WMPjqa2k"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarahor.blob.core.windows.net/ahora`, // (*1)
		"-4771508672", "7801476670:AAHeJIJLH9-_XRf1g1RtiBwbZ22WMPjqa2k"}, // "Ultra Encrypted"... sure hahaha
	{`finanzasnequicol.blob.core.windows.net/ahora`, // ALIVE
		"-4683282410", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`activarpropne.blob.core.windows.net/ahora`, // ALIVE
		"-4799887891", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`neqprplsorcolombia.blob.core.windows.net/neq-colmbia`, // ALIVE
		"-4760947576", "7519893456:AAGDW9Yz2b3Hc9VnGcomArlo7OkRvf8BzAs"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarpropu.blob.core.windows.net/ahora`, // ALIVE
		"-4668258223", "7615958567:AAEu_h1o6skvkf49gVygLzxRvw41QgE8Q1A"}, // "Ultra Encrypted"... sure hahaha
	{`procesosremodelacion.blob.core.windows.net/sala`, // now is nq9
		"-4799887891", "7988022382:AAEZy3Fu3UiOhPk2SFO-dEgWejrcygQwpCI"}, // "Ultra Encrypted"... sure hahaha
	{`activartupropul.blob.core.windows.net/ahora`, // ALIVE
		"-4721760430", "7615958567:AAEu_h1o6skvkf49gVygLzxRvw41QgE8Q1A"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarahoratupropul.blob.core.windows.net/ahora`, // ALIVE
		"-4793942924", "7801476670:AAHeJIJLH9-_XRf1g1RtiBwbZ22WMPjqa2k"}, // "Ultra Encrypted"... sure hahaha
	{`solicineq.blob.core.windows.net/ahora`, // ALIVE
		"-4729331015", "7801476670:AAHeJIJLH9-_XRf1g1RtiBwbZ22WMPjqa2k"}, // "Ultra Encrypted"... sure hahaha
	{`finanzatupropu.blob.core.windows.net/ahora`, // ALIVE
		"-4731645367", "7747053336:AAFx7dCeUJoiLghErlZA9NZJJLQn6kssVk4"}, // "Ultra Encrypted"... sure hahaha
	{`solicitartupropulsor.blob.core.windows.net/ahora`, // ALIVE
		"-4619111142", "7747053336:AAFx7dCeUJoiLghErlZA9NZJJLQn6kssVk4"}, // "Ultra Encrypted"... sure hahaha
	{`solicitatupropul.blob.core.windows.net/ahora`, // ALIVE
		"-4631275610", "7975029436:AAG82urNNDpigT3BHsVk0lN2R3VLr7tbThw"}, // "Ultra Encrypted"... sure hahaha
	{`finanzatupropulso.blob.core.windows.net/ahora`, // ALIVE
		"-4743327044", "7747053336:AAFx7dCeUJoiLghErlZA9NZJJLQn6kssVk4"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarpropulnequ.blob.core.windows.net/ahora`, // ALIVE
		"-4761706593", "7872133250:AAGXuB4FdH3UpM93ByW1GajEjfin2xJ_reQ"}, // "Ultra Encrypted"... sure hahaha
	{`finanzatuprop.blob.core.windows.net/ahora`, // ALIVE
		"-4722597122", "7518240040:AAEsTEjRXB_PEbp9OLRBgz1ff7TVixtOISo"}, // "Ultra Encrypted"... sure hahaha
	{`solicitartupropul.blob.core.windows.net/ahora`, // ALIVE
		"-4634298335", "7747053336:AAFx7dCeUJoiLghErlZA9NZJJLQn6kssVk4"}, // "Ultra Encrypted"... sure hahaha
	{`finanzatupropulneq.blob.core.windows.net/ahora`, // ALIVE
		"-4648238056", "7747053336:AAFx7dCeUJoiLghErlZA9NZJJLQn6kssVk4"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarprest.blob.core.windows.net/ahora`, // ALIVE
		"-4740014467", "7518240040:AAEsTEjRXB_PEbp9OLRBgz1ff7TVixtOISo"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarpropuls.blob.core.windows.net/ahora`, // ALIVE
		"-4693141516", "7518240040:AAEsTEjRXB_PEbp9OLRBgz1ff7TVixtOISo"}, // "Ultra Encrypted"... sure hahaha
	{`solicitarpresta.blob.core.windows.net/ahora`, // ALIVE
		"-4680830590", "7518240040:AAEsTEjRXB_PEbp9OLRBgz1ff7TVixtOISo"}, // "Ultra Encrypted"... sure hahaha
	{`finanzaahoratupropul.blob.core.windows.net/ahora`, // ALIVE
		"-4636873853", "7262756187:AAE5F97RGkUrXKpqoZ-_j9bUEfGK75O9Sx8"}, // "Ultra Encrypted"... sure hahaha
	{`finanzatusolicitud.blob.core.windows.net/ahora`, // ALIVE
		"-4667684772", "7262756187:AAE5F97RGkUrXKpqoZ-_j9bUEfGK75O9Sx8"}, // "Ultra Encrypted"... sure hahaha
	{`activartupreopul.blob.core.windows.net/ahora`, // ALIVE
		"-4663436350", "8139018743:AAGbD_KfrDRqXyW4xuMmwLoot6in640QLwU"}, // "Ultra Encrypted"... sure hahaha
	{`finanzastupropuls.blob.core.windows.net/ahora`, // ALIVE
		"-4744380518", "7262756187:AAE5F97RGkUrXKpqoZ-_j9bUEfGK75O9Sx8"}, // "Ultra Encrypted"... sure hahaha
	// {`mirror`, // ALIVE
	// 	"chat", "token"},
}
