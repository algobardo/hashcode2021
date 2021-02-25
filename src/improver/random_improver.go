package improver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"runtime"
	"sync"

	"hashcode2021/m/v2/src/parallelexecutor"
	"hashcode2021/m/v2/src/problemset"
)

type Output *struct{}
type Input *struct{}

type Improver struct {
	problems problemset.ProblemSet
	score    func(Input, Output) int
	strategy func(Input) Output
	loader   func(string) Input

	inputs  []Input
	outputs []Output
	scores  []int
	lock    sync.RWMutex
}

func New(problems problemset.ProblemSet,
	score func(Input, Output) int,
	strategy func(Input) Output,
	loader func(string) Input,
) *Improver {
	return &Improver{
		problems: problems,
		score:    score,
		strategy: strategy,
		loader:   loader,
	}
}

func (i *Improver) ImproveAllSolutions() {
	rand.Seed(42)

	for _, name := range i.problems.ListNames() {
		input := i.loader(i.problems.GetProblemInputPath(name))
		output := i.loadBestOutput(i.problems.GetProblemSolutionPath(name))
		if output == nil {
			res := i.strategy(input)
			println(fmt.Sprintf("Computed simulation for %s", name))
			output = res
		}
		score := i.score(input, output)

		i.inputs = append(i.inputs, input)
		i.outputs = append(i.outputs, output)
		i.scores = append(i.scores, score)
	}

	exec := parallelexecutor.New(context.Background(), runtime.GOMAXPROCS(-1))
	for {
		exec.Go(func() error {
			chosenProblem := rand.Intn(len(i.outputs))
			problemName := i.problems.ListNames()[chosenProblem]

			i.lock.RLock()
			curScore := i.scores[chosenProblem]
			curOutput := i.copyOutput(i.outputs[chosenProblem])
			input := i.inputs[chosenProblem]
			i.lock.RUnlock()

			newOutput, newScore := i.tryImprove(input, curOutput, curScore)
			if newOutput != nil {
				i.lock.Lock()
				println(fmt.Sprintf("Improved %s from %d to %d, total: %d", problemName, i.scores[chosenProblem], newScore, i.total(i.scores)))

				i.outputs[chosenProblem] = newOutput
				i.scores[chosenProblem] = newScore
				i.saveBestOutput(newOutput, problemName)
				i.lock.Unlock()
			}
			return nil
		})
	}
}

func (i *Improver) loadBestOutput(problemName string) Output {
	f, err := ioutil.ReadFile(i.problems.GetProblemSolutionImprovedPath(problemName))
	if err != nil {
		return nil
	}

	var output Output
	err = json.Unmarshal(f, &output)
	if err != nil {
		panic("invalid output")
	}
	return output
}

func (i *Improver) saveBestOutput(output Output, problemName string) {
	bytes, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic("failed marshalling output")
	}

	err = ioutil.WriteFile(i.problems.GetProblemSolutionImprovedPath(problemName), bytes, 0775)
	if err != nil {
		panic("failed writing output")
	}
}

func (i *Improver) total(scores []int) int {
	cur := 0
	for _, score := range scores {
		cur += score
	}
	return cur
}

func (i *Improver) tryImprove(input Input, output Output, curScore int) (Output, int) {
	operation := rand.Intn(1)
	switch operation {
	case 0:
		// do something
		print("implement me")
	}
	newScore := i.score(input, output)
	if newScore > curScore {
		return output, newScore
	}
	return nil, -1
}

func (i *Improver) copyOutput(out Output) Output {
	bytes, err := json.Marshal(out)
	if err != nil {
		panic(err.Error())
	}
	var copy Output
	err = json.Unmarshal(bytes, &copy)
	if err != nil {
		panic(err.Error())
	}
	return copy
}
