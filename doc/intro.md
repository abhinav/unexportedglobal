# Introduction

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
