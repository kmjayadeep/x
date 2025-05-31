package weather

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `weather`,
	Alias: `weat`,
	Short:  `get weather based on current location`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden(), basicCmd},
	Def: basicCmd,
}

var basicCmd = &Z.Cmd{
	Name:     `basic`,
	Short:  `basic weather info`,
	Do: func(x *Z.Cmd, args ...string) error {

		url := "https://wttr.in?format=3"

		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		fmt.Print(strings.TrimSpace(string(body)))
		return nil
	},
}
