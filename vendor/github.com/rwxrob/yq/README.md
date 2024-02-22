# 🌳 Query YAML Data

[![GoDoc](https://godoc.org/github.com/rwxrob/yq?status.svg)](https://godoc.org/github.com/rwxrob/yq)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This is the same popular `yq` tool (written in Go) that uses the same
`yqlib` code but as a [Bonzai](https://github.com/rwxrob/bonzai)
composite command (instead of Cobra) making it more portable and usable
in one's own Bonzai command tree monoliths and multicall binaries along
with other commands. It includes simplified `Evaluate` and
`EvaluateToString` high-level functions calling into the `yqlib`
with reasonable defaults allowing other non-Bonzai applications to
easily duplicate the same functionality.

## Install

This command can be installed as a standalone program or composed into
a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/yq/cmd/yq@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/yq"
)

var Cmd = &Z.Cmd{
	Name:     `cmds`,
	Commands: []*Z.Cmd{help.Cmd, yq.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C yq yq
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
