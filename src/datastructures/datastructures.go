package datastructures

type BookID int

type LibraryID int

type Input struct {
	BooksTotal     int
	LibrariesTotal int
	Days           int

	BooksScore BooksScore
	Libraries  []*Library
}

type BooksScore []int

type Library struct {
	DaysForSignUp      int
	BooksShippedPerDay int
	Books              []BookID
	BestBooks          []SortedBook
	BestUniqueBooks    []SortedBook
	ID                 LibraryID
}

type SortedBook struct {
	Book  BookID
	Score int
}
