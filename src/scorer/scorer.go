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
	Now int
}

type Meta struct {
	StreetIntersections map[datastructures.StreetID][]datastructures.IntersectionID
}

func Step(input datastructures.Input, output datastructures.Solution, state State, meta Meta) {
	for _, intersectionScheduled := range output.Schedules {
		remaining := state.GreenRemainingTime[intersectionScheduled.IntersectionID]
		state.GreenRemainingTime[intersectionScheduled.IntersectionID] =  remaining-1

		currentGreenStreet := state.IncomingGreenStreet[intersectionScheduled.IntersectionID]

		// move cars that can be moved, i.e. they are at front of the line
		for _, car := range input.Cars {
			if _, isCompleted := state.CarCompletionTime[car.ID]; !isCompleted {
				newPathPosition := state.CarPositionInItsPath[car.ID]+1
				if canMove(car, input, output, meta) {
					if newPathPosition < len(car.Path) {
						// still some way to go
						state.CarPositionInItsPath[car.ID] = newPathPosition
					} else {
						state.CarCompletionTime[car.ID] = state.Now
					}
				}
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
}

func canMove(car *datastructures.Car, input datastructures.Input, output datastructures.Solution, meta Meta) bool {

}


func Score(input datastructures.Input, output datastructures.Solution) int {

}
