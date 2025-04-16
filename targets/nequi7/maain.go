package nequi7

import (
	"time"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetWithParams
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "nq7"
	target.Description = "attack fake nequi7"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *mirrorTarget) Handler(mirror string) {
	d := mirrData(mirror)
	u := newUser()

	// ini.html form message
	d.SendToTelegram(d.selectTemplate(&u, 1))
	common.RandDelayRange(10*time.Second, 30*time.Second)

	// neq.html form message
	d.SendToTelegram(d.selectTemplate(&u, 2))
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// otp.html form message
	d.SendToTelegram(d.selectTemplate(&u, 3))
	common.RandDelayRange(15*time.Second, 30*time.Second)

	// loading.html has no form data

	// dinamica.html form message
	d.SendToTelegram(d.selectTemplate(&u, 4))
	common.RandDelayRange(35*time.Second, 50*time.Second)

	// load.html has no form data

	// dinamica2.html form message
	d.SendToTelegram(d.selectTemplate(&u, 5))
	common.RandDelayRange(35*time.Second, 50*time.Second)
}

func (t *mirrorTarget) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/ini.html")
	return state
}
