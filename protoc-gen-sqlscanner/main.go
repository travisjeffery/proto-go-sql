package main

import (
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	p := NewGenerator()
	resp := command.GeneratePlugin(req, p, "_sqlscanner.go")
	if !p.Write() {
		return
	}
	command.Write(resp)
}
