package nequi6

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "nq6"
	target.Description = "attack fake nequi6"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

// Mirrors. The comment depending on last checked state:
//   - (*1): "no such host" - Down
//   - (*3): "no such host" - Down, but domain still alive...
var mirrors = []string{
	`impulsayatuidea.com`,      // (*1)
	`salvavidatepresta.online`, // (*3)
	`tumejornq.com`,            // (*1)
}

func (t *mirrorTarget) Handler(base string) {
	h := common.ReqHandler{}
	h.UseJar(true)
	commonHeaders := []common.SimpleTerm{
		{K: "Host", V: base},
		{K: "Accept", V: "*/*"},
		{K: "Accept-Language", V: "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3"},
		{K: "Accept-Encoding", V: "gzip, deflate, br, zstd"},
		{K: "X-Requested-With", V: "XMLHttpRequest"},
		{K: "Origin", V: "https://" + base},
	}

	devices := []string{"Android", "iPhone"}
	device := common.PickRand(devices)
	cel := common.AddSeparator(common.GeneraCelColombia(), 0, " ")
	pin := common.GeneraPin(4)
	otp := common.GeneraPin(6)

	h.SendPostEncoded("https://"+base+"/process2/pasousuario.php", []common.SimpleTerm{
		{K: "pass", V: pin},
		{K: "user", V: cel + " Dna: " + otp},
		{K: "dis", V: device},
		{K: "banco", V: "Nequi"},
	}, common.JoinSlices(commonHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/inicio.php?p"},
	}), nil)

	h.SendPostEncoded("https://"+base+"/process2/pasoOTP.php", []common.SimpleTerm{
		{K: "otp", V: common.GeneraPin(6)},
	}, common.JoinSlices(commonHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp.php?p"},
	}), nil)
}

func (t *mirrorTarget) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/inicio.php?p")
	return state
}
