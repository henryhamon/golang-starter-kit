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
}
