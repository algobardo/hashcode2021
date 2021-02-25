package datastructures

import "encoding/json"

type StreetID string

type IntersectionID int

type CarID int

type Car struct {
	ID       CarID
	Path     []*Street
	Position int
}

type Street struct {
	ID     StreetID
	Start  IntersectionID
	End    IntersectionID
	Length int
	Queue  []CarID
}

type Intersection struct {
	ID         IntersectionID
	StreetsIn  []*Street
	StreetsOut []*Street
}

type Input struct {
	Duration          int
	IntersectionCount int
	StreetCount       int
	CarCount          int
	BonusPoints       int

	Cars    []*Car
	Streets map[StreetID]*Street
}

func (i *Input) Dumps() string {
	bs, _ := json.MarshalIndent(i, "", "  ")
	return string(bs)
}
