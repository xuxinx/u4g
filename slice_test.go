package u4g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceChunks(t *testing.T) {
	s := []interface{}{"1", "2", "3", "4", "5", "6"}

	cases := []struct {
		size      int
		expectRes [][]interface{}
	}{
		{
			size:      1,
			expectRes: [][]interface{}{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}, {"6"}},
		},
		{
			size:      2,
			expectRes: [][]interface{}{{"1", "2"}, {"3", "4"}, {"5", "6"}},
		},
		{
			size:      3,
			expectRes: [][]interface{}{{"1", "2", "3"}, {"4", "5", "6"}},
		},
		{
			size:      4,
			expectRes: [][]interface{}{{"1", "2", "3", "4"}, {"5", "6"}},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectRes, SliceChunks(s, c.size))
	}
}
