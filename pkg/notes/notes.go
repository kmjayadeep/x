package notes

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `notes`,
	Short:  `Manage notes`,
	Cmds: []*Z.Cmd{help.Cmd, editCmd},
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

var editCmd = &Z.Cmd{
	Name:     `edit`,
	Short:  `Edit notes`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {
		f, err := getFile()
		if err != nil {
			return err
		}
		fmt.Println(f)
		return nil
	},
}
