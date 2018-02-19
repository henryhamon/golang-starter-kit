package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultParameters(t *testing.T) {
	params, err := NewParameter()
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
