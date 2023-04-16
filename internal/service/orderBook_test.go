package service

import (
	"database/sql"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
	"v001_onelab/internal/model"
)

func TestCombineBooks(t *testing.T) {
	date := time.Date(2023, 4, 16, 10, 0, 0, 0, time.UTC)
	testTable := []struct {
		name         string
		withReturned bool
		input        []model.UserOrderBooks
		expected     []model.UserOrderBooksResponse
	}{
		{
			name:         "OK without returned",
			withReturned: true,
			input: []model.UserOrderBooks{
				{
					ID:       1,
					FullName: "John Doe",
					Login:    "johndoe",
					Book: model.BookWithDate{
						ID:          1,
						Name:        "Book 1",
						Description: "Description 1",
						Author:      "Author 1",
						OrderDate:   date,
						ReturnDate:  sql.NullTime{},
					},
				},
				{
					ID:       1,
					FullName: "John Doe",
					Login:    "johndoe",
					Book: model.BookWithDate{
						ID:          2,
						Name:        "Book 2",
						Description: "Description 2",
						Author:      "Author 2",
						OrderDate:   date,
						ReturnDate: sql.NullTime{
							Time:  date.Add(24 * time.Hour),
							Valid: true,
						},
					},
				},
				{
					ID:       2,
					FullName: "Jane Smith",
					Login:    "janesmith",
					Book: model.BookWithDate{
						ID:          3,
						Name:        "Book 3",
						Description: "Description 3",
						Author:      "Author 3",
						OrderDate:   date,
						ReturnDate:  sql.NullTime{},
					},
				},
			},
			expected: []model.UserOrderBooksResponse{
				{
					ID:       1,
					FullName: "John Doe",
					Login:    "johndoe",
					Book: []model.BookWithDate{
						{
							ID:          1,
							Name:        "Book 1",
							Description: "Description 1",
							Author:      "Author 1",
							OrderDate:   date,
							ReturnDate:  sql.NullTime{},
						},
						{
							ID:          2,
							Name:        "Book 2",
							Description: "Description 2",
							Author:      "Author 2",
							OrderDate:   date,
							ReturnDate: sql.NullTime{
								Time:  date.Add(24 * time.Hour),
								Valid: true,
							},
						},
					},
				},
				{
					ID:       2,
					FullName: "Jane Smith",
					Login:    "janesmith",
					Book: []model.BookWithDate{
						{
							ID:          3,
							Name:        "Book 3",
							Description: "Description 3",
							Author:      "Author 3",
							OrderDate:   date,
							ReturnDate:  sql.NullTime{},
						},
					},
				},
			},
		},
		{
			name:         "OK with returned",
			withReturned: false,
			input: []model.UserOrderBooks{
				{
					ID:       1,
					FullName: "John Doe",
					Login:    "johndoe",
					Book: model.BookWithDate{
						ID:          1,
						Name:        "Book 1",
						Description: "Description 1",
						Author:      "Author 1",
						OrderDate:   date,
						ReturnDate:  sql.NullTime{},
					},
				},
				{
					ID:       1,
					FullName: "John Doe",
					Login:    "johndoe",
					Book: model.BookWithDate{
						ID:          2,
						Name:        "Book 2",
						Description: "Description 2",
						Author:      "Author 2",
						OrderDate:   date,
						ReturnDate: sql.NullTime{
							Time:  date.Add(24 * time.Hour),
							Valid: true,
						},
					},
				},
				{
					ID:       2,
					FullName: "Jane Smith",
					Login:    "janesmith",
					Book: model.BookWithDate{
						ID:          3,
						Name:        "Book 3",
						Description: "Description 3",
						Author:      "Author 3",
						OrderDate:   date,
						ReturnDate:  sql.NullTime{},
					},
				},
			},
			expected: []model.UserOrderBooksResponse{
				{
					ID:       1,
					Login:    "johndoe",
					FullName: "John Doe",
					Book: []model.BookWithDate{
						{
							ID:          1,
							Name:        "Book 1",
							Description: "Description 1",
							Author:      "Author 1",
							OrderDate:   date,
							ReturnDate:  sql.NullTime{},
						},
					},
				},
				{
					ID:       2,
					FullName: "Jane Smith",
					Login:    "janesmith",
					Book: []model.BookWithDate{
						{
							ID:          3,
							Name:        "Book 3",
							Description: "Description 3",
							Author:      "Author 3",
							OrderDate:   date,
							ReturnDate:  sql.NullTime{},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			res := combineBooks(testCase.input, testCase.withReturned)

			for i, r := range res {
				assert.Equal(t, r.ID, testCase.expected[i].ID)
				assert.Equal(t, r.Book, testCase.expected[i].Book)
				assert.Equal(t, r.FullName, testCase.expected[i].FullName)
				assert.Equal(t, r.Login, testCase.expected[i].Login)
			}
		})
	}
}
