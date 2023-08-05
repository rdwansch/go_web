package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Programmer Zaman Now")

	header := r.Header.Get("Authorization")

	fmt.Fprint(w, header)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	request.Header.Add("Authorization", "Token cnv38hfh02a8hf3")

	HeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	poweredBy := response.Header.Get("X-Powered-By")

	fmt.Println("Request header " + string(body))
	fmt.Println("Response header " + poweredBy)
}
