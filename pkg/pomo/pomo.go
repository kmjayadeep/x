package pomo

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/kmjayadeep/x/internal/store"
	"github.com/spf13/cobra"
)

var props = store.New("pomo.json")

var (
	Duration   = "30m"
	Warn       = "1m"
	Prefix     = "🍅"
	PrefixWarn = "💢"
	WarnTime   = 5 * time.Minute
)

var Cmd = &cobra.Command{
	Use:   "pomo",
	Short: "manage a Pomodoro timer",
	RunE:  printStatus,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize pomo",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}

var printCmd = &cobra.Command{
	Use:     "print",
	Aliases: []string{"show", "p"},
	Short:   "print pomo status",
	Args:    cobra.NoArgs,
	RunE:    printStatus,
}

func printStatus(_ *cobra.Command, _ []string) error {
	started := props.Get("started")
	if started == "" {
		return nil
	}

	endt, err := time.Parse(time.RFC3339, started)
	if err != nil {
		return err
	}

	sec := time.Second
	left := endt.Sub(time.Now()).Round(sec)
	prefix := Prefix

	if left < WarnTime && left%(sec*2) == 0 {
		prefix = PrefixWarn
	}

	if left > 0 {
		fmt.Printf("%s%s", prefix, stopwatch(left))
		return nil
	}

	fmt.Printf("%sPomo up!", prefix)

	notified := props.Get("notified")

	if notified == "" {
		if err := exec.Command("notify-send", "-u", "critical", "Pomo time up").Run(); err != nil {
			return err
		}
		return props.Set("notified", "1")
	}

	return nil
}

var startCmd = &cobra.Command{
	Use:   "start [hour|DURATION]",
	Short: "start the pomo clock",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) > 0 {
			if args[0] == `hour` {
				now := time.Now()
				args[0] = now.Truncate(time.Hour).Add(time.Hour).Sub(now).String()
			}
			if err := props.Set("duration", args[0]); err != nil {
				return err
			}
		}
		s := props.Get("duration")
		if s == "" {
			s = Duration
		}
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		started := time.Now().Add(dur).Format(time.RFC3339)
		if err := props.Set("notified", ""); err != nil {
			return err
		}
		return props.Set("started", started)
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop pomo clock",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		if err := props.Set("started", ""); err != nil {
			return err
		}
		return props.Set("notified", "")
	},
}

func stopwatch(d time.Duration) string {
	d = d.Round(time.Second)
	hours := int(d / time.Hour)
	minutes := int(d % time.Hour / time.Minute)
	seconds := int(d % time.Minute / time.Second)
	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func init() { Cmd.AddCommand(initCmd, printCmd, startCmd, stopCmd) }
