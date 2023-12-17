package middleware

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	Use(e)
	assert.NotNil(e)
}
