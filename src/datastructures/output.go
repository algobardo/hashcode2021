package datastructures

import (
	"strconv"
	"strings"
)

type Solution struct {
	Schedules map[IntersectionID]*IntersectionSchedule
}

type IntersectionSchedule struct {
	IntersectionID      IntersectionID
	StreetSchedules     map[StreetID]*StreetSchedule
	StreetSchedulesList []*StreetSchedule
}

func (s *IntersectionSchedule) NextScheduleAfter(street StreetID) *StreetSchedule {
	idx := -1
	for i, streetSchedule := range s.StreetSchedulesList {
		if streetSchedule.StreetID == street {
			idx = i
			break
		}
	}
	return s.StreetSchedulesList[(idx+1)%len(s.StreetSchedulesList)]
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
			lines = append(lines)
			elements := []string{
				string(streetSchedule.StreetID),
				strconv.Itoa(streetSchedule.GreenLightDuration),
			}
			lines = append(lines, strings.Join(elements, " "))
		}
	}
	return lines
}
