package main

import (
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestCheckNumber(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"/1", ""},
		{"/3", "fizz"},
		{"/5", "buzz"},
		{"/15", "fizzbuzz"},
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", c.in, nil)
		if err != nil {
			t.Fatal(err)
		}

		resrec := httptest.NewRecorder()
		handler := http.HandlerFunc(CheckNumber)
		handler.ServeHTTP(resrec, req)

		expectedStatus := http.StatusOK
		if status := resrec.Code; status != expectedStatus {
			t.Errorf("CheckNumber: Bad status code: expected: %d, got: %d", expectedStatus, status)
		}

		if resrec.Body.String() != c.expected {
			t.Errorf("CheckNumber: Unexpected body: expected: %q, got: %q", c.expected, resrec.Body.String())
		}
	}

	errorcases := []struct {
		in       string
		expected int
	}{
		{"/a", http.StatusBadRequest},
		{"/%21", http.StatusBadRequest},
	}

	for _, ec := range errorcases {
		req, err := http.NewRequest("GET", ec.in, nil)
		if err != nil {
			t.Fatal(err)
		}

		resrec := httptest.NewRecorder()
		handler := http.HandlerFunc(CheckNumber)
		handler.ServeHTTP(resrec, req)

		if status := resrec.Code; status != ec.expected {
			t.Errorf("CheckNumber: Unexpected status code: expected: %d, got %d", ec.expected, status)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in       int
		expected string
	}{
		{1, ""},
		{3, "fizz"},
		{5, "buzz"},
		{15, "fizzbuzz"},
	}

	for _, c := range cases {
		got := FizzBuzz(c.in)
		if got != c.expected {
			t.Errorf("FizzBuzz(%d): expected: %q, got: %q", c.in, c.expected, got)
		}
	}
}
