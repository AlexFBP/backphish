package nequi8

import (
	"strings"

	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	Back1, Back2 string
	Opts         []string
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Back1 = v[1]
			d.Opts = strings.Split(v[3], ",")
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

func (m *mirrorData) hasOption(opt string) bool {
	for _, v := range m.Opts {
		if strings.TrimSpace(v) == strings.TrimSpace(opt) {
			return true
		}
	}
	return false
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
	{"sucursalnequi.blob.core.windows.net/base003", "git3-sand", "", "step1"},     // reported
	{"sucursalnequi.blob.core.windows.net/base002", "git7", "", "step1"},          // reported
	{"sucursalnequi.blob.core.windows.net/base013", "git2-sigma", "", "step1"},    // reported
	{"sucursalnequi.blob.core.windows.net/base026", "git20", "", "step1"},         // reported
	{"prestamopropulsor.blob.core.windows.net/salvavidas", "git10", "", ""},       // reported
	{"prestamopropulsor.blob.core.windows.net/propulsor1", "git15", "", "step1"},  // reported
	// {"", "", "", ""},
}
