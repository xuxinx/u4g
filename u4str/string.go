package u4str

import (
	"strings"
	"unicode/utf8"
)

// Utf8Index returns unicode index
// strings.Index("一二三", "二") returns 3, Utf8Index("一二三", "二") returns 1
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}

	return utf8.RuneCountInString(str[:index])
}
