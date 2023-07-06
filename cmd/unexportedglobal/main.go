// unexportedglobal is a linter for Go code
// that verifies that unexported global variables and constants
// are prefixed with '_' in their names.
//
// To use this linter, run the 'unexportedglobal' binary directly:
//
//	$ unexportedglobal ./...
//
// Alternatively, you can use the 'go vet' command:
//
//	$ go vet -vettool=$(which unexportedglobal) ./...
package main

import (
	"go.abhg.dev/unexportedglobal"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(unexportedglobal.Analyzer)
}
