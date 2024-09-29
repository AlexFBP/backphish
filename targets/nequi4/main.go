package nequi4

import (
	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	const base = "prestamos.website"

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
		{K: "usuario", V: "3224684616"},
		{K: "contrasena", V: "8438"},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/nequi.html"},
	}...), nil)
	// Set-Cookie PHPSESSID
	// h.PrintCookies()

	h.SendPostEncoded("https://"+base+"/tele2.php", []common.SimpleTerm{
		{K: "otp", V: "435049"},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp.html"},
	}...), nil)

	h.SendPostEncoded("https://"+base+"/tele3.php", []common.SimpleTerm{
		{K: "otp1", V: "782364"},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp-2.html"},
	}...), nil)

	h.SendPostEncoded("https://"+base+"/tele4.php", []common.SimpleTerm{
		{K: "otp2", V: "298373"},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/otp-3.html"},
	}...), nil)

	h.SendPostEncoded("https://"+base+"/telegra.php", []common.SimpleTerm{
		{K: "titular", V: "Juan Camilo Rodriguez"},
		{K: "cardNumber", V: "4513 7649 7005 4394"},
		{K: "expiryDate", V: "08/32"},
		{K: "cvv", V: "969"},
	}, append(defaultHeaders, []common.SimpleTerm{
		{K: "Referer", V: "https://" + base + "/prueba-autenti.html"},
	}...), nil)
}
