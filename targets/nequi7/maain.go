package nequi7

import (
	"time"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "nq7",
		Description: "attack fake nequi7",
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
	d := mirrData(mirror)
	u := newUser()

	// ini.html form message
	d.SendToTelegram(&h, u.template1())
	common.RandDelayRange(10*time.Second, 30*time.Second)

	// neq.html form message
	d.SendToTelegram(&h, u.template2(0))
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// otp.html form message
	d.SendToTelegram(&h, u.template2(1))
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// loading.html has no form data

	// dinamica.html form message
	d.SendToTelegram(&h, u.template2(2))
	common.RandDelayRange(35*time.Second, 50*time.Second)

	// load.html has no form data

	// dinamica2.html form message
	d.SendToTelegram(&h, u.template2(3))
	common.RandDelayRange(35*time.Second, 50*time.Second)
}
