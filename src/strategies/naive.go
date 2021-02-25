package strategies

import "hashcode2021/m/v2/src/datastructures"

type Naive struct {
}

func (n *Naive) Apply(input Input) Output {
	return &datastructures.Solution{
		Schedules: map[datastructures.IntersectionID]*datastructures.IntersectionSchedule{},
	}
	//for _, x := range input. {
	//
	//}
}

func NewNaiveStrategy() Strategy {
	return &Naive{}
}
