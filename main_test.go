package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	e = addRoutes(e)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e.ServeHTTP(rec, req)

	expected := "Hello, World!"
	actual := rec.Body.String()
	assert.Equal(expected, actual)
}
