package weather

import (
	"fmt"
	"io"
	"net/http"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `weather`,
	Short:  `Get weather based on current location`,
	Cmds: []*Z.Cmd{help.Cmd, basicCmd},
}

var basicCmd = &Z.Cmd{
	Name:     `basic`,
	Short:  `Basic weather info`,
	Cmds: []*Z.Cmd{help.Cmd},
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
		fmt.Println(string(body))
		return nil
	},
}
