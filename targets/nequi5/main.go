package nequi5

import (
	"time"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "nq5"
	target.Description = "attack fake nequi5"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

// Mirrors. The comment depending on last checked state:
//   - (*1): "no such host" - Down
//   - (*2): Timeout?
//   - (*3): "no such host" - Down, but domain still alive...
var mirrors = []string{
	"neqprepropul.online",         // (*1)
	"obtentupropulsivo.online",    // ALIVE (*3)
	"propulsorempresasneq.online", // Down? (*3)
}

func (t *mirrorTarget) Handler(base string) {
	h := common.ReqHandler{}
	h.UseJar(true)

	commonHeaders := []common.SimpleTerm{
		// {K: "Host", V: base},
		{K: "Accept", V: "*/*"},
		{K: "Accept-Language", V: "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3"},
		{K: "Accept-Encoding", V: "gzip, deflate, br, zstd"},
		{K: "X-Requested-With", V: "XMLHttpRequest"},
		{K: "Origin", V: "https://" + base},
	}

	devices := []string{"Android", "iPhone"}
	device := devices[1]
	cel := common.AddSeparator(common.GeneraCelColombia(), 0, " ")
	pin := common.GeneraPin(4)

	req1 := func(end string) {
		h.SendPostEncoded("https://"+base+"/process/pasousuario.php", []common.SimpleTerm{
			{K: "pass", V: cel + " - " + pin},
			{K: "dis", V: device},
		}, append(commonHeaders, []common.SimpleTerm{
			{K: "Referer", V: "https://" + base + "/bdigital/" + end},
		}...), nil)
		// Set-Cookie usuario,contrasena,registro,estado
	}
	req1("login.html")

	checkStat := func() {
		// Check "status" by polling...
		// POST https:// base /process/estado.php
		// (Send all cookies)
		// ... while last status changes
	}
	checkStat()

	req2 := func() {
		// Cookie contrasena,registro,estado
		// Cookie bmuid, cdContextId, cdSNum ???
		h.SendPostEncoded("https://"+base+"/process/pasoOTP.php", []common.SimpleTerm{
			{K: "otp", V: common.GeneraPin(6)},
		}, append(commonHeaders, []common.SimpleTerm{
			{K: "Referer", V: "https://" + base + "/bdigital/dinamica.php"},
		}...), nil)
		// Set-Cookie cdinamica
	}
	req2()
	checkStat()

	req1("inicio.php")
	checkStat()

	req2()
	checkStat()
}

func (t *mirrorTarget) EstimateDuration() time.Duration {
	return 60 * time.Second
}

func (t *mirrorTarget) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/bdigital/login.html")
	return state
}
