package router

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	Route(e)
	assert.NotNil(e)
}
