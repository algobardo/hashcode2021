package strategies

import (
	"hashcode2021/m/v2/src/datastructures"
	"sort"
)

type NaiveV2 struct {
}

func (n *NaiveV2) Apply(input *datastructures.Input) Output {
	hits := map[datastructures.StreetID]int{}
	for _, car := range input.Cars {
		for _, street := range car.Path {
			hits[street.ID] = hits[street.ID] + 1
		}
	}
	schedules := map[datastructures.IntersectionID]*datastructures.IntersectionSchedule{}
	for intersectionID, intersection := range input.Intersections {
		schedule := &datastructures.IntersectionSchedule{
			IntersectionID:  intersectionID,
			StreetSchedules: map[datastructures.StreetID]*datastructures.StreetSchedule{},
		}

		streets := sortByHits(input, intersection, hits)
		sum := sumByHits(input, intersection, hits)
		for _, street := range streets {
			if hits[street.ID] == 0 {
				continue
			}
			streetID := street.ID
			duration := 1
			if hits[streetID]*len(streets) > sum {
				duration = 2
			}
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

func sortByHits(input *datastructures.Input, intersection *datastructures.Intersection, hits map[datastructures.StreetID]int) []*datastructures.Street {
	var streets []*datastructures.Street
	for _, streetID := range intersection.StreetsIn {
		streets = append(streets, input.GetStreet(streetID))
	}
	sort.Slice(streets, func(i, j int) bool {
		return hits[streets[i].ID] > hits[streets[j].ID]
	})

	for i := 0; i < len(streets); i++ {
		if hits[streets[i].ID] == 0 {
			return streets[:i]
		}
	}
	return streets
}

func sumByHits(input *datastructures.Input, intersection *datastructures.Intersection, hits map[datastructures.StreetID]int) int {
	var streets []*datastructures.Street
	for _, streetID := range intersection.StreetsIn {
		streets = append(streets, input.GetStreet(streetID))
	}

	sum := 0
	for _, street := range streets {
		sum += hits[street.ID]
	}
	return sum
}

func NewNaiveV2Strategy() Strategy {
	return &NaiveV2{}
}

func computeDuration(sum int, hits int, streets int) int {
	duration := 1
	for i := 1; i <= streets; i++ {

	}
	return duration
}
