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

func InitEchoUser() *echo.Echo {
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get users",
			path:                 "/users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"users\":",
		},
		// {
		// 	name:                 "invalid get users",
		// 	path:                 "/users",
		// 	expectStatus:         http.StatusNotFound,
		// 	expectBodyStartsWith: "{\"message\":[",
		// },
	}

	e := InitEchoUser()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			// assert.equal(t, userJSON, rec.Body.String())

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get user",
			path:                 "/users/:id",
			userID:               "2",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"user\":[",
		},
		// {
		// 	name:                 "invalid get user",
		// 	path:                 "/users/:id",
		// userID:               "a12",
		// 	expectStatus:         http.StatusBadRequest,
		// 	expectBodyStartsWith: "{\"message\":[",
		// },
	}

	e := InitEchoUser()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		// Assertions
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			// assert.equal(t, userJSON, rec.Body.String())

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateUserController(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		userJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success create user",
			path:                 "/users/",
			userJSON:             `{"name": "newUser","email": "newUser@gmail.com","password":"newPassword"}`,
			expectStatus:         http.StatusCreated,
			expectBodyStartsWith: "{\"users\":",
		},
	}

	e := InitEchoUser()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, CreateUserController(c)) {
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		userJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success update book",
			path:                 "/users/:id",
			userID:               "1",
			userJSON:             `{"judul": "newJudulUpdated","penulis": "newPenulisUpdated","penerbit":"newPenerbitUpdated"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"users\":",
		},
		// {
		// 	name:                 "success update get book",
		// 	path:                 "/users/:id",
		// userID:               "asd123",
		// 	userJSON:             `{"judul": "newJudulUpdated","penulis": "newPenulisUpdated","penerbit":"newPenerbitUpdated"}`,
		// 	expectStatus:         http.StatusBadRequest,
		// 	expectBodyStartsWith: "{\"users\":",
		// },
	}

	e := InitEchoUser()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

func TestDeleteUserController(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success delete book",
			path:                 "/users/:id",
			userID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"users\":",
		},
		// {
		// 	name:                 "invalid get book",
		// 	path:                 "/users/:id",
		// userID:               "a12",
		// 	expectStatus:         http.StatusBadRequest,
		// 	expectBodyStartsWith: "{\"message\":[",
		// },
	}

	e := InitEchoUser()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		//Assertion
		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			//assert.Equal(t, testCase.expectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
