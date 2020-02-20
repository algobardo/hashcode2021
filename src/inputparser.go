package src

type Input struct {
	BooksTotal int
	LibrariesTotal int
	Days int

	BooksScore BooksScore
	Libraries []Library
}

type BooksScore []int

type Library struct {
	DaysForSignUp int
	BooksShippedPerDay int
	Books map[BookID]struct{}
}