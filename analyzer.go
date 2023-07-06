// Package unexportedglobal implements a linter
// that verifies that unexported global variables and constants
// in a Go program are prefixed with '_' in their names.
package unexportedglobal

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var _errorInterface *types.Interface

func init() {
	errObj := types.Universe.Lookup("error")
	if errObj == nil {
		panic("no error type in universe scope")
	}

	errIface, ok := errObj.Type().Underlying().(*types.Interface)
	if !ok {
		panic("error type is not an interface")
	}

	_errorInterface = errIface
}

// Analyzer implements the unexportedglobal linter.
//
// See package documentation for details.
var Analyzer = &analysis.Analyzer{
	Name: "unexportedglobal",
	Doc:  "requires a '_' prefix on unexported global variables and constants",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	(&analyzer{
		Fset:   pass.Fset,
		Info:   pass.TypesInfo,
		Report: pass.Report,
	}).Run(inspect)

	return nil, nil
}

type analyzer struct {
	Fset   *token.FileSet
	Info   *types.Info
	Report func(analysis.Diagnostic)
}

var _filter = []ast.Node{
	&ast.GenDecl{},
	&ast.ValueSpec{},

	// We don't actually need to look at function declarations,
	// but catching that node allows us to terminate recursion over its
	// children.
	&ast.FuncDecl{},
}

func (a *analyzer) Run(inspect *inspector.Inspector) {
	inspect.Nodes(_filter, func(n ast.Node, push bool) (proceed bool) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			// A function can't define a global variable.
			return false

		case *ast.GenDecl:
			return n.Tok == token.CONST || n.Tok == token.VAR

		case *ast.ValueSpec:
			a.valueSpec(n)
		}

		return false
	})
}

func (a *analyzer) valueSpec(spec *ast.ValueSpec) {
	for _, name := range spec.Names {
		a.ident(name)
	}
}

func (a *analyzer) ident(id *ast.Ident) {
	name := id.Name
	if ast.IsExported(name) {
		return
	}

	if strings.HasPrefix(name, "_") {
		return
	}

	typ := a.Info.ObjectOf(id).Type()
	if types.Implements(typ, _errorInterface) {
		return
	}

	a.Report(analysis.Diagnostic{
		Pos:     id.NamePos,
		End:     id.End(),
		Message: fmt.Sprintf("unexported global %q should be prefixed with '_'", name),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "add '_' prefix to unexported global",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     id.NamePos,
						End:     id.NamePos,
						NewText: []byte("_"),
					},
				},
			},
		},
	})
}
