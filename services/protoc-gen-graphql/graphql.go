package main

import (
	"text/template"

	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	pgs "github.com/lyft/protoc-gen-star"
)

type GraphqlModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func Graphql() *GraphqlModule {
	return &GraphqlModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (p *GraphqlModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("graphql").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
	})

	p.tpl = template.Must(tpl.Parse(graphqlTpl))
}

func (p *GraphqlModule) Name() string { return "graphql" }

func (p *GraphqlModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
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
{{ range .AllMessages }}
type {{ name . }} {
}
{{ end }}
`
