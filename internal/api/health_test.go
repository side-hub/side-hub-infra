package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	assert := assert.New(t)

	ctx, rec := NewMockAPI(http.MethodGet, "/health", nil)
	assert.Nil(GetHealth(ctx))
	assert.Equal(http.StatusOK, rec.Code)
}
