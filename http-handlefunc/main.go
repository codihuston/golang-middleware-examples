package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped

}

func LogMiddleware1(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.SetOutput(os.Stdout) // logs go to Stderr by default
		log.Println("1", r.Method, r.URL)
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler

	})
}

func LogMiddleware2(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.SetOutput(os.Stdout) // logs go to Stderr by default
		log.Println("2", r.Method, r.URL)
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler

	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Index!")
}

func main() {
	// example 1
	http.HandleFunc("/", LogMiddleware1(LogMiddleware2(IndexHandler)))
	// example 2
	http.HandleFunc("/2", MultipleMiddleware(IndexHandler, LogMiddleware1, LogMiddleware2))

	log.Fatal(http.ListenAndServe(":8082", nil))
}
