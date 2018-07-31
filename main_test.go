package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetFizzbuzzHandler(t *testing.T) {
	tests := []struct {
		param    string
		expected string
	}{
		{"1", "1"},
		{"2", "2"},
		{"3", "fizz"},
		{"5", "buzz"},
		{"15", "fizzbuzz"},
		{"NaN", "invalid number"},
	}

	for _, tt := range tests {
		c, rec := makeContext(tt.param)

		if assert.NoError(t, getFizzbuzzHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, tt.expected, rec.Body.String())
		}
	}
}

func makeContext(param string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:number")
	c.SetParamNames("number")
	c.SetParamValues(param)
	return c, rec
}
