package bdv1

import (
	"time"

	"github.com/brianvoe/gofakeit"

	"github.com/AlexFBP/backphish/common"
)

type mirrorTarget struct {
	common.TargetSimple
}

var target *mirrorTarget

func init() {
	target = &mirrorTarget{}
	target.Prefix = "bdv1"
	target.Description = "attack fake banco de venezuela 1"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func del() {
	common.RandDelayRange(15*time.Second, 30*time.Second)
}

func (t *mirrorTarget) Handler(mirror string) {
	h := common.ReqHandler{}
	h.UseJar(true)

	user := gofakeit.Username()
	pass := common.RandPass1(8, 15)

	h.SendJSON("https://"+mirror+"/usuario/", map[string]string{
		"contrasena": pass,
		"usuario":    user,
	}, []common.SimpleTerm{}, nil)

	del()
	attempts := gofakeit.Number(2, 5)

	for ; attempts > 0; attempts-- {

		h.SendJSON("https://"+mirror+"/codigo/", map[string]string{
			"sms":     common.GeneraPin(8),
			"usuario": user,
		}, []common.SimpleTerm{}, nil)

		del()
	}

}

func (t *mirrorTarget) EstimateDuration() time.Duration {
	return 90 * time.Second
}

var mirrors = []string{
	"victorr2-e98cda745618.herokuapp.com", // Apparently down
}

func (t *mirrorTarget) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("https://" + mirror + "/")
	return state
}
