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
	common.MainMenu.AddMany(getAllCmds())
}

func getAllCmds() []menu.CommandOption {
	return target.GetAllCmds()
}

// Mirrors. The comment depending on last checked state:
//   - (*1): "no such host" - Down
//   - (*2): TLS error? internal error? Down?
//   - (*3): "no such host" - Down, but domain still alive...
var mirrors = []string{
	`cl.davivendacol.website`,           // (*1)
	`co.davivnda.com`,                   // (*1)
	`col.davivendacol.website`,          // (*1)
	`col.davivnda.com`,                  // (*1)
	`davivenda.website`,                 // (*1)
	`davivendacol.online`,               // (*1)
	`davivendacol.website`,              // (*1)
	`davivnda.com`,                      // REPORTED (*1)
	`dvviendacol.site`,                  // (*1)
	`inf.davivnda.com`,                  // REPORTED (*1)
	`inf.dvviendacol.site`,              // (*1)
	`informacion.davivendacol.website`,  // (*1)
	`pre-aprobado.davivendacol.website`, // (*1)
	`tucrdtodavivienda.com`,             // (*3)
	`val.davivnda.com`,                  // (*1)
	`web.davivendacol.website`,          // (*1)
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
