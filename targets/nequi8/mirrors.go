package nequi8

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	Back1, Back2, Opts string
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Back1, d.Opts = v[1], v[3]
			if v[2] == "" {
				d.Back2 = v[1]
			} else {
				d.Back2 = v[2]
			}
			break
		}
	}
	return
}

func (m *mirrorData) send(h *common.ReqHandler, data userData, one bool) {
	back := ""
	if one {
		back = m.Back1
	} else {
		back = m.Back2
	}
	h.SendJSON("https://"+back+".vercel.app/api/send-message", data, nil, nil)
}

func getMirrorNames() (names []string) {
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}

var mirrors = [][]string{
	{"tuappnequi.blob.core.windows.net/app1", "mijo-one", "pricles", ""},
	{"tuappnequi.blob.core.windows.net/admin5", "server5-flame", "", "no-step"},
	{"nequiapp.blob.core.windows.net/admin", "sooooooes", "", ""},
	{"tuappnequi.blob.core.windows.net/app3", "pepsi-zeta-eight", "", ""},
	{"nequiapp.blob.core.windows.net/inicio8", "bebesssa", "", ""},
	{"nequiapp.blob.core.windows.net/recurso", "verde-theta", "", ""},
	{"nequiapp.blob.core.windows.net/personales", "aventura-kappa", "", ""},
	{"nequiapp.blob.core.windows.net/inicio7", "como-eight", "", ""},
	{"tuappnequi.blob.core.windows.net/salvavidas3", "sudao", "", ""},   // reported
	{"tuappnequi.blob.core.windows.net/salvavidas4", "dosdias", "", ""}, // reported
	// {"", "", "", ""},
}
