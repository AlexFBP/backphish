package bdv1

import (
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

func init() {
	common.MainMenu.Add(menu.CommandOption{
		Command: "bdv1", Description: "attack fake banco de venezuela 1", Function: cmd, // ALIVE
	})
}

func cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func del() {
	common.RandDelayRange(15*time.Second, 30*time.Second)
}

func attempt() {

	const mirror = "victorr2-e98cda745618.herokuapp.com"

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
