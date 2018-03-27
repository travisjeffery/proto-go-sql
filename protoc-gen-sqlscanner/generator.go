package main

import (
	"bytes"
	"html/template"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/proto"
	google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type Generator struct {
	*generator.Generator
	generator.PluginImports
	overwrite bool
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (p *Generator) Name() string {
	return "sqlscanner"
}

func (p *Generator) Overwrite() {
	p.overwrite = true
}

func (p *Generator) Init(g *generator.Generator) {
	p.Generator = g
}

var tmpl = `
{{ range $message := . }}
func (t *{{ $message.Name }}) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

func (t *{{ $message.Name }}) Value() (driver.Value, error) {
	return json.Marshal(t)
}
{{ end }}
`

func (p *Generator) GenerateImports(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	p.P()
}

func (p *Generator) Generate(file *generator.FileDescriptor) {
	var messages []*generator.Descriptor
	for _, message := range file.Messages() {
		if !p.overwrite && proto.GetBoolExtension(message.Options, Extension, false) {
			continue
		}
		messages = append(messages, message)
	}
	t := template.Must(template.New("sqlscanner").Parse(tmpl))
	var buf bytes.Buffer
	t.Execute(&buf, messages)
	spew.Fdump(os.Stderr, "buf:", buf.String())
	p.P(buf.String())
}

var Extension = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         90000,
	Name:          "sqlscanner",
	Tag:           "varint,90000,opt,name=sqlscanner",
	Filename:      "generator.go",
}

func init() {
	proto.RegisterExtension(Extension)
	generator.RegisterPlugin(NewGenerator())
}
