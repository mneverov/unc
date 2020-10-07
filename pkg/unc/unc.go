package unc

import (
	"bytes"
	"go/ast"
	"go/format"
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyser reports c-style nil checks.
var Analyzer = &analysis.Analyzer{
	Name:     "unc",
	Doc:      "Reports c style nil checks.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	fileFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspector.Preorder(fileFilter, func(node ast.Node) {
		be, ok := node.(*ast.BinaryExpr)
		x, ok := be.X.(*ast.Ident)
		if !ok {
			return
		}

		if x.Name == "nil" && x.Obj == nil {
			buf := &bytes.Buffer{}
			if err := format.Node(buf, pass.Fset, be.Y); err != nil {
				log.Fatalf("failed to format new node: %+v", err)
			}

			pass.Report(analysis.Diagnostic{
				Pos:     be.Pos(),
				Message: "lhs nil",
				SuggestedFixes: []analysis.SuggestedFix{{
					Message: "Reorder nil check",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     be.X.Pos(),
							End:     be.X.End(),
							NewText: buf.Bytes(),
						},
						{
							Pos:     be.Y.Pos(),
							End:     be.Y.End(),
							NewText: []byte("nil"),
						},
					},
				}},
			})
		}
	})

	return nil, nil
}
