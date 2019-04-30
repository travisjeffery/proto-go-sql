package main

import (
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	request := command.Read()
	plugin := NewGenerator()
	response := command.GeneratePlugin(request, plugin, "_sql.go")
	command.Write(response)
}
