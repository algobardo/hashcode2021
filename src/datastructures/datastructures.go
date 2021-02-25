package datastructures

type StreetID string

type IntersectionID int

type CarID int

type Car struct {
	ID     CarID
	Path   []*Street
	Street int
}

type Street struct {
	ID     StreetID
	Start  *Intersection
	End    *Intersection
	Length int
	Queue  []*Car
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
}
