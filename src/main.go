package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"hashcode2021/m/v2/src/input"
	"hashcode2021/m/v2/src/parserutils"
	"hashcode2021/m/v2/src/problemset"
	"hashcode2021/m/v2/src/strategies"
)


type (
	NaiveStrategy struct {
		Print     bool `help:"print"`
		Folder string `arg name:"folder" help:"Folder with problems." type:"string"`
		ProblemName string `arg name:"problem" help:"File with a problem." type:"string"`
	}

	DoSomethingElse struct {
		Paths []string `arg optional name:"path" help:"Paths to list." type:"path"`
	}

)

var CLI struct {
	Naive           NaiveStrategy   `cmd help:"Naive strategy."`
	DoSomethingElse DoSomethingElse `cmd help:"List paths."`
}

func (p *NaiveStrategy) Run() error {
	ps := problemset.NewProblemSet(p.Folder)
	inputPath := ps.GetProblemInputPath(p.ProblemName)
	lines, err := parserutils.LoadInputAsLines(inputPath)
	if err != nil {
		return err
	}
	input := input.InputFromLines(lines)
	naiveStrategy := strategies.NewNaiveStrategy()
	output := naiveStrategy.Apply(input)
	parserutils.ToStdOut(output.ToStrings())
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
