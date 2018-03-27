package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	sqlscanner "github.com/travisjeffery/sqlscannerpb"
)

type Generator struct {
	*generator.Generator
	generator.PluginImports
	overwrite bool
	write     bool
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

func (p *Generator) Write() bool {
	return p.write
}

func (p *Generator) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *Generator) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Messages()) == 0 {
		return
	}
	msgs := p.msgs(file)
	if len(msgs.JSON) > 0 {
		p.PrintImport("json", "json")
	}
	if len(msgs.GoGoProto) > 0 {
		p.PrintImport("gogoproto", "github.com/gogo/protobuf/proto")
	}
}

func (p *Generator) Generate(file *generator.FileDescriptor) {
	t := template.Must(template.New("sqlscanner").Parse(tmpl))
	var buf bytes.Buffer
	t.Execute(&buf, p.msgs(file))
	p.P(buf.String())
}

func (p *Generator) msgs(file *generator.FileDescriptor) Msgs {
	var msgs Msgs
	for _, msg := range file.Messages() {
		if reflect.ValueOf(msg.DescriptorProto.Options).IsNil() {
			continue
		}
		v, err := proto.GetExtension(msg.DescriptorProto.Options, sqlscanner.E_Sqlscanner)
		if err != nil || !p.overwrite {
			continue
		}
		ext := v.(*string)
		switch *ext {
		case "json":
			p.write = true
			msgs.JSON = append(msgs.JSON, msg)
		case "gogoprotobuf":
			p.write = true
			msgs.GoGoProto = append(msgs.GoGoProto, msg)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported marshal type: %s", *ext)
		}
	}
	return msgs
}

func init() {
	generator.RegisterPlugin(NewGenerator())
}

type Msgs struct {
	JSON      []*generator.Descriptor
	GoGoProto []*generator.Descriptor
}

var tmpl = `
{{ range $message := .JSON }}
func (t *{{ $message.Name }}) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

func (t *{{ $message.Name }}) Value() (driver.Value, error) {
	return json.Marshal(t)
}
{{ end }}

{{ range $message := .GoGoProto }}
func (t *{{ $message.Name }}) Scan(val interface{}) error {
	return gogoproto.Unmarshal(val.([]byte), t)
}

func (t *{{ $message.Name }}) Value() (driver.Value, error) {
	return gogoproto.Marshal(t)
}
{{ end }}
`
