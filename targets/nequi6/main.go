package nequi6

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq6",
		Description: "attack fake nequi6",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(GetAllCmds())
}

func GetAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

var mirrors = []string{
	`impulsayatuidea.com`,      // ALIVE
	`salvavidatepresta.online`, // ALIVE
	`tumejornq.com`,            // ALIVE
}

func attempt(base string) {
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
