package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

type Adapter func(http.Handler) http.Handler

func Notify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("before")
			defer logger.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}

// NOTE: adapters/middleware are run in reverse order!
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	// example 1 of middleware
	// http.Handle("/count", Notify(logger)(new(countHandler)))
	// example 2 of middleware (Recommended)
	http.Handle("/count", Adapt(new(countHandler), Notify(logger)))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
