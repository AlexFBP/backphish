package bc01

import (
	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	// POST https://desbloqueo--sucursalvirtua2.repl.co/finish9.php
	// cedula=224224242424

	// POST https://activacion--vitualclave.repl.co/finish9.php
	// clave=6248

	// POST https://dinamica.vitualclave.repl.co/finish9.php
	// clave=245871
}
