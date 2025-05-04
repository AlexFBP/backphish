package nequi10

import (
	"time"

	"github.com/brianvoe/gofakeit"

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

	// from lalo.html
	common.RandDelayRange(20*time.Second, 40*time.Second)
	m.SendDiscord(u.DataForStep(0))

	attempts := gofakeit.Number(1, 3)
	for ; attempts > 0; attempts-- {
		// from index.html
		common.RandDelayRange(20*time.Second, 40*time.Second)
		m.SendDiscord(u.DataForStep(1))

		// from prestamo.php.html
		common.RandDelayRange(15*time.Second, 30*time.Second)
		m.SendDiscord(u.DataForStep(2))

		// from error_dinamica.php.html
		common.RandDelayRange(15*time.Second, 30*time.Second)
		m.SendDiscord(u.DataForStep(3))

		// from error-otp.html
		common.RandDelayRange(15*time.Second, 30*time.Second)
		m.SendDiscord(u.DataForStep(4))
	}
}

func (t *TargetStr) EstimateDuration() time.Duration {
	return 107 * time.Second
}
