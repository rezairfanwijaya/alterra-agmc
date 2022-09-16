package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	type args struct {
		name string
		path string
		body string
	}

	testCase := args{
		name: "Success",
		path: "/books",
		body: "{\"data\"",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testCase.path)

	if assert.NoError(t, GetAllBook(c)) {
		body := rec.Body.String()
		assert.True(t, strings.HasPrefix(body, testCase.body))
	}
}

func TestBookById(t *testing.T) {
	type args struct {
		name       string
		path       string
		body       string
		paramName  string
		paramValue string
	}

	testCases := []args{
		{
			name:       "Success",
			path:       "/book:id",
			body:       "{\"data\":{\"id\":2",
			paramName:  "id",
			paramValue: "2",
		},
		{
			name:       "Book Not Found",
			path:       "/book:id",
			body:       "{\"data\":{\"id\":0",
			paramName:  "id",
			paramValue: "90",
		},
		{
			name:       "Invalid Id",
			path:       "/book:id",
			body:       "",
			paramName:  "id",
			paramValue: "jdbdjfb",
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames(testCase.paramName)
			c.SetParamValues(testCase.paramValue)

			GetBookById(c)
			body := rec.Body.String()
			log.Println("body" , body)
			expectation := strings.HasPrefix(body, testCase.body)
			assert.True(t, expectation)

		})

	}

}
