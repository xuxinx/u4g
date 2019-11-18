package u4g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtf8Index(t *testing.T) {
	cases := []struct {
		name   string
		str    string
		substr string
		expect int
	}{
		{
			name:   "404",
			str:    "一二三",
			substr: "四",
			expect: -1,
		},
		{
			name:   "first",
			str:    "一二三",
			substr: "一",
			expect: 0,
		},
		{
			name:   "not first",
			str:    "一二三",
			substr: "二",
			expect: 1,
		},
		{
			name:   "letters and chinese",
			str:    "q一二r",
			substr: "二",
			expect: 2,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, Utf8Index(c.str, c.substr))
	}
}
