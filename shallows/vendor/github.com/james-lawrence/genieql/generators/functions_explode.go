package generators

import (
	"fmt"
	"go/ast"
	"io"
	"strings"
	"text/template"

	"github.com/james-lawrence/genieql"
	"github.com/james-lawrence/genieql/astutil"
	"github.com/james-lawrence/genieql/internal/drivers"
)

type exploderFunction struct {
	fields []genieql.ColumnMap
	queryFunction
}

func (t exploderFunction) Generate(dst io.Writer) error {
	type context struct {
		Name       string
		Columns    []genieql.ColumnMap
		Parameters []*ast.Field
	}

	return t.queryFunction.Template.Execute(dst, context{
		Columns:    t.fields,
		Name:       t.Name,
		Parameters: t.Parameters,
	})
}

// NewExploderFunction ...
func NewExploderFunction(ctx Context, param *ast.Field, fields []genieql.ColumnMap, options ...QueryFunctionOption) genieql.Generator {
	const defaultQueryFunc = `// {{.Name}} generated by genieql
		func {{.Name}}({{ .Parameters | arguments }}) ([]interface{}, error) {
			var (
				{{- range $index, $column := .Columns }}
				c{{ $index }} {{ $column | sqltype -}} // {{ $column | name -}}
				{{ end }}
			)

			{{ range $index, $field := .Columns }}
			{{- $d := $field | type | typedef -}}
			{{- range $_, $stmt := encode $index $field error -}}
			{{ $stmt | ast }}
			{{ end }}
			{{ end }}
			return []interface{}{{"{"}}{{ .Columns | localvars }}{{"}"}}, nil
		}
		`

	var (
		typedef             = composeTypeDefinitionsExpr(ctx.Driver.LookupType, drivers.DefaultTypeDefinitions)
		defaultQueryFuncMap = template.FuncMap{
			"typedef": typedef,
			"type": func(field genieql.ColumnMap) ast.Expr {
				return ast.NewIdent(field.Definition.Type)
			},
			"sqltype": func(d genieql.ColumnMap) string {
				return d.Definition.ColumnType
			},
			"arguments": argumentsAsPointers,
			"ast":       astPrint,
			"encode":    ColumnMapEncoder(ctx),
			"error": func() func(string) ast.Node {
				return func(local string) ast.Node {
					return astutil.Return(
						ast.NewIdent("[]interface{}(nil)"),
						ast.NewIdent(local),
					)
				}
			},
			"localvars": func(fields []genieql.ColumnMap) (s string) {
				locals := make([]string, 0, len(fields))
				for idx := range fields {
					locals = append(locals, fmt.Sprintf("c%d", idx))
				}
				return strings.Join(locals, ",")
			},
			"name": func(field genieql.ColumnMap) string {
				return field.Name
			},
		}
		tmpl = template.Must(template.New("explode-function").Funcs(defaultQueryFuncMap).Parse(defaultQueryFunc))
	)

	qf := queryFunction{
		Context: ctx,
	}
	qf.Apply(append(options, QFOExplodeStructParam(param, QueryFieldsFromColumnMap(ctx, fields...)...), QFOTemplate(tmpl))...)

	return exploderFunction{
		fields:        fields,
		queryFunction: qf,
	}
}

// QFOExplodeStructParam explodes a structure parameter's fields in the query parameters.
func QFOExplodeStructParam(param *ast.Field, fields ...*ast.Field) QueryFunctionOption {
	selectors := structureQueryParameters(normalizeFieldNames(param)[0], fields...)
	return func(qf *queryFunction) {
		qf.Parameters = append(qf.Parameters, param)
		qf.QueryParameters = append(qf.QueryParameters, selectors...)
	}
}

// StructureQueryParameters - generates QueryParameters for the given struct and its component
// fields.
func StructureQueryParameters(param *ast.Field, fields ...*ast.Field) []ast.Expr {
	return structureQueryParameters(param, fields...)
}

func structureQueryParameters(param *ast.Field, fields ...*ast.Field) []ast.Expr {
	selectors := make([]ast.Expr, 0, len(fields)*len(param.Names))
	for _, name := range param.Names {
		for _, field := range fields {
			selectors = append(selectors, &ast.SelectorExpr{
				X:   name,
				Sel: astutil.MapFieldsToNameIdent(field)[0],
			})
		}
	}

	return selectors
}
