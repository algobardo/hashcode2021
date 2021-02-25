package datastructures

import "github.com/davecgh/go-spew/spew"

func (i *Input) String() string {
	return spew.Sdump(i)
}
