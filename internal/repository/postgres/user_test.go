package postgres_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
	"v001_onelab/configs"
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
	"v001_onelab/internal/repository/postgres"
)

// up test database
const PgURL = "host=localhost port=5436 user=postgres password=password dbname=postgres sslmode=disable"

func TestUser_Create(t *testing.T) {
	conf := &configs.Config{PgURL: PgURL}
	psql, err := postgres.NewDatabasePSQL(conf)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer func() {
		err := ClearDatabase(psql.DB)
		if err != nil {
			t.Fatal(err)
		}
	}()

	repo := repository.Repository{User: postgres.NewUser(psql.DB)}

	testCases := []struct {
		name          string
		inputUser     model.UserInput
		expectedError error
	}{
		{
			name: "OK",
			inputUser: model.UserInput{
				FullName: "testtest",
				Login:    "testtest",
				Password: "testtest",
			},
		},
		{
			name: "ERROR ALREADY EXIST",
			inputUser: model.UserInput{
				FullName: "testtest",
				Login:    "testtest",
				Password: "testtest",
			},
			expectedError: model.ErrorAlreadyExist,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.User.Create(testCase.inputUser)
			if err != nil {
				if testCase.expectedError == nil {
					t.Fatal(err.Error())
				} else {
					require.Error(t, testCase.expectedError, err.Error())
					return
				}
			}
			require.NoError(t, err)
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	conf := &configs.Config{PgURL: PgURL}
	psql, err := postgres.NewDatabasePSQL(conf)
	if err != nil {
		t.Fatal(err.Error())
	}

	repo := repository.Repository{User: postgres.NewUser(psql.DB)}

	testTable := []struct {
		name          string
		inputID       int
		expected      model.UserResponse
		expectedError error
	}{
		{
			name:    "OK",
			inputID: 1,
			expected: model.UserResponse{
				ID: 1,
				//this user got when I make migration
				FullName: "Dias Serikov",
				Login:    "qara-qurt",
			},
		},
		{
			name:          "ERROR USER NOT FOUND",
			inputID:       993,
			expectedError: model.ErrorNotFound,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := repo.User.GetByID(testCase.inputID)
			if err != nil {
				if testCase.expectedError == nil {
					t.Fatal(err.Error())
				} else {
					require.Error(t, testCase.expectedError, err.Error())
				}
			}
			assert.Equal(t, res, testCase.expected)
		})
	}
}

func ClearDatabase(psql *sqlx.DB) error {
	_, err := psql.DB.Exec("DELETE FROM users WHERE login='testtest'")
	if err != nil {
		return err
	}
	return nil
}
