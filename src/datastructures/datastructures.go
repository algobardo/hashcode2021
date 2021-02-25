package datastructures

import "encoding/json"

type StreetID string

type IntersectionID int

type CarID int

type Car struct {
	ID       CarID
	Path     []*Street
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
	StreetsIn  []StreetID
	StreetsOut []StreetID
}

type Input struct {
	Duration          int
	IntersectionCount int
	StreetCount       int
	CarCount          int
	BonusPoints       int

	Cars          map[CarID]*Car
	Streets       map[StreetID]*Street
	Intersections map[IntersectionID]*Intersection
}

func (i *Input) Dumps() string {
	bs, _ := json.MarshalIndent(i, "", "  ")
	return string(bs)
}

func (i *Input) GetStreet(id StreetID) *Street {
	return i.Streets[id]
}

func (i *Input) GetIntersection(id IntersectionID) *Intersection {
	return i.Intersections[id]
}
