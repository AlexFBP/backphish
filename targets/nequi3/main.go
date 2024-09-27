package nequi3

import (
	"fmt"
	"strings"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	base := "tucredirapido.com"
	h := common.ReqHandler{}
	h.UseJar(true)

	cel := common.GeneraCelColombia()
	h.SendPostEncoded(
		"https://"+base+"/process/pasousuario.php",
		map[string]string{
			"pass": fmt.Sprintf("%s - %sOTP %s", // "316 546 4981 - 3795OTP 524844"
				strings.Join([]string{cel[:3], cel[3:6], cel[6:]}, " "),
				common.GeneraPin(4), common.GeneraPin(6)),
			"dis":   "Android",
			"banco": "bancolombia",
		},
		map[string]string{
			"Host":             base,
			"Accept":           "*/*",
			"Accept-Language":  "es-MX,es;q=0.8,en-US;q=0.5,en;q=0.3",
			"Accept-Encoding":  "gzip, deflate, br, zstd",
			"X-Requested-With": "XMLHttpRequest",
			"Origin":           "https://" + base,
			"Referer":          "https://" + base + "/bdigital/otp.php?p",
		},
		nil,
	)
}
