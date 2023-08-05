package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHTTP(t *testing.T) {
	// server:= http.Server{
	// 	Addr: "localhost:8080",
	// 	Handler: HelloHandler,
	// }

	// dengan ini kita tidak perlu membuat server manual
	// melainkan hanya simulasi
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil) // Request
	recorder := httptest.NewRecorder()                                           // Response

	// eksekusi server nya
	HelloHandler(recorder, request)

	// mengambil response nya
	response := recorder.Result() // *http.Response

	// response body biasanya berbentuk blob, sehingga kita perlu memparsing nya
	// jika di JS adalah response.json()
	body, _ := io.ReadAll(response.Body)
	//helow world
	assert.Equal(t, "Hello World", string(body))
}
