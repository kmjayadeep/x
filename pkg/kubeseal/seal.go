package kubeseal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kmjayadeep/x/internal/store"
	"github.com/spf13/cobra"
)

var props = store.New("kubeseal.json")

var Cmd = &cobra.Command{
	Use:     "kubeseal [cert-path]",
	Short:   "kubeseal helper commands",
	Aliases: []string{"seal"},
	Long: `Helper commands for kubeseal.
     https://github.com/bitnami-labs/sealed-secrets/
     Cert path is cached after first use`,
	Args: cobra.MaximumNArgs(1),
	RunE: seal,
}

var sealCmd = &cobra.Command{
	Use:   "seal [cert-path]",
	Short: "seal the secret read from stdin",
	Args:  cobra.MaximumNArgs(1),
	RunE:  seal,
}

func seal(_ *cobra.Command, args []string) error {
	if len(args) > 0 {
		if err := props.Set("kubeseal-cert", args[0]); err != nil {
			return err
		}
	}

	cert := props.Get(`kubeseal-cert`)
	if cert == "" {
		return fmt.Errorf("no cert path provided; pass the path to the kubeseal certificate")
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
}

func init() { Cmd.AddCommand(sealCmd) }
