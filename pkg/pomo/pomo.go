package pomo

import (
	"time"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/to"
	"github.com/rwxrob/bonzai/term"
	"github.com/rwxrob/bonzai/dtime"
	"github.com/rwxrob/bonzai/run"
	"github.com/rwxrob/bonzai/persisters/inprops"
)

var (
	Duration   = "30m"
	Warn       = "1m"
	Prefix     = "🍅"
	PrefixWarn = "💢"
	WarnTime   = 5 * time.Minute
)

var Cmd = Z.Cmd{
	Name: `pomo`,
	Cmds: []*Z.Cmd{
		printCmd, // default
		help.Cmd,
		initCmd, startCmd, stopCmd,
	},
	Vars: Z.Vars{
		{
			K:  `duration`,
			S:  `Duration of pomo timer (default: 30m)`,
			P: true,
		},{
			K: `notified`,
			S: `Set to 1 when notification has been sent`,
			P: true,
		},{
			K: `started`,
			S: `Time when pomo was started (RFC3339 format)`,
			P: true,
		},
	},
}.WithPersister(inprops.NewUserCache("x","pomo"))

var initCmd = &Z.Cmd{
	Name:     `init`,
	Short:  `initialize pomo`,
	Cmds: []*Z.Cmd{help.Cmd},

	Do: func(x *Z.Cmd, _ ...string) error {
		return nil
	},
}

var printCmd = &Z.Cmd{
	Name:     `print`,
	Alias:  `show|p`,
	Cmds: []*Z.Cmd{help.Cmd},
	Short:  `print pomo status`,

	Do: func(x *Z.Cmd, _ ...string) error {

		started := x.Caller().Get(`started`)
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

		notified := x.Caller().Get("notified")

		if notified == "" {
			if err := run.Exec("notify-send", "-u", "critical", "Pomo time up"); err != nil {
				return err
			}
			x.Caller().Set("notified", "1")
			return nil
		}

		return nil
	},
}

var startCmd = &Z.Cmd{
	Name:     `start`,
	Usage:    `[help|hour|DURATION]`,
	Cmds: []*Z.Cmd{help.Cmd},
	// Params:   []string{`hour`},
	MaxArgs:  1,

	Do: func(x *Z.Cmd, args ...string) error {
		if len(args) > 0 {
			if args[0] == `hour` {
				t := time.Now()
				args[0] = dtime.Until(dtime.NextHourOf, &t).String()
			}
			x.Caller().Set("duration", args[0])
		}
		s := x.Caller().Get("duration")
		if s == "" {
			s = Duration
		}
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		started := time.Now().Add(dur).Format(time.RFC3339)
		x.Caller().Set("notified","")
		x.Caller().Set("started", started)
		return nil
	},
}

var stopCmd = &Z.Cmd{
	Name:     `stop`,
	Cmds: []*Z.Cmd{help.Cmd},
	Short:  `stop pomo clock`,

	Do: func(x *Z.Cmd, args ...string) error {
		x.Caller().Set("started","")
		x.Caller().Set("notified","")
		return nil
	},
}
