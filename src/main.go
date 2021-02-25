package main

import (
	"fmt"

	"hashcode2021/m/v2/src/datastructures"
	"hashcode2021/m/v2/src/parser"
	"hashcode2021/m/v2/src/parserutils"
	"hashcode2021/m/v2/src/problemset"
	"hashcode2021/m/v2/src/scorer"
	"hashcode2021/m/v2/src/strategies"

	"github.com/alecthomas/kong"
)

type (
	NaiveStrategy struct {
		Print       bool   `help:"print"`
		Folder      string `arg name:"folder" help:"Folder with problems." type:"string"`
		ProblemName string `arg name:"problem" help:"File with a problem." type:"string"`
	}

	NaiveV2Strategy struct {
		Print       bool   `help:"print"`
		Folder      string `arg name:"folder" help:"Folder with problems." type:"string"`
		ProblemName string `arg name:"problem" help:"File with a problem." type:"string"`
	}

	SimulationStrategy struct {
		Print       bool   `help:"print"`
		Folder      string `arg name:"folder" help:"Folder with problems." type:"string"`
		ProblemName string `arg name:"problem" help:"File with a problem." type:"string"`
	}

	DoSomethingElse struct {
		Paths []string `arg optional name:"path" help:"Paths to list." type:"path"`
	}
)

var CLI struct {
	Naive           NaiveStrategy   `cmd help:"Naive strategy."`
	NaiveV2         NaiveV2Strategy `cmd help:"NaiveV2 strategy."`
	Simul           SimulationStrategy   `cmd help:"Simulation strategy."`
	DoSomethingElse DoSomethingElse `cmd help:"List paths."`
}

func (p *NaiveStrategy) Run() error {
	// ./main naive ~/tmp foo.txt
	ps := problemset.NewProblemSet(p.Folder)
	inputPath := ps.GetProblemInputPath(p.ProblemName)
	lines, err := parserutils.LoadInputAsLines(inputPath)
	if err != nil {
		return err
	}
	input := parser.Parse(lines.Lines)
	//fmt.Println(input.Dumps())
	naiveStrategy := strategies.NewNaiveStrategy()
	output := naiveStrategy.Apply(input).(*datastructures.Solution)
	parserutils.ToStdOut(output.ToStrings())

	println(fmt.Sprintf("Score %d", scorer.Score(input, output)))

	return nil
}

func (p *NaiveV2Strategy) Run() error {
	// ./main naive ~/tmp foo.txt
	ps := problemset.NewProblemSet(p.Folder)
	inputPath := ps.GetProblemInputPath(p.ProblemName)
	lines, err := parserutils.LoadInputAsLines(inputPath)
	if err != nil {
		return err
	}
	input := parser.Parse(lines.Lines)
	//fmt.Println(input.Dumps())
	strategy := strategies.NewNaiveV2Strategy()
	output := strategy.Apply(input)
	parserutils.ToStdOut(output.ToStrings())
	return nil
}

func (p *SimulationStrategy) Run() error {
	// ./main naive ~/tmp foo.txt
	ps := problemset.NewProblemSet(p.Folder)
	inputPath := ps.GetProblemInputPath(p.ProblemName)
	lines, err := parserutils.LoadInputAsLines(inputPath)
	if err != nil {
		return err
	}
	input := parser.Parse(lines.Lines)
	//fmt.Println(input.Dumps())
	s := strategies.NewBackwardsSimulationStrategy()
	output := s.Apply(input)
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
