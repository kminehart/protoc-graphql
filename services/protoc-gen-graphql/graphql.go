package main

import (
	"log"
	"text/template"

	pgsgql "github.com/kminehart/protoc-graphql/services/protoc-gen-graphql/lang/graphql"
	pgs "github.com/lyft/protoc-gen-star"
)

// The GraphqlModule defines the template and context for generating a graphql schema from a protobuf file
type GraphqlModule struct {
	*pgs.ModuleBase
	ctx      pgsgql.Context
	tpl      *template.Template
	rootPath string
}

// Graphql initializes a Graphql module
func Graphql() *GraphqlModule {
	return &GraphqlModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

// InitContext initializes the context with template functions which are passed to the template.
// InitContext also attempts to compile the internal graphql template file, and panics if it fails to compile.
func (p *GraphqlModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgql.InitContext(c.Parameters())

	tpl := template.New("graphql").Funcs(map[string]interface{}{
		"name":     p.ctx.Name,
		"typename": p.ctx.Type,
		"ismap":    p.ctx.IsMap,
	})

	p.tpl = template.Must(tpl.Parse(graphqlTpl))
}

// Name ...
func (p *GraphqlModule) Name() string { return "graphql" }

// Execute runs the compiled template and returns the generated artifacts
func (p *GraphqlModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		log.Println("got target", t)

		p.generate(t)
		for _, i := range t.Imports() {
			log.Println("got target via import", i)
			p.generate(i)
		}
	}

	return p.Artifacts()
}

func (p *GraphqlModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetExt(".graphql")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

const graphqlTpl = `
{{- range .AllEnums }}
enum {{ name . }} {
	{{- range .Values }}
	{{ name . }}
	{{- end }}
}
{{- end }}

{{- range .AllMessages }}
{{- range .NonOneOfFields }}

{{- if ismap . }}
{{ $type := typename . }}
type {{ $type }} {
	{{ $type.Key }}: {{ $type.Element }}
}
{{- end }}

{{- end }}
type {{ name . }} {
	{{- range .NonOneOfFields }}
	{{ name . }}: {{ typename . }}
	{{- end }}
}
{{- end }}
`
