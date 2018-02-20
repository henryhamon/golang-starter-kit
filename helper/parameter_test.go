package helper

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDefaultParameters(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/", nil)

	params, err := NewParameter(c)
	assert.NoError(t, err)

	assert.Equal(t, 1, params.Page)
	assert.Equal(t, 25, params.Limit)
	assert.Equal(t, "asc", params.Order)
}

func TestValidateParams(t *testing.T) {
	number, err := validate("")
	assert.NoError(t, err)
	assert.Equal(t, -1, number)

	number, err = validate("2")
	assert.NoError(t, err)
	assert.Equal(t, 2, number)

	number, err = validate("abcd")
	assert.Error(t, err)
	assert.Equal(t, 0, number)
}

func TestParameters(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/?limit=10&order=desc&page=4", nil)

	params, err := NewParameter(c)
	assert.NoError(t, err)
	assert.Equal(t, 10, params.Limit)
	assert.Equal(t, "desc", params.Order)
	assert.Equal(t, 4, params.Page)
}

func TestParams(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/?page=5&foo=bar", nil)

	params, err := NewParameter(c)
	assert.NoError(t, err)
	assert.Equal(t, "bar", params.Params["foo"])
}
