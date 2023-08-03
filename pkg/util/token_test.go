package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken()
	assert.NotNil(t, token, "must not be empty")
	assert.Nil(t, err, "must be nil")
}
