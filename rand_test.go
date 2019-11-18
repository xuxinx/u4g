package u4g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	s1 := RandomString(10)
	s2 := RandomString(10)
	s3 := RandomString(20)
	assert.Equal(t, 10, len(s1))
	assert.Equal(t, 20, len(s3))
	assert.NotEqual(t, s1, s2)
}

func TestRandomIntArray(t *testing.T) {
	s1 := RandomIntArray(10)
	s2 := RandomIntArray(10)
	s3 := RandomIntArray(20)

	assert.Equal(t, 10, len(s1))
	assert.Equal(t, 20, len(s3))
	assert.NotEqual(t, s1, s2)
}
