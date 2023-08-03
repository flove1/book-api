package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	_, err := HashPassword("password")
	assert.Nil(t, err, "must be nil")

	_, err = HashPassword(strings.Repeat("a", 73))
	assert.NotNil(t, err, "must not be nil")
}

func TestCheckPassword(t *testing.T) {
	password1 := "password"
	password2 := "pass"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)

	assert.Empty(t, CheckPassword(password1, hash), "must pass")
	assert.NotEmpty(t, CheckPassword(password2, hash), "must fail")
}
