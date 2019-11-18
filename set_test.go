package u4g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	s := NewSet()
	assert.Equal(t, 0, s.Len())

	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, true, s.Has(1))

	s.Add(2)
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, true, s.Has(2))

	s.Add(1, 2, 3, 3)
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, true, s.Has(3))
}

func TestSet_Delete(t *testing.T) {
	s := NewSet()

	s.Add(1, 2, 3, 4)
	assert.Equal(t, 4, s.Len())

	s.Delete(1)
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, false, s.Has(1))

	s.Delete(1, 3)
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, false, s.Has(3))
}

func TestSet_Clear(t *testing.T) {
	s := NewSet()

	s.Add(1, 2, 3, 4)
	assert.Equal(t, 4, s.Len())

	s.Clear()
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, false, s.Has(1))
}

func TestSet_Items(t *testing.T) {
	s := NewSet()
	assert.Equal(t, 0, len(s.Items()))

	s.Add(1, 2, 3, 4)
	assert.ElementsMatch(t, []interface{}{1, 2, 3, 4}, s.Items())
}

func TestSet_Has(t *testing.T) {
	s := NewSet()

	has := s.Has(1)
	assert.Equal(t, false, has)

	s.Add(1)
	assert.Equal(t, true, s.Has(1))
}
