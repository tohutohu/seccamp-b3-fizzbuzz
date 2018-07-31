package main

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestGetFizzbuzzHandler(t *testing.T) {
	tests := []struct {
		param        string
		expectedText string
		expectedCode int
	}{
		{"1", "1", 200},
		{"2", "2", 200},
		{"3", "fizz", 200},
		{"5", "buzz", 200},
		{"15", "fizzbuzz", 200},
		{"NaN", "invalid number", 400},
	}

	for i, tt := range tests {
		c, rec := makeContext(tt.param)

		if err := getFizzbuzzHandler(c); err != nil {
			t.Fatal(err)
		}

		if got, exp := rec.Code, tt.expectedCode; got != exp {
			t.Fatalf("case %d> wrong status code: got %d, expected %d", i+1, got, exp)
		}

		if got, exp := rec.Body.String(), tt.expectedText; got != exp {
			t.Fatalf("case %d> wrong response text: got %s, expected %s", i+1, got, exp)
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
