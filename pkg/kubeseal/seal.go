package kubeseal

import (
	"fmt"
	"os"
	"os/exec"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/vars"
)

func init() {
	Z.Vars.SoftInit()
}

var Cmd = &Z.Cmd{
	Name:    `kubeseal`,
	Summary: `Kubeseal helper commands`,
	Description: `Helper commands instead of kubeseal.
     https://github.com/bitnami-labs/sealed-secrets/
     Set cert path using 'x kubeseal var set kubeseal-cert <path>'`,
	Commands: []*Z.Cmd{help.Cmd, sealCmd, vars.Cmd},
}

var sealCmd = &Z.Cmd{
	Name:     `seal`,
	Summary:  `Seal given secret`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		cert, err := x.Caller.Get(`kubeseal-cert`)
		cmd := exec.Command("kubeseal", "--format=yaml", "--cert="+cert)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Println(string(out))
		return nil
	},
}
