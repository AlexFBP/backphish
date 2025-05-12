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
	{"consultapuntajeneq.blob.core.windows.net/bcs", "1366580698063835177/YGMx4A3pPM4cPWiQO9aVAITUTIELs0HFKs2lx-xLpADPx9P9N7eUUS-a6kKNXqpveI3f"},
	{"laoportunidadesahora.blob.core.windows.net/prestamo", "1366915502286573598/b0pWJjJnRuXbWetHCqunAJc4gYfpoFknnfMqFTLSPdZV0WE6glPHFEro--Ulwtx0wyi9"},
	{"prestamoinmediato.blob.core.windows.net/ch21", "1364990793529557033/XUxENAr4cGEz8c745sD34v7OhgFTNfYvTsOpU8P5Txfktq3QXR6DpCWD8tWkAt0lezYe"},
	{"propulsorfinanciero.blob.core.windows.net/mgne", "1366469597301047297/YCx3YZR-0zyUfZJkuvnzYyrLfH9VOwgVlBHJBLW0FcLQX8-qu7BVd2ew-MOcej3RkTKu"},
	{"prestamosalvavidas1.blob.core.windows.net/dd1", "1369710052159520788/IcqXT0buIaTIkGjP8ovXc2UPG-pFbYMZFMF1BTsn0XBoJLnOV4S1slHk-JyV2eAdhHct"},
	{"prestamosalvavidas1.blob.core.windows.net/dd2", "1369711269086494741/Y5pnVsORsm6Gvva1BfZffKVPOsXUCK7zYlsxYzMt4lRBm6dMxTR6LFevJWnJm5ih-v0V"},
	// {"MIRROR","WEBHOOK"},
}
