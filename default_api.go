package main

import (
	"net/http"
	"strconv"
	"strings"
)

func CheckNumber(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(FizzBuzz(num)))
}

func FizzBuzz(num int) string {
	fizz := num % 3
	buzz := num % 5

	res := ""
	if fizz == 0 {
		res += "fizz"
	}
	if buzz == 0 {
		res += "buzz"
	}

	return res
}
