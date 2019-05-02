package main

import (
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	request := command.Read()
	plugin := NewGenerator()
	response := command.GeneratePlugin(request, plugin, "_sql.go")
	if !plugin.Write() {
		return
	}
	command.Write(response)
}
