package nequi5

import (
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

func init() {
	common.MainMenu.Add(menu.CommandOption{
		Command: "nq5", Description: "attack fake nequi5", Function: Cmd,
	})
}

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	h := common.ReqHandler{}
	h.UseJar(true)

	const base = "obtentupropulsivo.online"
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
