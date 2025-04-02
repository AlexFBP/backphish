package nequi9

import (
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq9",
		Description: "attack fake nequi9",
		Mirrors:     getMirrorNames(),
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func attempt(mirror string) {
	h := common.ReqHandler{}
	h.UseJar(true)
	d := newUser()
	m := mirrData(mirror)

	back_hook := ""
	var op map[string]string
	cust_hook := func(hook string) {
		if hook != "" {
			back_hook = m.Webhook
			m.Webhook = hook
		}
	}
	alt_hook := func() {
		cust_hook(m.Token)
	}
	reset := func() {
		op = map[string]string{}
		if back_hook != "" {
			m.Webhook = back_hook
			back_hook = ""
		}
	}
	reset()

	if m.Token != "" && m.Chat != "" {
		m.SendToTelegram(&h, d.DataForStep(0, nil).Content)
		common.RandDelayRange(10*time.Second, 30*time.Second)
	}

	// (step 1) POST application/json webhook (from ini.html)
	if m.HasOption("alt1") {
		op["wrap"] = "😎😎"
	}
	m.SendDiscord(&h, d.DataForStep(1, op))
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
	m.SendDiscord(&h, d.DataForStep(2, op))
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
	m.SendDiscord(&h, d.DataForStep(3, op))
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
	m.SendDiscord(&h, d.DataForStep(4, op))
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
	m.SendDiscord(&h, d.DataForStep(5, op))
	reset()
	common.RandDelayRange(35*time.Second, 50*time.Second)
}
