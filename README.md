# workflow-example-go

A minimal Go CLI repository that demonstrates a complete build, install, and release workflow.

This repository is mainly useful as a template for Go projects that need cross-platform release
artifacts, GitHub release publishing, and an install script with optional artifact attestation verification.

## Installation

### Recommended

Install the latest release into `./bin`:

```sh
curl -sSfL https://get.lolico.me/workflow-example-go | sh
```

Install a specific release into `~/.local/bin`:

```sh
curl -fsSL https://get.lolico.me/workflow-example-go | sh -s -- -b ~/.local/bin -d v0.1.0
```

Install script options:

- `-b`: Specify a custom installation directory (defaults to `./bin`)
- `-d`: More verbose logging levels (-d for debug, -dd for trace)
- `-v`: Verify the GitHub attestations of the downloaded artifact before installation (requires `gh` to be installed)
- `-h`: Show install script help

### Manual

You can download from [GitHub Releases Page](https://github.com/c3b2a7/workflow-example-go/releases)

## Usage

Run the binary:

```sh
workflow-example-go
```

## Development

Run tests:

```sh
make test
```

Build all target binaries:

```sh
make all
```

Create release archives:

```sh
make releases
```

Clean build output:

```sh
make clean
```

## Release Assets

`make releases` creates these archives:

| Target        | Asset                                      |
|---------------|--------------------------------------------|
| Linux amd64   | `workflow-example-go-linux-amd64.tar.gz`   |
| Linux arm64   | `workflow-example-go-linux-arm64.tar.gz`   |
| FreeBSD amd64 | `workflow-example-go-freebsd-amd64.tar.gz` |
| FreeBSD arm64 | `workflow-example-go-freebsd-arm64.tar.gz` |
| macOS amd64   | `workflow-example-go-darwin-amd64.tar.gz`  |
| macOS arm64   | `workflow-example-go-darwin-arm64.tar.gz`  |
| Windows amd64 | `workflow-example-go-windows-amd64.zip`    |
| Windows 386   | `workflow-example-go-windows-386.zip`      |

Unix-like archives contain a binary named `workflow-example-go`. Windows
archives contain `workflow-example-go.exe`.

## License

This project is licensed under the [MIT License](LICENSE).
