package ms01

import (
	"fmt"
	"strings"
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
	target.Prefix = "ms1"
	target.Description = "attack fake MS login"
	target.SetMirrors(&mirrors)
	common.MainMenu.Register(target)
}

func (t *mirrorTarget) Handler(base string) {
	h := common.ReqHandler{}
	// common.RandDelay(20, 60)
	domains := []string{
		"hotmail.com",
		"live.com",
		"outlook.com",
	}
	em := gofakeit.Email()
	em = fmt.Sprintf("%s@%s", strings.Split(em, "@")[0], common.PickRand(domains))
	form := []common.SimpleTerm{
		{K: "argaml", V: em},
		{K: "argapw", V: common.RandPass1(5, 20)},
	}
	pin := common.GeneraPin(4)
	form = append(form, []common.SimpleTerm{
		{K: "pn1", V: pin},
		{K: "pn2", V: pin},
	}...)
	if gofakeit.Bool() {
		form = append(form, common.SimpleTerm{K: "KMSI", V: "on"})
	}
	h.SendPostEncoded("http://"+base+"/next.php",
		form,
		[]common.SimpleTerm{
			{K: "Host", V: "" + base + ""},
			{K: "Origin", V: "http://" + base + ""},
			{K: "Referer", V: "http://" + base + "/Iniciar%20Sesi%C3%B3n.html"},
		},
		nil,
	)
}

func (t *mirrorTarget) EstimateDuration() time.Duration {
	return 90 * time.Second
}

var mirrors = []string{
	"recuperacion004.0hi.me", // DOWN
}

func (t *mirrorTarget) MirrorStatus(mirror string) int {
	state, _ := common.CheckHostState("http://" + mirror + "/Iniciar%20Sesi%C3%B3n.html")
	return state
}
