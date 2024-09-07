package nequi1

import (
	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt, common.ArgsHaveTimes(args...))
}

func attempt() {
}
