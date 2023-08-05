package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "who are you")
	} else {
		fmt.Fprint(w, "hello "+name)
	}
}

func TestQueryParam(t *testing.T) {
	name := ""
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name="+name, nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	if name != "" {
		assert.Equal(t, "hello "+name, string(body))
	} else {
		assert.Equal(t, "who are you", string(body))
	}
}

func MultipleQueryParam(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")
	fmt.Println("Hello World")
	fmt.Fprintf(w, "%s %s", first_name, last_name)
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?first_name=Eko&last_name=Khannedy", nil)
	recorder := httptest.NewRecorder()
	// Hello World

	MultipleQueryParam(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	// Hello World
	assert.Equal(t, "Eko Khannedy", string(body))
}
