package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(Init("./log"))
	assert.Nil(os.RemoveAll("./log"))
}
