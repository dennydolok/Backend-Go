package helper

import (
	"WallE/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(1, 1, config.InitConfig().SECRET_KEY)
	assert.NotNil(t, token)
	assert.NoError(t, err)
}

func TestGetUserId(t *testing.T) {
	token := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
	id := GetUserId(token)
	assert.NotNil(t, id)
}

func TestGetClaim(t *testing.T) {
	token := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
	id := GetClaim(token)
	assert.NotNil(t, id)
}

func TestCheckAdmin(t *testing.T) {
	t.Run("admin", func(t *testing.T) {
		err := CheckAdmin(1)
		assert.NoError(t, err)
	})
	t.Run("not admin", func(t *testing.T) {
		err := CheckAdmin(2)
		assert.Error(t, err)
	})
}

func TestCheckCustomer(t *testing.T) {
	t.Run("customer", func(t *testing.T) {
		err := CheckCustomer(2)
		assert.NoError(t, err)
	})
	t.Run("not customer", func(t *testing.T) {
		err := CheckCustomer(1)
		assert.Error(t, err)
	})
}
