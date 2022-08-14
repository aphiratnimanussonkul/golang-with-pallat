package todo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

type fakeStore struct{}

func (s *fakeStore) NewTask(task string) error {
	return nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/todos", NewHandler(&fakeStore{}).NewTask)
	return r
}

func TestNewTaskHandler(t *testing.T) {
	router := setupRouter()

	payload := bytes.NewBuffer([]byte(`{"title":"task"}`))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/todos", payload)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"insert success"}`, w.Body.String())
}
