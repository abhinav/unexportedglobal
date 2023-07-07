# unexportedglobal

- [Introduction](#introduction)
  - [Why](#why)
- [Usage](#usage)
- [FAQ](#faq)
- [License](#license)

## Introduction

[![Go Reference](https://pkg.go.dev/badge/go.abhg.dev/unexportedglobal.svg)](https://pkg.go.dev/go.abhg.dev/unexportedglobal)
[![CI](https://github.com/abhinav/unexportedglobal/actions/workflows/ci.yml/badge.svg)](https://github.com/abhinav/unexportedglobal/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/abhinav/unexportedglobal/branch/main/graph/badge.svg?token=6C90SIVIL4)](https://codecov.io/gh/abhinav/unexportedglobal)

unexportedglobal is a linter for Go that verifies
that all unexported global constants and variables in a Go program
are prefixed with an underscore (`_`).

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
const defaultHost = "example.com"
```

</td><td>

```go
const _defaultHost = "example.com"
```

</td></tr>
</tbody></table>

### Why

The linter is inspired by the guidance,
[Prefix Unexported Globals with `_`](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_)
in the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

Prefixing unexported global variables with an underscore
effectively places them in a namespace separate from normal variables,
eliminating risk of accidentally shadowing or overwriting
such a global variable.

## Usage

To use unexportedglobal from the command line,
first install the standalone program by running:

```bash
go install go.abhg.dev/unexportedglobal/cmd/unexportedglobal@latest
```

Then use it like so:

```bash
unexportedglobal ./...
```

Use it with `go vet` for cleaner output.

```bash
go vet -vettool=$(which unexportedglobal) ./...
```

It is also usable as a `golangci-lint` plugin:

```bash
$ go build -buildmode=plugin go.abhg.dev/unexportedglobal/cmd/unexportedglobal
$ cat .golangci.yml
linter-settings:
  custom:
    unexportedglobal:
      path: unexportedglobal.so
      description: Verify unexported globals have an underscore prefix.
      original-url: go.abhg.dev/unexportedglobal
```

## FAQ

**What about unexported sentinel errors?**

This linter does not require you to rename sentinel error values.
That is, the following will be left alone:

```go
var errNotFound = errors.New("not found")
```

This aligns with the exception carved for this case in the style guide.

## License

This software is made available under the MIT License.
See LICENSE for details.
