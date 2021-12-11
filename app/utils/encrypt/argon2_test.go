package encrypt

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
	password := "test"

	hash, err := CreateHash(password, DefaultParams)
	assert.NotNil(t, hash)
	assert.NoError(t, err)
}

func TestComparePasswordAndHash(t *testing.T) {
	password := "test"
	argon, errArgon := CreateHash(password, DefaultParams)
	assert.NoError(t, errArgon)
	check, errCheck := ComparePasswordAndHash("test", argon)
	assert.NoError(t, errCheck)

	assert.Equal(t, true, check)
}

func TestCheckHash(t *testing.T) {
	password := "test"
	argon, errArgon := CreateHash(password, DefaultParams)
	assert.NoError(t, errArgon)
	check, errCheck := ComparePasswordAndHash("test", argon)
	assert.NoError(t, errCheck)

	assert.Equal(t, true, check)
}

func TestGenerateRandomBytes(t *testing.T) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(27))
	assert.NoError(t, err)

	assert.NotNil(t, nBig)
}
