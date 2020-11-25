package controller

import (
	"battery-analysis-platform/app/web/middleware"
	"battery-analysis-platform/pkg/jd"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getGinEngine() *gin.Engine {
	gin.SetMode("release")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Session("test"))
	Register(r)
	return r
}

func request(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func jsonResponseEquel(t *testing.T, w *httptest.ResponseRecorder, code int, msg string, data interface{}) {
	ast := assert.New(t)

	var response map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	decoder.UseNumber()
	err := decoder.Decode(&response)
	ast.Nil(err)

	tmp, err := response["code"].(json.Number).Int64()
	ast.Nil(err)

	assert.Equal(t, int(tmp), code)
	assert.Equal(t, response["msg"], msg)
	assert.Equal(t, response["data"], data)
}

func TestLogin(t *testing.T) {
	r := getGinEngine()
	w := request(r, "GET", "/login", nil)
	assert.Equal(t, http.StatusOK, w.Code)

	jsonResponseEquel(t, w, jd.ERROR, "", nil)
}

func TestLogout(t *testing.T) {
	r := getGinEngine()
	w := request(r, "POST", "/logout", nil)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
