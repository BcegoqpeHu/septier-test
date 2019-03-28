package handlers

import (
	"bytes"
	"compress/gzip"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddMany(t *testing.T) {
	e := echo.New()
	h := NewHandler()

	go h.Hub.Run()

	requestJson, _ := ioutil.ReadFile("add_many_handler_request_test.json")
	responseJson, _ := ioutil.ReadFile("add_many_handler_response_test.json")

	req := httptest.NewRequest(http.MethodPost, "/addmany", strings.NewReader(string(requestJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.AddManyHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(responseJson), rec.Body.String())
	}

	h.Hub.Stop()
}

//noinspection GoUnhandledErrorResult
func TestAddManyGzip(t *testing.T) {
	e := echo.New()
	h := NewHandler()

	e.Use(middleware.Gzip())
	go h.Hub.Run()

	e.POST("/addmany", h.AddManyHandler)

	requestJson, _ := ioutil.ReadFile("add_many_handler_request_test.json")

	req := httptest.NewRequest(http.MethodPost, "/addmany", strings.NewReader(string(requestJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAcceptEncoding, "gzip")

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	if cl := rec.Header().Get("Content-Length"); cl != "" {
		assert.Equal(t, cl, rec.Body.Len())
	}

	r, err := gzip.NewReader(rec.Body)
	if assert.NoError(t, err) {
		defer r.Close()
		responseJson, _ := ioutil.ReadFile("add_many_handler_response_test.json")
		if assert.NoError(t, err) {
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r)
			assert.Equal(t, responseJson, buf.Bytes())
		}
	}

	h.Hub.Stop()
}
