package yq

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/mikefarah/yq/v4/pkg/yqlib"
	logging "gopkg.in/op/go-logging.v1"
)

// NewPrinter is a convenience function for dealing with all the many
// things that must be configured properly to get a usable
// yqlib.Printer. This is useful for testing as well.
func NewPrinter(
	w io.Writer,
	f yqlib.PrinterOutputFormat,
	unwrap bool,
	colors bool,
	indent int,
	sep bool,
) yqlib.Printer {
	prefs := yqlib.YamlPreferences{
		PrintDocSeparators: sep,
		UnwrapScalar:       unwrap,
	}
	enc := yqlib.NewYamlEncoder(indent, colors, prefs)
	pwr := yqlib.NewSinglePrinterWriter(w)
	return yqlib.NewPrinter(enc, pwr)
}

// Evaluate creates a yqlib Evaluator and applies it to one or more
// files with reasonable, predictable defaults for logging, decoding,
// and printing. Only YAML files are supported.  Will read from standard
// input if no arguments are passed.
func Evaluate(expr string, files ...string) error {

	if len(files) == 0 {
		files = append(files, "-")
	}

	// FIXME(rwxrob) shitty log shit, god i HATE op/go-logging (circa 2012)
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05} %{shortfunc} [%{level:.4s}]%{color:reset} %{message}`,
	)
	b1 := logging.NewLogBackend(os.Stderr, "", 0)
	b2 := logging.AddModuleLevel(logging.NewBackendFormatter(b1, format))
	b2.SetLevel(logging.ERROR, "")
	logging.SetBackend(b2)

	ev := yqlib.NewAllAtOnceEvaluator()
	pr := NewPrinter(os.Stdout, yqlib.YamlOutputFormat, true, false, 2, true)
	prefs := yqlib.YamlPreferences{}
	dc := yqlib.NewYamlDecoder(prefs)

	return ev.EvaluateFiles(expr, files, pr, dc)
}

// EvaluateToString is the same as Evaluate but returns a string with
// the output instead.
func EvaluateToString(expr string, files ...string) (string, error) {

	if len(files) == 0 {
		files = append(files, "-")
	}

	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05} %{shortfunc} [%{level:.4s}]%{color:reset} %{message}`,
	)
	b1 := logging.NewLogBackend(os.Stderr, "", 0)
	b2 := logging.AddModuleLevel(logging.NewBackendFormatter(b1, format))
	b2.SetLevel(logging.ERROR, "")
	logging.SetBackend(b2)

	buf := new(bytes.Buffer)

	ev := yqlib.NewAllAtOnceEvaluator()
	pr := NewPrinter(buf, yqlib.YamlOutputFormat, true, false, 2, true)
	prefs := yqlib.YamlPreferences{}
	dc := yqlib.NewYamlDecoder(prefs)

	if err := ev.EvaluateFiles(expr, files, pr, dc); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}
