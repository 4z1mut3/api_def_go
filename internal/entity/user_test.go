package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("paulo andre", "andre@andre.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "paulo andre", user.Name)
	assert.Equal(t, "andre@andre.com", user.Email)

}
func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("paulo andre", "andre@andre.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.validatePassword("123456"))
	assert.False(t, user.validatePassword("123123"))
	assert.NotEqual(t, "123456", user.Password)
}
