package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/civera17/fintech-test/database"
	"github.com/civera17/fintech-test/models"
	"github.com/stretchr/testify/assert"
)

func TestAllBooksRoute(t *testing.T) {
	mockDb := database.ConnectMockDb()
	// Setup the app as it is done in the main function
	app := SetupAPI()

	expectedBook := []models.Book{
		{
			Title:  "title",
			Author: "Author",
		},
	}

	jsonBook, err := json.Marshal(expectedBook)
	assert.NoError(t, err)

	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "allbooks route",
			route:         "/allbooks",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  string(jsonBook),
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Not Found",
		},
	}

	mockDb.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL`)).WillReturnRows(sqlmock.NewRows([]string{"title", "author"}).
		AddRow(&expectedBook[0].Title, &expectedBook[0].Author))

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.EqualValues(t, test.expectedBody, string(body))
	}
}
