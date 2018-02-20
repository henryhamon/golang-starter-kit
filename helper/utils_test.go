package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	assert.True(t, StringInSlice("a", []string{"a", "b"}))
	assert.False(t, StringInSlice("x", []string{"a", "b"}))
}
