package api

import (
	"io"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

// NewMockAPI
func NewMockAPI(method, target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}
