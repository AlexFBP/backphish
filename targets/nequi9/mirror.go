package nequi9

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	common.Discord
	common.Telegram
	common.MirrorOptions
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Webhook, d.Token, d.Chat = v[1], v[2], v[3]
			d.ParseOptions(v[4])
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
	{"sucursalvirtuall.blob.core.windows.net/nequi001", "1337922396681932920/gUxfz9zq-H068Qh0IEr6G2qncQ7XEN0E-Nev5Z1zJC4_PuwBHa_9rCll0RC0hyZkaRzX", "", "", ""},                                                             // (*1) reported -  "Z1xGSEFXVlxcQlQuJDUvCVAqOT0OFDcoVyUTVSAHEzYjJDYGCA4tNQU5GlECNA==", "fV5CQUdWVVtQQF5YUEo=", "Porqueblesnoesejemplodenadaaaaa"
	{"sucursalvirtuall.blob.core.windows.net/nequi002", "1338548943050637343/b_xUBPOsA4RwRmZfZZTSzJUXFZ_QDmEFOvYdhFtb9RrTGZ3ptAt9Fn69HIPFODQL0Llq", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi003", "1338558488074391605/-2wgbBayVSJDEWkkEjTaBQO7vR_wVRVCRINmB8My-216llzcPbL_4z6qa6SoNXuFqCPu", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi004", "1338559433109667893/-5gdmIXjr-eMPnL6_ONtIDKbXbPUQl2CrWok1s7EOb0ygEMi7cJxh71YIOqASWamCY3I", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi005", "1338560850050547744/uzAHiOnJ_t3pEvlGqorGxIwcPX49xTgTFG-GH-YYzEIeDg17LtTny8UPgbgGssn16yjA", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi006", "1338562843296923721/BCcezLTQBFgy6lgEEO077YAe4LGLXWTqdGWv3QLupG-gW0IXzHuHdO1hF34Ci1Ebdmm-", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi007", "1338563500209078303/RSOYlk4En5LCJQOagUn12jqsgyxAnqO-BTfpxDf2z2nPPlH0fCOztbu0m1Fhw5tHLbm9", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi008", "1338564040821309461/A2IHsI_A4DE9qsVSTw3Jrl5w-2nTszdbzJrqZAibOwBIem1KQYDGPzazDlxsutw1InXi", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi009", "1338566707609862235/KvS8FT3dGbEUvKH-bQGV2_xzVoGxKArj4aq39PsEA53ww2p8qYJ5srXgauApMa5WmR0Z", "", "", ""},                                                             // (*1) reported - (same hardcoded) - HOLD: TK1 TK2 TK2 TK2
	{"sucursalvirtuall.blob.core.windows.net/nequi010", "1338567600279523450/u--zHUjQZ5p9ZcnuUN5yKzSjAN4iAsaKEKB1E97rqUYhmbrQ058QwHW1c-WLy4Io-S58", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi011", "1338568450125336669/gMtBmDXuqzn-QA_9NkvM5YXZK9U1yB3kmZ4zpYQqzvac_3PqVKXX-n9dK_Lp_0GevYSp", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi012", "1338569266336895049/gIyNCUM92yX7B5_7od9WelHUFVrhMgLBfXOTtH6W9EPTkEIu1ortl5iPYqE6sMKoT3_D", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi013", "1338570085090201620/uCeggXpE5SJvYq6r_bsGeYVVTtycxZbAZf_E7IT4gX-fWgYode3LtfH545543m4iIdq6", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi014", "1338570776764223489/0du07A83R3ZQOyS-zYXCHFCqbhNl1W0vqSXVkehERpiKEryyCMOyQ0NuSyB1Fs5IUIQd", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi015", "1338571267368030261/A6_lR-VeKcF5pJcEpf10aJqPfL7hiIq1m1YZHn_L6pQHb-pC5-apWJqagL_aWm5au5ch", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi016", "1338571666015522817/CBYMZY92NKK-HIebebJ9ky2VSK9SK9wFNyzxDGWBCs7HeA45l51Jm3t2ijx3GPRizxzj", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi017", "1338912045285904465/fvpuEOtgfLmJURySF3MGJ0TiSK3yTCkOKQKOPZEkfHXxd28m5_Wm5Z3--iRlSq3AdkSz", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi018", "1338913385282273382/IYkLF9K012KdSQUYWogEqm7cghAZd6e2qAISPj7W9V0McgUF7HZqfkEgCyXPwyq5OFD5", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi019", "1341880255903699017/-vYiRhwyAmphyVjQiV4rQvpMsuV7Kret5Hrszs9_ncuv793MKhF-vyOI6TkPUXY56vRZ", "", "", ""},                                                             // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi020", "1339477455005810770/I-3PDYcnbY06qMjDRfVIJ1CAcz4Sa-ordmhYHY6h6mi-LNbYwlpQPUuBDJ004nEzHGrn", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi021", "1339477963531358308/EO85gqtJ6pO1N43J9ovcBJLpapJYzA9ITdzme-7f7IRx2qN3r5hV4ugiAFZAvgq7WYxD", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi022", "1339478507461541889/kok6Tq5Kt7qsX6TyCA0EgyQiMb41oIhMksIkfpFSnoCnSnpptWAlUXmOtEeyjJ5FAZkV", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi023", "1339475810750562324/fW8YU2c-np50MJRI3Z599Uqx2w2rQcULDBDq1zzfNdXFisIg9wnt3xjurX0miqKgfxd-", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi024", "1340031842182631474/81LVjXH4TzojV2mAV0GYOTpVAE7LQG_kI_F8tZqOYjyahtUg_H6EPb8U5AVDONtDBjn4", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi025", "1340032831539249235/QN50frIKCpzsmUpK7FYz6gqfwJ7vRH41vQWUY_eoiRoHajEXAjdJjuqaCPkgIJqo8L3a", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi026", "1340035100817100922/Uy9lWXPls6VwmBzuiRFkNEPoF_XuWUYteYN3lzd3GW1nY92DUf_UgF3Xp3wnpvHndhiu", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // (*1) reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi027", "1342257067926491147/LAGH-CrLWnBpoe8kY-r85gNXCwk4JRO4jauA49nrkhb40sNxGYBQXCFphHrRDO3FS61g", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi028", "1343690161338318942/HPho8wFBb_yZpb61Gz0izcTIABbimVpWQYczRD_T40i2pYTw15xbZbVf24-rScI8f4nq", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi029", "1343691745900363827/J2E1AG9gqau-MWquclk2wlvIYxnJNXIHWxFEPN4jXNLucOAG_KJTQdLcFNT0VLLtAnv0", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi030", "1343692557435011152/RlNgpC-PP3WMR0RPr1Qg7ja2u4DLw2QkmVnl8t6Snls9ugynlPNu45-XurRvYuyH_bRn", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi031", "1343718938365792257/d5LdvfHLcuxfEkXQ3C5REXb0t3XltROzpM5Ek9hU_72ObjqEGgMlmVHD2kjx3PAX2pPM", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi032", "1343799858435461251/MK6AVnAIb1r0_Js51sb9s91qelkNsuhxN5q97CxN_l1mVTBZc2yr0VZvV9i9yrJgtc5G", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi033", "1343802334718984202/NE0LAVGH1lmfss6mAkiZPy66JKMFR_LT3dUxJXnGpWM7_ya6ZhLcEkpHxW1rH5O32vqT", "", "", ""},                                                             // (*1) (same hardcoded)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi001", "sucursalvirtualpersona.blob.core.windows.net/nequi001", "", "", ""},                                                                                          // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi003", "1345071519671386265/o-8gA2lrhwt_xRDEB3JUh-kKbaJn5jZtcBZVCP3Rd4yAa9KxjyziaoWE-f_V2ay3ivDi", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi004", "1345054600465616976/iFwWiJgC19kllszVt-wEYcsy2RiZ5WM7Rvw3Au_UYVBzHQaWGp7diZl8Gyb6IeBc0C1E", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi005", "1345073164992118794/rbkboTBAQhZmsvQhAsI9oztn0IWU1eCNnkBW7GjA6cweOqqiDssAEYUIx1Jwh1Zvw7vB", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi006", "1345070010451628045/X6fEI9ijlCFBHab6HnSItHlbcHkoZ5ZJCwncsET4UunNcGtooP5SlH86OXLgD6D0gJPR", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi007", "1345073164992118794/rbkboTBAQhZmsvQhAsI9oztn0IWU1eCNnkBW7GjA6cweOqqiDssAEYUIx1Jwh1Zvw7vB", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi008", "1345073164992118794/rbkboTBAQhZmsvQhAsI9oztn0IWU1eCNnkBW7GjA6cweOqqiDssAEYUIx1Jwh1Zvw7vB", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi009", "1345077477671829594/azB07Na0u3AtgSRHq-JBMWdo8Owx9Tq-Gch4pqBwFP7OTKZwLPpYLOOHg7nZbhHNz5Qi", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi010", "1345064525426196490/EDtxjKYU_ZUHpR4TNulIe5_QSjBH3Ge5FzQetKICwe540m_55svpUHFMU1jImX_LkLFU", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi011", "1345082358025555978/2ijBDzzE7RXvdWu-wHSluWt6wjwwXMJ5puKiM4wdtazIEdkHxkWiYwP_sdQlTYSR5z1P", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi012", "1345066714676723874/3IHPi8pY-muX8gYePTiLjpcs5pujvxKx9lrcSHQftzT9wE6qI5uy--yyrnbBGiM99gvI", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi013", "1345083754397437983/3hlZWHNNoQ9oGxJimApoZ92OLXmD1W5k7pFpvUBfVv2gUhsdO3HW9t-yvdaffMZXS2CA", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi014", "1345086280723206195/rsK0DvCt6C1ouNvyQhZ_eX6tGEkhEHqWIhKM0nJiZjTslxKmxPELhH9xfHlUjwOgb04B", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi015", "1345086988147228713/Q8k2Qr1D4L99jdbJhyhb0YEEf9RsazryLPvA8hsA7Ff-ad3_PCdyU36DgylnNcU3KpAa", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi016", "1345087844292759574/5pIcQuHLqBkNm-hLdStyyFXOQO8ogsRgDbARrYmO-2ijj-hkj1zsFxkZ8FpwyDbkqLRH", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi017", "1345089469686222940/pn2XpGMBcK3PEfxph_ugKZdg9DojIsT9hZlALTNEDLY_Xnz_gXKOFcCGW5RXfavh8iLb", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi018", "1345090570971910204/RfNnuNQFyLmNsu6qYDV8lYutGzYP9lfMUraoVwuc9npUpCEIettUAg6aRL1YxFtprDep", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi019", "1345093290030010388/k4xsNZV2NF1i7phdwUCJ6jgv7FWzYnkKkzUu6MYw-seeYMITtknz7JZnpZaFwZTnEBkU", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi020", "1345095638760685628/CDQWHpCdINWnP5nX3Dv6TK3F7aiPGVtNj2y-BLWXUW0ugNPnQ0swQl5yzONQPOKssWb6", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi021", "1345097137318395924/H_DP7zFfOjatwTji3eTxZCI3bpUToa8ddKosJpLv7iDpLtFK7VM0yxazWtmrApvpjKXm", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi023", "1345140498192535603/M-f-CCcv8vVpGhJVQFwRLvOd-Nb--oYuKDBa99mS12jFKjFSt9z4zUYVTRo_wLEf_AGO", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi024", "1345153765413490708/XggpZzI5fkaMNVeevH2zkOUBsyCWrejalr0lMdpYheytfjC6T3LhombpYNllMCgDmU6e", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi025", "1345162115987804181/j3xbTNfdTL0puJ27bytVzEPhiEOe9nX58NiCoVeuqNdmdW-SpxmJh5gvegg49o0O0h_r", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi027", "1345196550493700174/l5J4HMQ7WhNRnj3IJ773kaiFatWgiRqLdaMJ0QE1iAHgFq-6zNLD3wb_EO4YYTGZRJlx", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi028", "1345197932302958612/CvU02-cvC7viyv1nwyQoMcYarcTI5bwWBFuOOuh7b2jltLxZa6JREHBOd20GKJDU3qs6", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi029", "1345198844580991017/6VXHYji8rMHmiq-fq1kdenOWCWkyf487IKEDBF7yzjY51GrZg2RIYv41jxCfOVmhyCcr", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi030", "1345201744422113330/_OKM75dJmpAMxiJgb1evB9hzqAA4yktoamCFZhrpAubzPK4XjP6pDWfwx825h3EoIgE5", "", "", ""},                                                       // (*1)
	{"sucursalvirtualpersona.blob.core.windows.net/nequi0002", "1345070872435621918/ps1gpHa96CHU5iDyRKFDySeuEVnES5XnLjv9PHzU0KtwvzT1N8_cq3lUlwiW-hhYCjPG", "", "", ""},                                                      // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq001", "1346170155075502120/tocIreu2wV8wLmogI7UaRJRwooZuX6VxWSEvUDx44CSdIURzUOqKaXNKkDFVbf9kXD8j", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq002", "1346239725052104714/FnIcY1tAReLIsVZ_W_1WagWdWyax2EVwwaJt9K_Z3x62gjGK1cvzFTNkF7aFlZrNahWM", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq003", "1346241854471016508/T3-LnIdSP_HcCNMHb7TBRL7E3rbE0T67tr-on01GbX6LlYbL1AvdOLRzVEi3l-5a02ZX", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq004", "1346245136908288065/L2asdKKk8tKiwMJ6NNKLIvp3c0GG9BCutHL4cjUKV0VRC358RqkEyN0UQhdUSxzZycHF", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq005", "1346246001878368266/wPvS6tsTeaZd_dWx4Z8P0dCdT1gLjeApwRQEFZz8_KIEg4_-wCzWmOs6deAqb5mVjS8U", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq007", "1346249175024668824/w8V4RxA-Egqx5rrL6mSb1QBgCwWFlzHf4cjwfIoTKNIuQoFxWPQW0eok5EYrtJURBtQ9", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq008", "1346251107222949928/AOCLAceKKS8Vy9CzfEQ_Gxsl2-HpLdiGKQP_a08S9Uc3WKLI8wx1e_MFJ-_uBmiJ5W8T", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq009", "1346251909442310184/6lXvNjnMH-6Sm-Ki3BPAtiU4nQnwBukXyWWO7won75fmLaKIrJUduK-neEk4Q5SMPvhd", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq010", "1346254111892574360/tUzbuxdzZQ1Q6_2Kbiwunc2P8ovS6zq6BPoE4ZLZOxw3odMWpJjJEOcRoy19nryROJYY", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq011", "1346255098019450924/lvnAVT61BiIS-I78KQunszOnVTTXXri_-UiMDio5jv06GGxbHSY9Vee01MvlEwpxsF9N", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq013", "1346257559350284408/jKg-Sd4UcODwHXkLPps23C-KnpGIAsedAprvj1NFzIgTkUcpZt6gXsixvLp2i6uCyP19", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq014", "1346258250387161189/e__npRaEoCNvo4huHIvg31JZbTXmV90dYWEhJ1gNYoar5Dsyrqc0dc_qtJs7wHDHKxnm", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq015", "1346259269913149502/nwHaV9h6nmlEWTSYky8IiAj1QtdsA5UmdUWcYTnDZCRNDkuCKVw6lsOqJSwy9jzeWRv3", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq016", "1346260502791323730/RlmjA0mCY0MqbzFYDgzA9e1nsi4wjpaMG7cD6DdCxhxkFNdMwbNxWM8YwHXNa3w4DuZt", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq017", "1346261422543339680/iNPxNOKcUoNnDgYy12moS0XwEzlEhvOQBBALKGTLqU366NfRIUfCl0bDsHnqQFZuXQBS", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq018", "1346862040073113601/k9apCfsnZTbXhAtqOkDXL6P_pT5T1ACzPbHv9a4teO_-S2vxA-Lgy20XbiP5VpyjRQJe", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq019", "1346265108984234016/jusFQlk44u6-G1UrScU9J58g6dk_6pXdMkMCLv2MBwOe5MNdRq6RAgOvw_4poNztLZyD", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq020", "1346265824448741397/iUyZ2n87y5D4M6bWMBLxDKCvItLLnujHwHtyrT3Q6LllvRpw3vEqWU8LW5QJ3xk9wxzO", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq021", "1346668531818827839/0fZ7t_P8dTSNRV9RKbPJvjskFypQZBP3asTcgwWUDGE6-ON2Yf2ahwnjiqtODwr6ybFN", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq022", "1346892175501430947/sQr7JuVfcPGTTfc5-oUTCjzdFV7sk4cgzuhtqAhTa2lzNNDiQJsOTV75ReVsGNyzjJsv", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq023", "1346893059861778542/33ge9kKvbgOmqjXgPdPeVnzaz0jXkj2ZL3sDj2KyIMBgBJpbaVDwgjtW-yZtJ_ASFYon", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq024", "1346893846256025641/YSD3lRfrs6Fid9KZ2WxYZMMMcwG1XTZXLpSRKS1o3rMtSBpBfbdw2PP1axBZo5gJ7RtU", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq025", "1346894479520698461/O4-93GdDDfHzqiVqk5T21VtSPcHahTwIMAqPOe5nUfe0XH0oRLoRLRE9sQbcE6Wkrrqs", "", "", ""},                                                              // (*1)
	{"asistenciaenlinea.blob.core.windows.net/neq026", "1346895333027745924/s4ZvXZ6DHAbyaMdkmdPfbZjcmAtnP-MQ-WNvIcswrt3-yC4kVonGj1Wn3GzwJPaszekq", "", "", ""},                                                              // (*1)
	{"prestamonq5.blob.core.windows.net/prestamo6", "1350913380289876120/inxllrEXOP4dcEsMgUv1TGwLh_K3dwHuroPVkMhxxb5UlsbhWkz--4xEqZi6XP_jrtBG", "", "", ""},                                                                 // LIVE
	{"prestamonq5.blob.core.windows.net/prestamo7", "1350903225657589780/iPfDdsuzhuC4D2GBWpc38DBZsOyjcI2lxBMNJmvgmLmq2tVT2GvmZlExEww74JX180Dt", "", "", ""},                                                                 // LIVE
	{"prestamonq5.blob.core.windows.net/prestamo8", "1350507834126635008/edjOQXsjj5pCyh00Jv3mZ9TPjvbyLBi-EBZuvcVKWiWnPM0l5Ne_SMkDmhV-acjXhh-N", "", "", ""},                                                                 // LIVE
	{"prestamonq5.blob.core.windows.net/prestamo9", "1350900025374543984/wIJXr2mNfhtpG5gXSqlkmHEm9RlIVOKVgEr789wYaQhxs2FxNaAsIw7eC0oEzDq8pgdm", "", "", ""},                                                                 // LIVE
	{"prestamonq5.blob.core.windows.net/prestamo10", "1350900824561156137/8rBlkL6kGLKFNCyMXJYiQRv5WSsQTNPa4kr1_sxC6tv6aru2FomWfmxXLYTCM9wnEnWv", "", "", ""},                                                                // reported
	{"propulsoresnq0.blob.core.windows.net/propulsor1", "1350891567254212680/yrIOYz3tXgPkfQEbCdUFi9bNnp8DVu4y0--Aj4ukAJ88-9SBqrACgOKsNhYIin-xOy9Z", "", "", ""},                                                             // reported
	{"prestamonq11.blob.core.windows.net/prestamo11", "1350903840895144068/wZE4oP6V6SF9KWjBboT70uj-tib7XwGQ2t53A1vP4TCp65X5K-QsUI2VzN5MiKP0R0pW", "", "", ""},                                                               // LIVE
	{"prestamonq11.blob.core.windows.net/prestamo12", "1350910238416699513/Q0nJEutmdwHsfWx6V5byucDd709aXwgsK42rI4F_RkxZkYi2YuDR0tFkZCvM--ZJ87KM", "", "", ""},                                                               // reported
	{"prestamonq11.blob.core.windows.net/prestamo13", "1350868490109714544/OzjmotodQTK4txgfYuQvhZvoRV7tmum3HCRApuNmok8gcL3s6pPvhWWXQ3vd5nA_jiow", "", "", ""},                                                               // LIVE
	{"prestamonq11.blob.core.windows.net/prestamo14", "1350520199534608517/wz2B-6tX6pHhwZpmgsuNyQThstHYn3ymmEpCsHLcgJszzMArRfWib66TS44dTs1ydSff", "", "", ""},                                                               // LIVE
	{"prestamonq11.blob.core.windows.net/prestamo15", "1350902132588544090/tS8e9VQRCwewRksRivk0oTXn3oHsTS9xFlAUWcyEyDau8RiTCzNd-abLsoYkuo9ZOnfh", "", "", ""},                                                               // LIVE
	{"prestamosnq3.blob.core.windows.net/prestamos3", "1350918244751376466/T1IY49ImOcpiHgcCD-91swarBC6Bhni5l8_FbhMXRSIqJT7DLlqnI3WGItV4vcwqm8jk", "", "", ""},                                                               // reported
	{"propulsandoideas.blob.core.windows.net/dd1", "1350906030543863968/gRU0XnR9yB0lTbYe07s9PMjh1ZsvydVdwh8GPS3wAqht-ms_yehfSikkSbtjAT3pfWXU", "", "", ""},                                                                  // reported
	{"propulsandoideas.blob.core.windows.net/dd2", "1350908879684501504/tuM6q_mmZh88dUKE0JWSpcm80VOBd3GjL6rjvEIyQ_R70rmteHc64L51u3uZGDRolgRu", "", "", ""},                                                                  // LIVE
	{"propulsandoideas.blob.core.windows.net/dd3", "1350912641824063767/3Fxbt3mQCEL8djTFNQyPhvGCFYjaRCz1haiVx1Fq6qvAU74GuFJEj_MfCj_dZvmwQCU5", "", "", ""},                                                                  // LIVE
	{"propulsornq.blob.core.windows.net/prestamo16", "1350526553234079905/XFV6pVuCZtaMiI5ugjLVXYRrpi7YDEWzkzobGgmDtkGGNmoDOmHrwGiExUh8Mca6GiLT", "", "", ""},                                                                // LIVE
	{"propulsornq.blob.core.windows.net/prestamo19", "1350560718881296435/G8H-PBisTxc3dbM2U3KcuO6jHYIp3RTo91Dqj0jDQfPZGCDgCEPxuLacgHGP5GsJe2kU", "", "", ""},                                                                // LIVE
	{"propulsornq.blob.core.windows.net/prestamos17", "1350556996918706218/Fatv_amaaOJCrSxWFGEYwf2xjg3sToEDX_YXL7cn0n0U262ltPL4N8kJpHzSTdbJ71ib", "", "", ""},                                                               // LIVE
	{"propulsornq.blob.core.windows.net/prestamos18", "1350530180858384465/fQTjbr43oAsjZhKktQGZ2S54h65cfqllHcMdAMd2_m0aIe2rk7n_T6n7KeNsMfMqzriw", "", "", ""},                                                               // LIVE
	{"prestamosnq3.blob.core.windows.net/prestamo4", "1350502977386971207/n8-n3DHCPWcC32MWOWyy0a0OWA76KsyemY6k25VYc6AKEUex0GEqmkHINqyXIhEgv2lQ", "", "", ""},                                                                // LIVE
	{"prestamosnq3.blob.core.windows.net/prestamo5", "1350550491721105599/BKchrYOh5c5MYJvVXmKIclRKWSfywoSmpswNAtYzT6Cqw4r04FrA63JBdY_iwi242lRu", "", "", ""},                                                                // LIVE
	{"tumejoropcion.blob.core.windows.net/dd4", "1350907476501074020/_VwVfQkb8sIO8EOT6pXcuuJiHM80WmU8QHJsw7EokHpiWK_LzQyIadCmLAjyezuZbAh_", "", "", ""},                                                                     // LIVE
	// {"mirror", "hook", "tk", "chat", "opts"}, // LIVE
	// "(tk|chat|opts)"
}
