package strategies

import "hashcode2021/m/v2/src/datastructures"

type Naive struct {
}

func (n *Naive) Apply(input *datastructures.Input) Output {
	schedules := map[datastructures.IntersectionID]*datastructures.IntersectionSchedule{}
	for intersectionID, intersection := range input.Intersections {
		schedule := &datastructures.IntersectionSchedule{
			IntersectionID:  intersectionID,
			StreetSchedules: map[datastructures.StreetID]*datastructures.StreetSchedule{},
		}

		for _, streetID := range intersection.StreetsIn {
			streetSchedule := &datastructures.StreetSchedule{
				StreetID:           streetID,
				GreenLightDuration: 1,
			}
			schedule.StreetSchedules[streetID] = streetSchedule
			schedule.StreetSchedulesList = append(schedule.StreetSchedulesList, streetSchedule)
		}
		schedules[intersectionID] = schedule
	}
	return &datastructures.Solution{
		Schedules: schedules,
	}
}

func NewNaiveStrategy() Strategy {
	return &Naive{}
}
