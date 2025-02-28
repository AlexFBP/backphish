package nequi9

import (
	"strings"

	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	Webhook, Tk, Chat string
	Opts              []string
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Webhook, d.Tk, d.Chat = v[1], v[2], v[3]
			d.Opts = strings.Split(v[4], ",")
			break
		}
	}
	return
}

func (m *mirrorData) hasOption(opt string) bool {
	for _, v := range m.Opts {
		if strings.TrimSpace(v) == strings.TrimSpace(opt) {
			return true
		}
	}
	return false
}

func (m *mirrorData) sendDiscord(h *common.ReqHandler, data hookData) {
	h.SendJSON("https://discord.com/api/webhooks/"+m.Webhook, data, nil, nil)
}

func (m *mirrorData) sendToTelegram(h *common.ReqHandler, msg string) {
	h.SendJSON("https://api.telegram.org/bot"+m.Tk+"/sendMessage", common.TgMsg{
		ChatID:    m.Chat,
		Text:      msg,
		ParseMode: "Markdown",
	}, nil, nil)
}

func getMirrorNames() (names []string) {
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}

var mirrors = [][]string{
	{"sucursalvirtuall.blob.core.windows.net/nequi001", "1337922396681932920/gUxfz9zq-H068Qh0IEr6G2qncQ7XEN0E-Nev5Z1zJC4_PuwBHa_9rCll0RC0hyZkaRzX", "", "", ""},                                                             // reported -  "Z1xGSEFXVlxcQlQuJDUvCVAqOT0OFDcoVyUTVSAHEzYjJDYGCA4tNQU5GlECNA==", "fV5CQUdWVVtQQF5YUEo=", "Porqueblesnoesejemplodenadaaaaa"
	{"sucursalvirtuall.blob.core.windows.net/nequi002", "1338548943050637343/b_xUBPOsA4RwRmZfZZTSzJUXFZ_QDmEFOvYdhFtb9RrTGZ3ptAt9Fn69HIPFODQL0Llq", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi003", "1338558488074391605/-2wgbBayVSJDEWkkEjTaBQO7vR_wVRVCRINmB8My-216llzcPbL_4z6qa6SoNXuFqCPu", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi004", "1338559433109667893/-5gdmIXjr-eMPnL6_ONtIDKbXbPUQl2CrWok1s7EOb0ygEMi7cJxh71YIOqASWamCY3I", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi005", "1338560850050547744/uzAHiOnJ_t3pEvlGqorGxIwcPX49xTgTFG-GH-YYzEIeDg17LtTny8UPgbgGssn16yjA", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi006", "1338562843296923721/BCcezLTQBFgy6lgEEO077YAe4LGLXWTqdGWv3QLupG-gW0IXzHuHdO1hF34Ci1Ebdmm-", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi007", "1338563500209078303/RSOYlk4En5LCJQOagUn12jqsgyxAnqO-BTfpxDf2z2nPPlH0fCOztbu0m1Fhw5tHLbm9", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi008", "1338564040821309461/A2IHsI_A4DE9qsVSTw3Jrl5w-2nTszdbzJrqZAibOwBIem1KQYDGPzazDlxsutw1InXi", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi009", "1338566707609862235/KvS8FT3dGbEUvKH-bQGV2_xzVoGxKArj4aq39PsEA53ww2p8qYJ5srXgauApMa5WmR0Z", "", "", ""},                                                             // reported - (same hardcoded) - HOLD: TK1 TK2 TK2 TK2
	{"sucursalvirtuall.blob.core.windows.net/nequi010", "1338567600279523450/u--zHUjQZ5p9ZcnuUN5yKzSjAN4iAsaKEKB1E97rqUYhmbrQ058QwHW1c-WLy4Io-S58", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi011", "1338568450125336669/gMtBmDXuqzn-QA_9NkvM5YXZK9U1yB3kmZ4zpYQqzvac_3PqVKXX-n9dK_Lp_0GevYSp", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi012", "1338569266336895049/gIyNCUM92yX7B5_7od9WelHUFVrhMgLBfXOTtH6W9EPTkEIu1ortl5iPYqE6sMKoT3_D", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi013", "1338570085090201620/uCeggXpE5SJvYq6r_bsGeYVVTtycxZbAZf_E7IT4gX-fWgYode3LtfH545543m4iIdq6", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi014", "1338570776764223489/0du07A83R3ZQOyS-zYXCHFCqbhNl1W0vqSXVkehERpiKEryyCMOyQ0NuSyB1Fs5IUIQd", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi015", "1338571267368030261/A6_lR-VeKcF5pJcEpf10aJqPfL7hiIq1m1YZHn_L6pQHb-pC5-apWJqagL_aWm5au5ch", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi016", "1338571666015522817/CBYMZY92NKK-HIebebJ9ky2VSK9SK9wFNyzxDGWBCs7HeA45l51Jm3t2ijx3GPRizxzj", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi017", "1338912045285904465/fvpuEOtgfLmJURySF3MGJ0TiSK3yTCkOKQKOPZEkfHXxd28m5_Wm5Z3--iRlSq3AdkSz", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi018", "1338913385282273382/IYkLF9K012KdSQUYWogEqm7cghAZd6e2qAISPj7W9V0McgUF7HZqfkEgCyXPwyq5OFD5", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi019", "1341880255903699017/-vYiRhwyAmphyVjQiV4rQvpMsuV7Kret5Hrszs9_ncuv793MKhF-vyOI6TkPUXY56vRZ", "", "", ""},                                                             // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi020", "1339477455005810770/I-3PDYcnbY06qMjDRfVIJ1CAcz4Sa-ordmhYHY6h6mi-LNbYwlpQPUuBDJ004nEzHGrn", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi021", "1339477963531358308/EO85gqtJ6pO1N43J9ovcBJLpapJYzA9ITdzme-7f7IRx2qN3r5hV4ugiAFZAvgq7WYxD", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi022", "1339478507461541889/kok6Tq5Kt7qsX6TyCA0EgyQiMb41oIhMksIkfpFSnoCnSnpptWAlUXmOtEeyjJ5FAZkV", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi023", "1339475810750562324/fW8YU2c-np50MJRI3Z599Uqx2w2rQcULDBDq1zzfNdXFisIg9wnt3xjurX0miqKgfxd-", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi024", "1340031842182631474/81LVjXH4TzojV2mAV0GYOTpVAE7LQG_kI_F8tZqOYjyahtUg_H6EPb8U5AVDONtDBjn4", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi025", "1340032831539249235/QN50frIKCpzsmUpK7FYz6gqfwJ7vRH41vQWUY_eoiRoHajEXAjdJjuqaCPkgIJqo8L3a", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi026", "1340035100817100922/Uy9lWXPls6VwmBzuiRFkNEPoF_XuWUYteYN3lzd3GW1nY92DUf_UgF3Xp3wnpvHndhiu", "7397391450:AAEZLjSyIv6kvnpnBFWsoj6gEoIGl8KmYZ8", "-1002463610738", ""}, // reported - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi027", "1342257067926491147/LAGH-CrLWnBpoe8kY-r85gNXCwk4JRO4jauA49nrkhb40sNxGYBQXCFphHrRDO3FS61g", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi028", "1343690161338318942/HPho8wFBb_yZpb61Gz0izcTIABbimVpWQYczRD_T40i2pYTw15xbZbVf24-rScI8f4nq", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi029", "1343691745900363827/J2E1AG9gqau-MWquclk2wlvIYxnJNXIHWxFEPN4jXNLucOAG_KJTQdLcFNT0VLLtAnv0", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi030", "1343692557435011152/RlNgpC-PP3WMR0RPr1Qg7ja2u4DLw2QkmVnl8t6Snls9ugynlPNu45-XurRvYuyH_bRn", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi031", "1343718938365792257/d5LdvfHLcuxfEkXQ3C5REXb0t3XltROzpM5Ek9hU_72ObjqEGgMlmVHD2kjx3PAX2pPM", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi032", "1343799858435461251/MK6AVnAIb1r0_Js51sb9s91qelkNsuhxN5q97CxN_l1mVTBZc2yr0VZvV9i9yrJgtc5G", "", "", ""},                                                             // LIVE - (same hardcoded)
	{"sucursalvirtuall.blob.core.windows.net/nequi033", "1343802334718984202/NE0LAVGH1lmfss6mAkiZPy66JKMFR_LT3dUxJXnGpWM7_ya6ZhLcEkpHxW1rH5O32vqT", "", "", ""},                                                             // LIVE - (same hardcoded)
	// {"mirror", "hook", "tk", "chat", "opts"}, // LIVE
}
