package nequi4

import (
	"strconv"

	"github.com/brianvoe/gofakeit"
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq4",
		Description: "attack fake nequi4",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(base string) {

	h := common.ReqHandler{}
	h.UseJar(true)

	defaultHeaders := []common.SimpleTerm{
		{K: "Host", V: base},
		{K: "Accept", V: "*/*"},
		{K: "Accept-Language", V: "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3"},
		{K: "Accept-Encoding", V: "gzip, deflate, br, zstd"},
		{K: "Origin", V: "https://" + base},
	}

	h.SendPostEncoded("https://"+base+"/tele1.php", []common.SimpleTerm{
		{K: "usuario", V: common.GeneraCelColombia()},
		{K: "contrasena", V: common.GeneraPin(4)},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/nequi.html"},
	}...), nil)
	// Set-Cookie PHPSESSID
	// h.PrintCookies()

	h.SendPostEncoded("https://"+base+"/tele2.php", []common.SimpleTerm{
		{K: "otp", V: common.GeneraPin(6)},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp.html"},
	}...), nil)

	h.SendPostEncoded("https://"+base+"/tele3.php", []common.SimpleTerm{
		{K: "otp1", V: common.GeneraPin(6)},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp-2.html"},
	}...), nil)

	h.SendPostEncoded("https://"+base+"/tele4.php", []common.SimpleTerm{
		{K: "otp2", V: common.GeneraPin(6)},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp-3.html"},
	}...), nil)

	p := gofakeit.Person()

	h.SendPostEncoded("https://"+base+"/telegra.php", []common.SimpleTerm{
		{K: "titular", V: "Juan Camilo Rodriguez"},
		{K: "cardNumber", V: common.AddSeparator(strconv.Itoa(p.CreditCard.Number), 1, " ")},
		{K: "expiryDate", V: p.CreditCard.Exp},
		{K: "cvv", V: p.CreditCard.Cvv},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/prueba-autenti.html"},
	}...), nil)
}

var mirrors = []string{
	"prestamos.website", // DOWN
}
