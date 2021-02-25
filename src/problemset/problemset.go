package problemset

import (
	"fmt"
	"path"
)

type ProblemSet struct {
	folder string
}

func NewProblemSet(folder string) ProblemSet {
	return ProblemSet{folder: folder}
}

func (p ProblemSet) GetProblemInputPath(name string) string {
	return path.Join(p.folder, name)
}

func (p ProblemSet) GetProblemSolutionPath(name string) string {
	return path.Join(p.folder, name)
}

func (p ProblemSet) GetProblemSolutionJsonPath(name string) string {
	return path.Join(p.folder, fmt.Sprintf("%s.out", name))
}


func (p ProblemSet) GetProblemInputJsonPath(name string) string {
	return path.Join(p.folder, fmt.Sprintf("%s.json", name))
}

