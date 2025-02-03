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

// Mirrors. The comment depending on last checked state:
//   - (*1): Apparently Down by host provider
var mirrors = [][]string{
	{"tuappnequi.blob.core.windows.net/app1", "mijo-one", "pricles", ""},          // (*1)
	{"tuappnequi.blob.core.windows.net/admin5", "server5-flame", "", "no-step"},   // (*1)
	{"nequiapp.blob.core.windows.net/admin", "sooooooes", "", ""},                 // (*1)
	{"tuappnequi.blob.core.windows.net/app3", "pepsi-zeta-eight", "", ""},         // (*1)
	{"nequiapp.blob.core.windows.net/inicio8", "bebesssa", "", ""},                // (*1)
	{"nequiapp.blob.core.windows.net/recurso", "verde-theta", "", ""},             // (*1)
	{"nequiapp.blob.core.windows.net/personales", "aventura-kappa", "", ""},       // (*1)
	{"nequiapp.blob.core.windows.net/inicio7", "como-eight", "", ""},              // (*1)
	{"tuappnequi.blob.core.windows.net/salvavidas3", "sudao", "", ""},             // (*1)
	{"tuappnequi.blob.core.windows.net/salvavidas4", "dosdias", "", ""},           // (*1)
	{"nequiapp.blob.core.windows.net/inicio9", "calleperoelegante", "", ""},       // (*1)
	{"nequico.blob.core.windows.net/prestamo", "owners-livid", "", ""},            // (*1)
	{"prestamopropulsor.blob.core.windows.net/salvavida", "inlove-kappa", "", ""}, // reported
	// {"", "", "", ""},
}
