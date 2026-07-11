package notes

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "notes",
	Short: "manage notes",
}

func getFile() (string, error) {
	dir := os.Getenv("PSUITE_NOTES_DIR")

	cmd := exec.Command("fzf", "--preview", fmt.Sprintf(`bat --style numbers,changes --color always %s/{}`, dir))
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return dir + "/" + strings.TrimSpace(string(out)), nil
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit notes",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		f, err := getFile()
		if err != nil {
			return err
		}
		fmt.Println(f)
		return nil
	},
}

func init() { Cmd.AddCommand(editCmd) }
