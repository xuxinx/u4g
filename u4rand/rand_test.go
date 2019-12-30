package u4rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	s1 := RandString(10)
	s2 := RandString(10)
	s3 := RandString(20)
	assert.Equal(t, 10, len(s1))
	assert.Equal(t, 20, len(s3))
	assert.NotEqual(t, s1, s2)
}

func TestRandIntArray(t *testing.T) {
	s1 := RandIntArray(10)
	s2 := RandIntArray(10)
	s3 := RandIntArray(20)

	assert.Equal(t, 10, len(s1))
	assert.Equal(t, 20, len(s3))
	assert.NotEqual(t, s1, s2)
}
