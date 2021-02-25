package parser

import (
	"strconv"
	"strings"

	"hashcode2021/m/v2/src/datastructures"
)

func LoadInput(folder, name string) *datastructures.Input {
	return nil
	/*
		dataFiles, err := ioutil.ReadDir(folder)
		if err != nil {
			panic(err.Error())
		}
		fileName := ""
		for _, file := range dataFiles {
			if strings.HasPrefix(file.Name(), name+"_") && strings.HasSuffix(file.Name(), ".txt") {
				fileName = file.Name()
				break
			}
		}
		if len(fileName) == 0 {
			panic("wrong filename")
		}
		fileContent, err := ioutil.ReadFile(filepath.Join(folder, fileName))
		if err != nil {
			panic(err.Error())
		}

		return PostProcess(Parse(string(fileContent)))
	*/
}

func PostProcess(i *datastructures.Input) *datastructures.Input {
	return i
}

func Parse(lines []string) *datastructures.Input {
	lineNumber := 0
	firstLine := strings.Split(lines[lineNumber], " ")
	lineNumber++
	input := &datastructures.Input{
		Duration:          parseInt(firstLine[0]),
		IntersectionCount: parseInt(firstLine[1]),
		StreetCount:       parseInt(firstLine[2]),
		CarCount:          parseInt(firstLine[3]),
		BonusPoints:       parseInt(firstLine[4]),
	}

	streets := map[datastructures.StreetID]*datastructures.Street{}
	for s := 0; s < input.StreetCount; s++ {
		vals := strings.Split(lines[lineNumber], " ")
		street := &datastructures.Street{
			ID:     datastructures.StreetID(vals[2]),
			Start:  datastructures.IntersectionID(parseInt(vals[0])),
			End:    datastructures.IntersectionID(parseInt(vals[1])),
			Length: parseInt(vals[3]),
			Queue:  []datastructures.CarID{},
		}
		streets[street.ID] = street
		lineNumber++
	}

	var cars []*datastructures.Car
	for c := 0; c < input.CarCount; c++ {
		vals := strings.Split(lines[lineNumber], " ")
		pathLength := parseInt(vals[0])
		car := &datastructures.Car{
			ID: datastructures.CarID(c),
		}

		for p := 0; p < pathLength; p++ {
			streetID := datastructures.StreetID(vals[1+p])
			car.Path = append(car.Path, streets[streetID])
		}
		streets[car.Path[0].ID].Queue = append(streets[car.Path[0].ID].Queue, car.ID)
		cars = append(cars, car)
		lineNumber++
	}

	input.Streets = streets
	input.Cars = cars

	return input
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}
