package controllers

import (
	"myAppTesting/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoBook() *echo.Echo {
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get books",
			path:                 "/books",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":[",
		},
		// {
		// 	name:                 "invalid get books",
		// 	path:                 "/bookss",
		// 	expectStatus:         http.StatusNotFound,
		// 	expectBodyStartsWith: "{\"message\":[",
		// },
	}

	e := InitEchoBook()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			// assert.equal(t, userJSON, rec.Body.String())

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		bookID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get book",
			path:                 "/books/:id",
			bookID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":[",
		},
		// {
		// 	name:                 "invalid get book",
		// 	path:                 "/books/:id",
		// 	bookID:               "a12",
		// 	expectStatus:         http.StatusBadRequest,
		// 	expectBodyStartsWith: "{\"message\":[",
		// },
	}

	e := InitEchoBook()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.bookID)

		// Assertions
		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			// assert.equal(t, userJSON, rec.Body.String())

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateBookController(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		bookJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success create book",
			path:                 "/books/",
			bookJSON:             `{"judul": "newJudul","penulis": "newPenulis","penerbit":"newPenerbit"}`,
			expectStatus:         http.StatusCreated,
			expectBodyStartsWith: "{\"books\":",
		},
	}

	e := InitEchoBook()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.bookJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, CreateBookController(c)) {
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

func TestUpdateBookController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		bookID               string
		bookJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success update book",
			path:                 "/books/:id",
			bookID:               "1",
			bookJSON:             `{"judul": "newJudulUpdated","penulis": "newPenulisUpdated","penerbit":"newPenerbitUpdated"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":",
		},
		// {
		// 	name:                 "success update get book",
		// 	path:                 "/books/:id",
		// 	bookID:               "asd123",
		// 	bookJSON:             `{"judul": "newJudulUpdated","penulis": "newPenulisUpdated","penerbit":"newPenerbitUpdated"}`,
		// 	expectStatus:         http.StatusBadRequest,
		// 	expectBodyStartsWith: "{\"books\":",
		// },
	}

	e := InitEchoBook()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.bookJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.bookID)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

func TestDeleteBookController(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		bookID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success delete book",
			path:                 "/books/:id",
			bookID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":",
		},
		// {
		// 	name:                 "invalid get book",
		// 	path:                 "/books/:id",
		// 	bookID:               "a12",
		// 	expectStatus:         http.StatusBadRequest,
		// 	expectBodyStartsWith: "{\"message\":[",
		// },
	}

	e := InitEchoBook()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.bookID)

		//Assertion
		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			//assert.Equal(t, testCase.expectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
