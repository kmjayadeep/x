package yq

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/compfile"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{

	Name:      `yq`,
	Summary:   `query YAML and JSON files`,
	Usage:     `(help|<expression>) [<file>...]`,
	Comp:      compfile.New(),
	Version:   `v0.3.2`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands:  []*Z.Cmd{help.Cmd},

	Description: `
		The {{cmd .Name}} command allows YAML (and JSON, since all JSON is YAML)
		files to be queried using a simple syntax that is nearly identical
		to the popular stedolan/jq tool (written in C). In fact, {{cmd .Name}} uses
		the same *yqlib* that the {{cmd .Name}} tool does (just without the Cobra).

		The first argument is the <expression> and almost always begins with
		a dot (.). See the <https://github.com/stedolan/jq> project for
		exact syntax (until more documentation is updated here to contain
		the same).

		The remaining arguments are the names of one or more files. If no
		file argument is passed standard input is assumed (per UNIX filter
		philosophy). Note that the special dash (-) filename is not
		supported even though it was in the original {{cmd .Name}} tool.`,

	Call: func(x *Z.Cmd, args ...string) error {
		var files []string
		switch len(args) {
		case 1:
			files = append(files, "-")
		case 0:
			return x.UsageError()
		default:
			files = append(files, args...)
		}
		return Evaluate(args[0], args[1:]...)
	},
}
