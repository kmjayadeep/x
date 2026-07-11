package net

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "net",
	Short: "network-related utilities",
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "get public IP",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {

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

func init() { Cmd.AddCommand(ipCmd) }
