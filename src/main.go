package main

import (
	"github.com/alecthomas/kong"
	"hashcode2021/m/v2/src/parserutils"
	"hashcode2021/m/v2/src/strategies"
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
		var s strategies.Strategy
		var input strategies.Input
		output := s.Apply(input)
		parserutils.ToStdOut(output.ToStrings())
	case "do-something-else":
	default:
		panic(ctx.Command())
	}
}
