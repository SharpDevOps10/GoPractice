package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HandleHello(t *testing.T) {
	s := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	s.handleHello().ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Body.String(), "Hello")
}
