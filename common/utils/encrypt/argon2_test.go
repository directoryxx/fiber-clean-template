package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHashDefaultHash(t *testing.T) {
	password := "test"

	hash, err := CreateHash(password, DefaultParams)
	assert.NotNil(t, hash)
	assert.NoError(t, err)
}
