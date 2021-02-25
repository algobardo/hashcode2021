package parserutils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Lines struct {
	Lines []string
}

func (l Lines) GetNth(n uint) string {
	return l.Lines[n]
}

func (l Lines) GetNthValues(n uint) []string {
	return strings.Split(l.Lines[n], " ")
}

func LoadInputAsLines(filename string) (*Lines, error) {
	s, err := LoadInputAsString(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(s, "\n")
	return &Lines{Lines: lines}, err
}

func LoadInputAsString(filename string) (string, error) {
	bytes, err := LoadInput(filename)
	return string(bytes), err
}

func LoadInput(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func ToStdOut(ls []string) {
	for _, s := range ls {
		fmt.Print(s)
	}
}
