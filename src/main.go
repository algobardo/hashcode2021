package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"hashcode2021/m/v2/src/parser"
	"hashcode2021/m/v2/src/problemset"
)


type (
	Parse struct {
		Print     bool `help:"print"`
		Folder string `arg name:"folder" help:"Folder with problems." type:"string"`
		ProblemName string `arg name:"problem" help:"File with a problem." type:"string"`
	}

	DoSomethingElse struct {
		Paths []string `arg optional name:"path" help:"Paths to list." type:"path"`
	}

)

var CLI struct {
	Parse Parse `cmd help:"Remove files."`
	DoSomethingElse DoSomethingElse `cmd help:"List paths."`
}

func (p *Parse) Run() error {
	ps := problemset.NewProblemSet(p.Folder)
	inputPath := ps.GetProblemInputPath(p.ProblemName)
	parser.LoadInput()
	//	var s strategies.Strategy
	//	var input strategies.Input
	//	output := s.Apply(input)
	//	parserutils.ToStdOut(output.ToStrings())
	fmt.Println("rm", p.Path)
	return nil
}

func (p *DoSomethingElse) Run() error {
	fmt.Println("sh", p.Paths)
	return nil
}

func main() {
	ctx := kong.Parse(&CLI)
	ctx.Run()
	//
	//switch ctx.Command() {
	//case "parse <path>":

	//case "do-something-else":
	//default:
	//	panic(ctx.Command())
	//}
}
