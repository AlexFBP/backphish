package dv1

import (
	"io"
	"strings"

	"github.com/turret-io/go-menu/menu"

	"github.com/AlexFBP/backphish/common"
)

var target *common.Target

func init() {
	target = &common.Target{
		Prefix:      "dv1",
		Description: "attack fake davivienda 1",
		Mirrors:     mirrors,
		Handler:     attempt,
	}
}

func GetAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

var mirrors = []string{
	`inf.davivnda.com`,
	`davivnda.com`,
	`tucrdtodavivienda.com`,
}

func attempt(base string) {

	h := common.ReqHandler{}
	// h.UseJar(true)

	h.SendMultipart("https://"+base+"/ban.php", map[string]io.Reader{
		"t1": strings.NewReader("c&eacutedula"),
		"t2": strings.NewReader("78264583"),
		"t3": strings.NewReader("832411"),
	}, nil, nil)

}
