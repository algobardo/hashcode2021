package scorer

import (
	"fmt"

	"hashcode2021/m/v2/src/datastructures"
)


type State struct {
	CarCompletionTime map[datastructures.CarID]int
	CarPositionInItsPath map[datastructures.CarID]int
	GreenRemainingTime map[datastructures.IntersectionID]int
	IncomingGreenStreet map[datastructures.IntersectionID]datastructures.StreetID
	StreetQueues map[datastructures.StreetID][]datastructures.CarID
	Now int
}

type Meta struct {
	StreetIntersections map[datastructures.StreetID][]datastructures.IntersectionID
}

func setupState(input *datastructures.Input, output *datastructures.Solution) *State {
	s := &State{
		CarCompletionTime:    map[datastructures.CarID]int{},
		CarPositionInItsPath: map[datastructures.CarID]int{},
		GreenRemainingTime:   map[datastructures.IntersectionID]int{},
		IncomingGreenStreet:  map[datastructures.IntersectionID]datastructures.StreetID{},
		StreetQueues:         map[datastructures.StreetID][]datastructures.CarID{},
		Now:                  0,
	}

	for _, street := range input.Streets {
		s.StreetQueues[street.ID] = append([]datastructures.CarID{}, street.Queue...)
	}
	for _, intersectionSchedule := range output.Schedules {
		if len(intersectionSchedule.StreetSchedulesList) > 0 {
			firstSchedule := intersectionSchedule.StreetSchedulesList[0]
			s.IncomingGreenStreet[intersectionSchedule.IntersectionID] = firstSchedule.StreetID
			s.GreenRemainingTime[intersectionSchedule.IntersectionID] = firstSchedule.GreenLightDuration
		}
	}

	return s
}

func Step(input *datastructures.Input, output *datastructures.Solution, state *State) {
	for _, intersectionScheduled := range output.Schedules {
		remaining := state.GreenRemainingTime[intersectionScheduled.IntersectionID]
		state.GreenRemainingTime[intersectionScheduled.IntersectionID] = remaining-1

		currentGreenStreet := state.IncomingGreenStreet[intersectionScheduled.IntersectionID]

		// moving the cars that are front of the queue on the green incoming street
		streetQueue := state.StreetQueues[currentGreenStreet]
		if len(streetQueue) > 0 {
			firstCar := streetQueue[0]
			firstCarInfo := input.Cars[firstCar]
			state.StreetQueues[currentGreenStreet] =streetQueue[1:]
			newPathPosition := state.CarPositionInItsPath[firstCar]+1
			if newPathPosition < len(firstCarInfo.Path) {
				state.CarPositionInItsPath[firstCar] = newPathPosition
				state.StreetQueues[currentGreenStreet] = append(state.StreetQueues[currentGreenStreet],firstCar)
			} else {
				state.CarCompletionTime[firstCar] = state.Now
			}
		}

		// move the semaphore to the next incoming street
		if remaining == 0 {
			intersectionSchedule := output.Schedules[intersectionScheduled.IntersectionID]
			nextSchedule := intersectionSchedule.NextScheduleAfter(currentGreenStreet)
			if nextSchedule.GreenLightDuration == 0 {
				panic(fmt.Sprintf("unexpected 0 green light duration in the solution for street %s at intersection %d", nextSchedule.StreetID, intersectionScheduled.IntersectionID))
			}
			state.GreenRemainingTime[intersectionScheduled.IntersectionID] = nextSchedule.GreenLightDuration
			state.IncomingGreenStreet[intersectionScheduled.IntersectionID] = nextSchedule.StreetID
		}
	}
	state.Now+=1
}

func Score(input *datastructures.Input, output *datastructures.Solution) int {
	currentState := setupState(input,output)
	for i := 0; i < input.Duration; i++ {
		Step(input, output, currentState)
	}
	totalPoints := 0
	for _, completionTime := range currentState.CarCompletionTime {
		totalPoints += input.BonusPoints + (input.Duration - completionTime)
	}
	return totalPoints
}
