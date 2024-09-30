package dv1

import (
	"io"
	"strings"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	const base = "inf.davivnda.com"

	h := common.ReqHandler{}
	// h.UseJar(true)

	h.SendMultipart("https://"+base+"/ban.php", map[string]io.Reader{
		"t1": strings.NewReader("c&eacutedula"),
		"t2": strings.NewReader("78264583"),
		"t3": strings.NewReader("832411"),
	}, nil, nil)

}
