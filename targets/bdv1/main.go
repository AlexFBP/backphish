package bdv1

import (
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "bdv1",
		Description: "attack fake banco de venezuela 1",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

func del() {
	common.RandDelayRange(15*time.Second, 30*time.Second)
}

func attempt(mirror string) {
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

var mirrors = []string{
	"victorr2-e98cda745618.herokuapp.com", // Apparently down
}
