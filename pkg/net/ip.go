package net

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `net`,
	Short:  `network related utilities`,
	Cmds: []*Z.Cmd{help.Cmd, ipCmd},
}

var ipCmd = &Z.Cmd{
	Name:     `ip`,
	Short:  `get public ip`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {

		url := "https://ipconfig.io"

		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		fmt.Println(strings.TrimSpace(string(body)))
		return nil
	},
}
