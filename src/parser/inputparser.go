package parser

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"hashcode2021/m/v2/src/datastructures"
)

func LoadInput(folder, name string) *datastructures.Input {
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
}

func PostProcess(i *datastructures.Input) *datastructures.Input {
	for _, library := range i.Libraries {
		var sortedBooks []datastructures.SortedBook
		for _, book := range library.Books {
			sortedBooks = append(sortedBooks, datastructures.SortedBook{
				Book:  book,
				Score: i.BooksScore[book],
			})
		}
		library.BestBooks = sortedBooks
		sort.Slice(library.BestBooks, func(i, j int) bool {
			return library.BestBooks[i].Score > library.BestBooks[j].Score
		})
	}
	return i
}

func Parse(s string) *datastructures.Input {
	lines := strings.Split(s, "\n")
	firstLine := strings.Split(lines[0], " ")

	input := &datastructures.Input{
		BooksTotal:     parseInt(firstLine[0]),
		LibrariesTotal: parseInt(firstLine[1]),
		Days:           parseInt(firstLine[2]),
	}

	secondLine := strings.Split(lines[1], " ")
	for _, bookScoreString := range secondLine {
		input.BooksScore = append(input.BooksScore, parseInt(bookScoreString))
	}

	for i := 2; i+1 < len(lines); i += 2 {
		firstLibraryLine := strings.Split(lines[i], " ")
		secondLibraryLine := strings.Split(lines[i+1], " ")

		if len(firstLibraryLine) < 3 {
			continue
		}

		//books := make(map[BookID]struct{}, parseInt(firstLibraryLine[0]))
		var books []datastructures.BookID
		for _, book := range secondLibraryLine {
			books = append(books, datastructures.BookID(parseInt(book)))
		}

		input.Libraries = append(input.Libraries, &datastructures.Library{
			DaysForSignUp:      parseInt(firstLibraryLine[1]),
			BooksShippedPerDay: parseInt(firstLibraryLine[2]),
			Books:              books,
			ID:                 datastructures.LibraryID(len(input.Libraries)),
		})
	}
	return input
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}
