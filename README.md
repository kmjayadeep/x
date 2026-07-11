<div align="center">

# `x`

**Small shell utilities, one fast binary.**

[![CI](https://github.com/kmjayadeep/x/actions/workflows/ci.yml/badge.svg)](https://github.com/kmjayadeep/x/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kmjayadeep/x)](https://goreportcard.com/report/github.com/kmjayadeep/x)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`x` collects the tiny commands that are useful enough to keep, but too small to
deserve a separate repository. It provides discoverable help, shell completion,
and a consistent interface without a pile of shell scripts.

</div>

## Highlights

- A single cross-platform Go binary with no runtime service or configuration.
- Built-in help and completion for Bash, Zsh, Fish, and PowerShell.
- Reproducible Nix build and ordinary Go installation.
- Automated dependency updates, CI, and tagged binary releases.
- Persistent command state stored in the operating system's user cache directory.

## Install

### Go

```sh
go install github.com/kmjayadeep/x/cmd/x@latest
```

Make sure `$(go env GOPATH)/bin` is on your `PATH`.

### Nix

Run directly:

```sh
nix run github:kmjayadeep/x -- --help
```

Or install into your profile:

```sh
nix profile install github:kmjayadeep/x
```

### From source

```sh
git clone https://github.com/kmjayadeep/x.git
cd x
make install
```

Tagged releases also include archives for Linux and macOS on AMD64 and ARM64.

## Commands

| Command | Alias | Purpose |
| --- | --- | --- |
| `x clip copy` | `x copy`, `x clip c` | Copy standard input to the clipboard. |
| `x clip paste` | | Print clipboard contents. |
| `x date` | `x d` | Print the current date as `YYYY/MM/DD`. |
| `x date datetime` | | Print the current date and time. |
| `x date full` | | Print a readable short date. |
| `x datehead` | `x dh` | Print a date suitable for a document heading. |
| `x env` | `x env data`, `x env all` | Print the environment. |
| `x env get NAME` | | Print an environment variable, with uppercase fallback. |
| `x git filter code [language]` | | Wrap standard input in a Markdown code block. |
| `x git filter tf` | | Wrap a Terraform plan in a Markdown details block. |
| `x kubeseal [cert-path]` | `x seal` | Seal standard input using a remembered certificate. |
| `x net ip` | | Print the public IP address. |
| `x notes edit` | | Select a note using `fzf` and print its path. |
| `x pomo start [duration]` | | Start a Pomodoro timer; accepts Go durations or `hour`. |
| `x pomo print` | `show`, `p` | Print the remaining Pomodoro time. |
| `x pomo stop` | | Stop the timer. |
| `x weather` | `x weat` | Print compact weather for the current location. |

Every command documents its arguments:

```sh
x --help
x pomo --help
```

## Optional tools and environment

Most commands are self-contained. These commands integrate with external tools:

| Feature | Requirement |
| --- | --- |
| Clipboard | A platform clipboard supported by [`atotto/clipboard`](https://github.com/atotto/clipboard). |
| `kubeseal` | The `kubeseal` executable. |
| Pomodoro notification | `notify-send` on Linux. |
| Notes | `fzf`, `bat`, and `PSUITE_NOTES_DIR` pointing to the notes directory. |

Command state is written beneath `os.UserCacheDir()/x` (usually
`~/.cache/x` on Linux). It contains only the Kubeseal certificate path and
Pomodoro state and can safely be deleted to reset them.

## Shell completion

Generate completion using Cobra's built-in command. For Zsh, for example:

```sh
mkdir -p ~/.zfunc
x completion zsh > ~/.zfunc/_x
```

See `x completion --help` for instructions for your shell.

## Development

Go 1.23.4 or newer is required.

```sh
make check       # format check, vet, and tests
make build       # create ./x
make nix-build   # verify the Nix package
make vendor      # refresh dependencies after changing go.mod
```

To add a command, create a package below `pkg/`, expose a `*cobra.Command`, and
register it in `cmd/x/main.go`. Add focused tests beside the implementation and
run `make check` before opening a pull request.

## Releasing

Push a semantic-version tag to publish a GitHub release:

```sh
git tag v1.2.3
git push origin v1.2.3
```

The release workflow builds checksummed Linux and macOS archives. Regular users
do not need to build or publish releases manually.

## License

[MIT](LICENSE) © Jayadeep KM
