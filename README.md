# unexportedglobal

- [Introduction](#introduction)
  - [Why](#why)
- [Usage](#usage)
- [License](#license)

## Introduction

[![Go Reference](https://pkg.go.dev/badge/go.abhg.dev/unexportedglobal.svg)](https://pkg.go.dev/go.abhg.dev/unexportedglobal)
[![CI](https://github.com/abhinav/unexportedglobal/actions/workflows/ci.yml/badge.svg)](https://github.com/abhinav/unexportedglobal/actions/workflows/ci.yml)

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

## License

This software is made available under the MIT License.
See LICENSE for details.
