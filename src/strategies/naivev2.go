package strategies

import (
	"hashcode2021/m/v2/src/datastructures"
	"sort"
)

type NaiveV2 struct {
}

func (n *NaiveV2) Apply(input *datastructures.Input) Output {
	isHit := map[datastructures.StreetID]bool{}
	for _, car := range input.Cars {
		for _, street := range car.Path {
			isHit[street.ID] = true
		}
	}
	schedules := map[datastructures.IntersectionID]*datastructures.IntersectionSchedule{}
	for intersectionID, intersection := range input.Intersections {
		schedule := &datastructures.IntersectionSchedule{
			IntersectionID:  intersectionID,
			StreetSchedules: map[datastructures.StreetID]*datastructures.StreetSchedule{},
		}

		streets := sortByQueue(input, intersection)
		for _, street := range streets {
			if !isHit[street.ID] {
				continue
			}
			streetID := street.ID
			duration := 1
			streetSchedule := &datastructures.StreetSchedule{
				StreetID:           streetID,
				GreenLightDuration: duration,
			}
			schedule.StreetSchedules[streetID] = streetSchedule
			schedule.StreetSchedulesList = append(schedule.StreetSchedulesList, streetSchedule)
		}
		if len(schedule.StreetSchedules) > 0 {
			schedules[intersectionID] = schedule
		}

	}
	return &datastructures.Solution{
		Schedules: schedules,
	}
}

func sumLengths(input *datastructures.Input, intersection *datastructures.Intersection) int {
	var streets []*datastructures.Street
	for _, streetID := range intersection.StreetsIn {
		streets = append(streets, input.GetStreet(streetID))
	}

	result := 0
	for _, street := range streets {
		result = result + street.Length
	}
	return result
}

func sortByLength(input *datastructures.Input, intersection *datastructures.Intersection) []*datastructures.Street {
	var streets []*datastructures.Street
	for _, streetID := range intersection.StreetsIn {
		streets = append(streets, input.GetStreet(streetID))
	}

	sort.Slice(streets, func(i, j int) bool {
		return streets[i].Length < streets[j].Length
	})
	return streets
}

func sortByQueue(input *datastructures.Input, intersection *datastructures.Intersection) []*datastructures.Street {
	var streets []*datastructures.Street
	for _, streetID := range intersection.StreetsIn {
		streets = append(streets, input.GetStreet(streetID))
	}
	sort.Slice(streets, func(i, j int) bool {
		return len(streets[i].Queue) < len(streets[j].Queue)
	})
	return streets
}

func NewNaiveV2Strategy() Strategy {
	return &NaiveV2{}
}
