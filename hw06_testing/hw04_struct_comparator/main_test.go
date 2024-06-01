package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		book1    *Book
		book2    *Book
		yearWant bool
		sizeWant bool
		rateWant bool
	}{
		{
			name:     "case1",
			book1:    NewBook(1, "Title1", "Author1", 2024, 100, 4.9),
			book2:    NewBook(2, "Title2", "Author2", 2023, 99, 4.5),
			yearWant: true,
			sizeWant: true,
			rateWant: true,
		},
		{
			name:     "case2",
			book1:    NewBook(3, "Title3", "Author3", 2023, 100, 4.5),
			book2:    NewBook(4, "Title4", "Author4", 2024, 99, 4.9),
			yearWant: false,
			sizeWant: true,
			rateWant: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			comparatorByYear := NewBookComparator(ByYear)
			yearGot := comparatorByYear.Compare(tc.book1, tc.book2)

			comparatorBySize := NewBookComparator(BySize)
			sizeGot := comparatorBySize.Compare(tc.book1, tc.book2)

			comparatorByRate := NewBookComparator(ByRate)
			rateGot := comparatorByRate.Compare(tc.book1, tc.book2)

			allMatch := yearGot == tc.yearWant && sizeGot == tc.sizeWant && rateGot == tc.rateWant

			// allMatch должно быть true, иначе выведется текст.
			assert.True(t, allMatch, "All conditions should match")
		})
	}
}
