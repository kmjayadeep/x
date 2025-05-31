package kubeseal

import (
	"fmt"
	"os"
	"os/exec"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/persisters/inprops"
	"github.com/rwxrob/bonzai/term"
	"github.com/rwxrob/bonzai/vars"
)

var props = inprops.NewUserCache("x", "kubeseal.props")

var Cmd = &Z.Cmd{
	Name:    `kubeseal`,
	Short: `kubeseal helper commands`,
	Alias: `seal`,
	Usage: `kubeseal [cert-path]`,
	Long: `Helper commands instead of kubeseal.
     https://github.com/bitnami-labs/sealed-secrets/
     Cert path is cached after first use`,
	Def: sealCmd,
	Cmds: []*Z.Cmd{help.Cmd, sealCmd, vars.Cmd},
}

var sealCmd = &Z.Cmd{
	Name:     `seal`,
	Short:  `Seal given secret`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(x *Z.Cmd, args ...string) error {
		if len(args) > 0 {
			props.Set(`kubeseal-cert`, args[0])
		}

		cert := props.Get(`kubeseal-cert`)
		if cert == "" {
			_ ,err := term.Print(`No cert path provided. Please provide the path to the kubeseal certificate.`)
			return err
		}

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
