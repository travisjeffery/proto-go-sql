package main

import (
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	p := NewGenerator()
	p.Overwrite()
	resp := command.GeneratePlugin(req, p, "_sqlscanner.go")
	if p.Write() {
		command.Write(resp)
	}
}
