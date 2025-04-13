package nequi9

import (
	"fmt"
	"time"

	"github.com/AlexFBP/backphish/common"
)

type mirrorManager struct {
	common.TargetWithParams
}

var target *mirrorManager

func init() {
	target = &mirrorManager{}
	target.Prefix = "nq9"
	target.Description = "attack fake nequi9"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *mirrorManager) Handler(mirror string) {
	d := newUser()
	m := mirrData(mirror)
	m.Discord.UseJar(true)

	var op map[string]string
	alt_hook := func() {
		m.SetAltHook(m.Token)
	}
	reset := func() {
		op = map[string]string{}
	}
	reset()

	if m.Token != "" && m.Chat != "" {
		m.SendToTelegram(d.DataForStep(0, nil).Content)
		common.RandDelayRange(10*time.Second, 30*time.Second)
	}

	// (step 1) POST application/json webhook (from ini.html)
	if m.HasOption("alt1") {
		op["wrap"] = "😎😎"
	}
	m.DetectHookFromURL(fmt.Sprintf("https://%s/ini.html", mirror))
	m.SendDiscord(d.DataForStep(1, op))
	reset()
	common.RandDelayRange(10*time.Second, 30*time.Second)

	// (step 2) POST webhook (from neq.html)
	if m.HasOption("alt1") {
		op["wrap-ini"] = "😎😎"
		op["wrap-end"] = "😎😎"
		if m.HasOption("alt3") {
			op["wrap-end"] += " goku"
		}
	}
	if m.HasOption("alt2") {
		op["alt2"] = ""
	}
	m.DetectHookFromURL(fmt.Sprintf("https://%s/neq.html", mirror))
	m.SendDiscord(d.DataForStep(2, op))
	reset()
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// (step 3) POST webhook (from otp.html)
	if m.HasOption("alt1") {
		if m.HasOption("alt2") {
			op["wrap-ini"] = "😎😎"
			op["wrap-end"] = "😎😎"
		} else {
			op["wrap-ini"] = "🥳🥳"
			op["wrap-end"] = "🥳"
			op["cont1"] = ""
		}
	}
	if m.HasOption("alt2") {
		op["usr1"] = ""
	}
	if m.HasOption("wh-bypass-3") {
		alt_hook()
	}
	m.DetectHookFromURL(fmt.Sprintf("https://%s/otp.html", mirror))
	m.SendDiscord(d.DataForStep(3, op))
	reset()
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// (step 4) POST back2 (from loading --> dinamica)
	if m.HasOption("alt1") {
		op["wrap"] = "🥳"
		op["cont1"] = ""
	}
	if m.HasOption("wh-bypass-4") {
		alt_hook()
	}
	m.DetectHookFromURL(fmt.Sprintf("https://%s/dinamica.html", mirror))
	m.SendDiscord(d.DataForStep(4, op))
	reset()
	common.RandDelayRange(35*time.Second, 50*time.Second)

	// (step 5) POST back2 (from load --> dinamica2)
	if m.HasOption("alt1") {
		op["wrap"] = "🥳"
		op["cont1"] = ""
	}
	if m.HasOption("wh-bypass-5") {
		alt_hook()
	}
	m.DetectHookFromURL(fmt.Sprintf("https://%s/dinamica2.html", mirror))
	m.SendDiscord(d.DataForStep(5, op))
	reset()
	common.RandDelayRange(35*time.Second, 50*time.Second)
}
