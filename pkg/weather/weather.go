package weather

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "weather",
	Aliases: []string{"weat"},
	Short:   "get weather based on current location",
	RunE:    getWeather,
}

var basicCmd = &cobra.Command{
	Use:   "basic",
	Short: "basic weather info",
	Args:  cobra.NoArgs,
	RunE:  getWeather,
}

func getWeather(_ *cobra.Command, _ []string) error {

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
}

func init() { Cmd.AddCommand(basicCmd) }
