package apiserver

import (
	"github.com/Kirillznkv/new_rest_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/users", nil)
	assert.NoError(t, err)

	s := NewServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
