package analyzer

import (
	"fmt"
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "importpackagename",
	Doc:      "Check package import naming",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspector.Preorder([]ast.Node{(*ast.ImportSpec)(nil)},
		func(node ast.Node) {
			spec := node.(*ast.ImportSpec)

			name := importName(spec)
			path := importPath(spec)

			switch name {
			case ".", "_":
				return
			}

			wanted := Config.AliasForPath(path)
			switch wanted {
			case "":
				return // No preferred name mapping configured.
			case name:
				return // Alias is already correct.
			}

			switch name {
			case "":
				// TODO(jpeach) The import doesn't have a package name,
				// we want to insert one before the import path.
			default:
				pass.Report(analysis.Diagnostic{
					Pos:     spec.Pos(),
					Message: fmt.Sprintf("import package name %q should be %q", name, wanted),
					SuggestedFixes: []analysis.SuggestedFix{{
						Message: fmt.Sprintf("replace %q with %q", name, wanted),
						TextEdits: []analysis.TextEdit{
							{
								Pos:     spec.Name.Pos(),
								End:     spec.Name.End(),
								NewText: []byte(wanted),
							},
						},
					}},
				})
			}

			// XXX(jpeach) since we suggested a fix that changed
			// the import package name, we ought to also suggest
			// fixes for all the qualified identifiers that use
			// that name.
		})

	return nil, nil
}

func importName(spec *ast.ImportSpec) string {
	if spec != nil && spec.Name != nil {
		return spec.Name.Name
	}

	return ""
}

func importPath(spec *ast.ImportSpec) string {
	if spec != nil && spec.Path != nil {
		s, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			panic(err.Error())
		}

		return s
	}

	return ""
}
