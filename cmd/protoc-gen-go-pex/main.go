package main

import (
	"embed"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	templateFilename        = "template/pex.pb.go.tpl"
	generatedFilenameSuffix = "_pex.pb.go"
)

//go:embed template/*
var efs embed.FS

func main() {
	tmpl, err := template.ParseFS(efs, templateFilename)
	if err != nil {
		panic(err)
	}

	protogen.Options{}.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			filename := f.GeneratedFilenamePrefix + generatedFilenameSuffix
			g := p.NewGeneratedFile(filename, f.GoImportPath)
			if err := tmpl.Execute(g, f); err != nil {
				return err
			}
		}
		return nil
	})
}
