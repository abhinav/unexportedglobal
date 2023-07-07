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
//
// Alternatively, you can use it as a golangci-lint plugin:
//
//	$ go build -buildmode=plugin go.abhg.dev/unexportedglobal/cmd/unexportedglobal
//	$ cat .golangci.yml
//	linter-settings:
//	  custom:
//	    unexportedglobal:
//	      path: unexportedglobal.so
//	      description: Verify unexported globals have an underscore prefix.
//	      original-url: go.abhg.dev/unexportedglobal
package main

import (
	"go.abhg.dev/unexportedglobal"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(unexportedglobal.Analyzer)
}

// AnalyzerPlugin provides the analyzer as a plugin.
var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

// GetAnalyzers returns the unexportedglobal analyzer.
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{unexportedglobal.Analyzer}
}
