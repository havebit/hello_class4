package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos", NewHandler(&fakeStore{}).NewTask)
	return r
}

type fakeStore struct{}

func (s *fakeStore) NewTask(task string) error {
	return nil
}

func TestNewTaskHandler(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	router.ServeHTTP(w, req)

	if http.StatusOK != w.Code {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", w.Code, http.StatusOK)
	}
}
