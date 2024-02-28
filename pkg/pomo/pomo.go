package pomo

import (
	"time"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/dtime"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	"github.com/rwxrob/to"
	"github.com/rwxrob/vars"
)

var (
	Duration   = "30m"
	Warn       = "1m"
	Prefix     = "🍅"
	PrefixWarn = "💢"
	WarnTime   = 5 * time.Minute
)

func init() {
	Z.Vars.SoftInit()
}

var Cmd = &Z.Cmd{
	Name: `pomo`,
	Commands: []*Z.Cmd{
		printCmd, // default
		help.Cmd, vars.Cmd,
		initCmd, startCmd, stopCmd,
	},
}

var initCmd = &Z.Cmd{
	Name:     `init`,
	Summary:  `Initialize pomo`,
	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, _ ...string) error {
		return nil
	},
}

var printCmd = &Z.Cmd{
	Name:     `print`,
	Aliases:  []string{`show`, `p`},
	Commands: []*Z.Cmd{help.Cmd},
	Summary:  `Print pomo status`,

	Call: func(x *Z.Cmd, _ ...string) error {

		started, err := x.Caller.Get(`started`)
		if err != nil {
			return err
		}
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
			term.Printf("%v%v", prefix, to.StopWatch(left))
			return nil
		}

		term.Printf("%v%v", prefix, "Pomo up!")

		notified, err := x.Caller.Get("notified")
		if err != nil {
			return err
		}

		if notified == "" {
			if err := Z.Exec("notify-send", "-u", "critical", "Pomo time up"); err != nil {
				return err
			}
			return x.Caller.Set("notified", "1")
		}

		return nil
	},
}

var startCmd = &Z.Cmd{
	Name:     `start`,
	Usage:    `[help|hour|DURATION]`,
	Commands: []*Z.Cmd{help.Cmd},
	Params:   []string{`hour`},
	MaxArgs:  1,

	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) > 0 {
			if args[0] == `hour` {
				t := time.Now()
				args[0] = dtime.Until(dtime.NextHourOf, &t).String()
			}
			if err := x.Caller.Set("duration", args[0]); err != nil {
				return err
			}
		}
		s, err := x.Caller.Get("duration")
		if err != nil {
			return err
		}
		if s == "" {
			s = Duration
		}
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		started := time.Now().Add(dur).Format(time.RFC3339)
		if err := x.Caller.Del("notified"); err != nil {
			return err
		}
		return x.Caller.Set("started", started)
	},
}

var stopCmd = &Z.Cmd{
	Name:     `stop`,
	Commands: []*Z.Cmd{help.Cmd},
	Summary:  `Stop pomo clock`,

	Call: func(x *Z.Cmd, args ...string) error {
		x.Caller.Del("started")
		return x.Caller.Del("notified")
	},
}
