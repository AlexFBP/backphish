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

// Mirrors. The comment depending on last checked state:
//   - (*1): "no such host" - Down
//   - (*2): TLS error? internal error? Down?
var mirrors = []string{
	`co.davivnda.com`,          // ALIVE (*1)
	`col.davivendacol.website`, // ALIVE (*1)
	`col.davivnda.com`,         // ALIVE? (*1)
	`davivnda.com`,             // REPORTED
	`davivendacol.website`,     // ALIVE
	`inf.davivnda.com`,         // REPORTED (*1)
	`tucrdtodavivienda.com`,    // (*2)
	`val.davivnda.com`,         // ALIVE (*1)
	`web.davivendacol.website`, // ALIVE
}

func attempt(base string) {

	h := common.ReqHandler{}
	// h.UseJar(true)

	h.SendMultipart("https://"+base+"/ban.php", map[string]io.Reader{
		"t1": strings.NewReader("c&eacutedula"),
		"t2": strings.NewReader(common.GeneraNIPcolombia()),
		"t3": strings.NewReader(common.GeneraPin(6)),
	}, nil, nil)

}
