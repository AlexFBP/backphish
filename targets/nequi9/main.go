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
	d := newUser()
	m := mirrData(mirror)

	if m.Tk != "" && m.Chat != "" {
		m.sendToTelegram(&h, d.DataForStep(0).Content)
		common.RandDelayRange(10*time.Second, 30*time.Second)
	}
	// (step 1) POST application/json webhook (from ini.html)
	m.sendDiscord(&h, d.DataForStep(1))
	common.RandDelayRange(10*time.Second, 30*time.Second)

	// (step 2) POST webhook (from neq.html)
	m.sendDiscord(&h, d.DataForStep(2))
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// (step 3) POST webhook (from otp.html)
	m.sendDiscord(&h, d.DataForStep(3))
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// (step 4) POST back2 (from loading --> dinamica)
	m.sendDiscord(&h, d.DataForStep(4))
	common.RandDelayRange(35*time.Second, 50*time.Second)

	// (step 5) POST back2 (from load --> dinamica2)
	m.sendDiscord(&h, d.DataForStep(5))
	common.RandDelayRange(35*time.Second, 50*time.Second)
}
