package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	// menambahkan body
	// jika di node js menggunakan JSON
	requestBody := strings.NewReader("first_name=Ridhwan&last_name=Siddiq")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	record := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	FormPost(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
