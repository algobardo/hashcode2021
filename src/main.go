package main

import (
	"github.com/alecthomas/kong"
)

var CLI struct {
	Parse struct {
		Print     bool `help:"print"`
		Paths []string `arg name:"path" help:"Paths to parse." type:"path"`
	} `cmd help:"Remove files."`

	DoSomethingElse struct {
		Paths []string `arg optional name:"path" help:"Paths to list." type:"path"`
	} `cmd help:"List paths."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "parse <path>":

	case "do-something-else":
	default:
		panic(ctx.Command())
	}
}
