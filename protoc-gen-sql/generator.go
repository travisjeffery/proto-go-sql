package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	sql "github.com/travisjeffery/proto-go-sql/sql"
)

type Generator struct {
	*generator.Generator
	generator.PluginImports
	write bool
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (p *Generator) Name() string {
	return "sql"
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
		p.PrintImport("json", "encoding/json")
	}
	if len(msgs.GoGoProto) > 0 {
		p.PrintImport("gogoproto", "github.com/gogo/protobuf/proto")
	}
}

func (p *Generator) Generate(file *generator.FileDescriptor) {
	p.write = false
	t := template.Must(template.New("sql").Parse(tmpl))
	var buf bytes.Buffer
	t.Execute(&buf, p.msgs(file))
	p.P(buf.String())
}

func forEachMessage(parent *descriptor.DescriptorProto, children []*descriptor.DescriptorProto, f func(parent *descriptor.DescriptorProto, child *descriptor.DescriptorProto)) {
	for _, child := range children {
		f(parent, child)
		forEachMessage(child, child.NestedType, f)
	}
}

func (p *Generator) msgs(file *generator.FileDescriptor) Msgs {
	var msgs Msgs

	forEachMessage(nil, file.MessageType, func(parent *descriptor.DescriptorProto, child *descriptor.DescriptorProto) {
		var name string
		if parent != nil {
			parentName := generator.CamelCase(*parent.Name)
			childName := generator.CamelCase(*child.Name)
			name = fmt.Sprintf("%s_%s", parentName, childName)
		} else {
			name = generator.CamelCase(*child.Name)
		}
		child.Name = &name

		if !reflect.ValueOf(child.Options).IsNil() {
			v, err := proto.GetExtension(child.Options, sql.E_All)
			if err == nil {
				ext := v.(*string)

				switch *ext {
				case "json":
					p.write = true
					msgs.JSON = append(msgs.JSON, child)
				case "gogoprotobuf":
					p.write = true
					msgs.GoGoProto = append(msgs.GoGoProto, child)
				default:
					fmt.Fprintf(os.Stderr, "Unsupported marshal type: %s", *ext)
				}
			}
		}
	})

	return msgs
}

func init() {
	generator.RegisterPlugin(NewGenerator())
}

type Msgs struct {
	JSON      []*descriptor.DescriptorProto
	GoGoProto []*descriptor.DescriptorProto
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
