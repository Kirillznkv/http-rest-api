package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandlHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	res, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, res)
	assert.Equal(t, rec.Body.String(), "Hello")
}
