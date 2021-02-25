package datastructures

import (
	"strconv"
)

type Solution struct {
	Schedules map[IntersectionID]*IntersectionSchedule
}

type IntersectionSchedule struct {
	IntersectionID      IntersectionID
	StreetSchedules     map[StreetID]*StreetSchedule
	StreetSchedulesList []*StreetSchedule
}

type StreetSchedule struct {
	StreetID           StreetID
	GreenLightDuration int
}

func (s *Solution) ToStrings() []string {
	var lines []string
	lines = append(lines, strconv.Itoa(len(s.Schedules)))
	for _, intersectionSchedule := range s.Schedules {
		lines = append(lines, strconv.Itoa(int(intersectionSchedule.IntersectionID)))
		lines = append(lines, strconv.Itoa(len(intersectionSchedule.StreetSchedules)))
		for _, streetSchedule := range intersectionSchedule.StreetSchedulesList {
			lines = append(lines, string(streetSchedule.StreetID))
			lines = append(lines, strconv.Itoa(streetSchedule.GreenLightDuration))
		}
	}
	return lines
}
