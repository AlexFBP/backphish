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

func getMirrorNames() (names []string) {
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}

var mirrors = [][]string{
	{"sucursalvirtuall.blob.core.windows.net/nequi001", "1337922396681932920/gUxfz9zq-H068Qh0IEr6G2qncQ7XEN0E-Nev5Z1zJC4_PuwBHa_9rCll0RC0hyZkaRzX", "", "", ""}, // "Z1xGSEFXVlxcQlQuJDUvCVAqOT0OFDcoVyUTVSAHEzYjJDYGCA4tNQU5GlECNA==", "fV5CQUdWVVtQQF5YUEo=", "Porqueblesnoesejemplodenadaaaaa"
	// {"", "", "", ""},
	// {"", "", "", ""},
	// {"", "", "", ""},
	// {"", "", "", ""},
}
