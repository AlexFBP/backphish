package nequi10

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	common.Discord
}

func CreateMirrorHandler(name string) (d mirrorData) {
	v := target.GetMirrorParams(name)
	d.Webhook = v[0]
	return
}

func (t *TargetStr) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/rastrear/informacion/login.php")
	return state
}

var mirrors = [][]string{
	{"consultapuntajeneq.blob.core.windows.net/bcs", ""},
	{"laoportunidadesahora.blob.core.windows.net/prestamo", ""},
	{"prestamoinmediato.blob.core.windows.net/ch21", ""},
	{"propulsorfinanciero.blob.core.windows.net/mgne", ""},
}
