package datastructures

import (
	"strconv"
	"strings"
)

type Solution struct {
	Schedules map[IntersectionID]*IntersectionSchedule
}

type IntersectionSchedule struct {
	IntersectionID  IntersectionID
	StreetSchedules map[StreetID]*StreetSchedule
}

type StreetSchedule struct {
	StreetID           StreetID
	GreenLightDuration int
}

func (s *Solution) ToStrings() []string {
	return strings.Split(MarshalToSolution(s), "\n")
}

func MarshalToSolution(solution *Solution) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(len(solution.Schedules)))
	sb.WriteString("\n")
	for _, intersectionSchedule := range solution.Schedules {
		sb.WriteString(strconv.Itoa(int(intersectionSchedule.IntersectionID)))
		sb.WriteString("\n")
		sb.WriteString(strconv.Itoa(len(intersectionSchedule.StreetSchedules)))
		sb.WriteString("\n")
		first := true
		for _, streetSchedule := range intersectionSchedule.StreetSchedules {
			if !first {
				sb.WriteString("\n")
			}
			sb.WriteString(string(streetSchedule.StreetID))
			sb.WriteString(" ")
			sb.WriteString(strconv.Itoa(streetSchedule.GreenLightDuration))
			first = false
		}
	}
	return sb.String()
}
