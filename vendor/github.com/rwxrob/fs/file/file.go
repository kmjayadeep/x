package file

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	_fs "io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rogpeppe/go-internal/lockedfile"
	"github.com/rwxrob/fs"
)

// DefaultPerms for new file creation.
var DefaultPerms = 0600
var DefaultDirPerms = 0700

// Touch creates a new file at path or updates the time stamp of
// existing. If a new file is needed creates it with 0600 permissions
// (instead of 0666 as default os.Create does). If the directory does
// not exist it is also created using DefaultDirPerms.
func Touch(path string) error {
	if fs.NotExists(path) {
		if err := os.MkdirAll(
			filepath.Dir(path),
			_fs.FileMode(DefaultDirPerms),
		); err != nil {
			return err
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
		return os.Chmod(path, _fs.FileMode(DefaultPerms))
	}
	now := time.Now().Local()
	if err := os.Chtimes(path, now, now); err != nil {
		return err
	}
	return nil
}

// Replace replaces a file at a specified location with another
// successfully retrieved file from the specified URL or file path and
// duplicates the original files permissions. Only http and https URLs
// are currently supported. For security reasons, no backup copy of the
// replaced executable is kept. Also no checksum validation of the file
// is performed (which is fine in most cases where the connection has
// been secured with HTTPS).
func Replace(orig, url string) error {
	if err := Fetch(url, orig+`.new`); err != nil {
		return err
	}
	if err := fs.DupPerms(orig, orig+`.new`); err != nil {
		return err
	}
	if err := os.Rename(orig, orig+`.orig`); err != nil {
		return err
	}
	if err := os.Rename(orig+`.new`, orig); err != nil {
		return err
	}
	if err := os.Remove(orig + `.orig`); err != nil {
		return err
	}
	return nil
}

// Fetch fetches the specified file at the give "from" URL and saves it
// "to" the specified file path. The name is *not* inferred. If
// timeouts, status, and contexts are required use the net/http package
// instead. Will block until the entire file is downloaded. For more
// involved downloading needs consider the github.com/cavaliercoder/grab
// package.
func Fetch(from, to string) error {

	file, err := os.Create(to)
	defer file.Close()
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Get(from)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf(res.Status)
	}

	if _, err := io.Copy(file, res.Body); err != nil {
		return err
	}

	return nil
}

func execute(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing name of executable")
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Edit opens the file at the given path for editing searching for an
// editor on the system using the following (in order of priority):
//
// * VISUAL
// * EDITOR
// * code
// * vim
// * vi
// * emacs
// * nano
//
func Edit(path string) error {
	ed := os.Getenv("VISUAL")
	if ed != "" {
		return execute(ed, path)
	}
	ed = os.Getenv("EDITOR")
	if ed != "" {
		return execute(ed, path)
	}
	ed, _ = exec.LookPath("code")
	if ed != "" {
		return execute(ed, path)
	}
	ed, _ = exec.LookPath("vim")
	if ed != "" {
		return execute(ed, path)
	}
	ed, _ = exec.LookPath("vi")
	if ed != "" {
		return execute(ed, path)
	}
	ed, _ = exec.LookPath("emacs")
	if ed != "" {
		return execute(ed, path)
	}
	ed, _ = exec.LookPath("nano")
	if ed != "" {
		return execute(ed, path)
	}
	return fmt.Errorf("unable to find editor")
}

// Exists calls fs.Exists and further confirms that the file is a file
// and not a directory.
func Exists(path string) bool { return fs.Exists(path) && !fs.IsDir(path) }

// HereOrAbove returns the full path to the file if the file is found in
// the current working directory, or if not exists in any parent
// directory recursively.
func HereOrAbove(name string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for ; len(dir) > 0 && dir != string(os.PathSeparator); dir = filepath.Dir(dir) {
		path := filepath.Join(dir, name)
		if Exists(path) {
			return path, nil
		}
	}
	return "", nil
}

// Head is like the UNIX head command returning only that number of
// lines from the top of a file.
func Head(path string, n int) ([]string, error) {
	lines := make([]string, 0, n)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	for c := 0; s.Scan() && c < n; c++ {
		lines = append(lines, s.Text())
	}
	return lines, nil
}

// Tail is like the UNIX tail command returning only that number of
// lines from the bottom of a file. If n is negative counts that many
// lines from the top of the file (effectively the line to start from).
func Tail(path string, n int) ([]string, error) {
	var lines []string
	if n >= 0 {
		lines = make([]string, 0, n)
	} else {
		lines = make([]string, 0)
	}
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if n > len(lines) {
		n = len(lines)
	}
	if n < 0 {
		return lines[-n:], nil
	}
	return lines[len(lines)-n:], nil
}

// ReplaceAllString loads the file at path into buffer, compiles the
// regx, and replaces all matches with repl same as function of same
// name overwriting the target file at path. Returns and error if unable
// to compile the regular expression or read or overwrite the file.
//
// Normally, it is better to pre-compile regular expressions. This
// function is designed for applications where the regular expression
// and replacement string are passed by the user at runtime.
func ReplaceAllString(path, regx, repl string) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	exp, err := regexp.Compile(regx)
	if err != nil {
		return err
	}
	return Overwrite(path, exp.ReplaceAllString(string(buf), repl))
}

// FindString reads the file at path into a string, dynamically compiles
// the regx regular expression, and returns FindString on it returning an
// error if file does not exist, or if regular expression could not
// compile. Note that it is not an error to not find the string, which
// causes an empty string to be returned. See regexp.FindString.
func FindString(path, regx string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	exp, err := regexp.Compile(regx)
	if err != nil {
		return "", err
	}
	return exp.FindString(string(buf)), nil
}

// Overwrite replaces the content of the target file at path with the
// string passed using the same file-level locking used by Go. File
// permissions are preserved if file exists.
func Overwrite(path, buf string) error {
	f, err := os.Open(path)
	var mode _fs.FileMode
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := Touch(path); err != nil {
				return err
			}
			f, err = os.Open(path)
			if err := Touch(path); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if !(uint32(mode) > 0) {
		info, err := f.Stat()
		if err != nil {
			return err
		}
		mode = info.Mode()
	}
	return lockedfile.Write(path, strings.NewReader(buf), mode)
}

// Cat just prints the contents of target file to stdout. If the file
// cannot be found or opened returns error. For performance, the entire
// file is loaded into memory before being written to stdout. Do not use
// this on anything where the size of the file is unknown and untrusted.
func Cat(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	os.Stdout.Write(f)
	return nil
}

// IsEmpty checks for files of zero length in an OS-agnostic way. If the
// file does not exist returns false.
func IsEmpty(path string) bool { return Size(path) == 0 }

// Size returns the info.Size() of the file from os.Stat(path). Returns
// -1 if unable to determine.
func Size(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return -1
	}
	return info.Size()
}

// Field returns the field (as returned by strings.Fields) from each
// line of the file located at path (like awk '{print $1}'). Always
// returns a slice even if empty. If that field does not exist on a line,
// that line is omitted. Note that field count starts at 1 (not 0).
func Field(path string, n int) []string {
	fields := make([]string, 0)

	if n <= 0 {
		return fields
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return fields
	}

	s := bufio.NewScanner(bytes.NewReader(f))
	for s.Scan() {
		f := strings.Fields(s.Text())
		if len(f) < n {
			continue
		}
		fields = append(fields, f[n-1])
	}

	return fields
}
