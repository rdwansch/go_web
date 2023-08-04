package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprintln(w, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	// sama seperti /images/[...] di JS
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Method)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
