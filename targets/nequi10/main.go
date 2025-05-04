package nequi10

import (
	"time"

	"github.com/AlexFBP/backphish/common"
)

type TargetStr struct {
	common.TargetWithParams
}

var target *TargetStr

func init() {
	target = &TargetStr{}
	target.Prefix = "nq10"
	target.Description = "attack fake nequi10"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *TargetStr) Handler(mirror string) {
	m := CreateMirrorHandler(mirror)
	u := CreateUserData()

	// from index.html
	m.SendDiscord(u.DataForStep(1))

	// from prestamo.php.html
	m.SendDiscord(u.DataForStep(2))

	// from error_dinamica.php.html
	m.SendDiscord(u.DataForStep(3))

	// from error-otp.html
	m.SendDiscord(u.DataForStep(4))
}

func (t *TargetStr) EstimateDuration() time.Duration {
	return 0
}
