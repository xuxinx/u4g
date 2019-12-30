package u4rand

import (
	"math/rand"
	"time"
)

var (
	letters    []rune
	lettersLen int
)

func init() {
	for i := 'A'; i <= 'Z'; i++ {
		letters = append(letters, i)
	}
	for i := 'a'; i <= 'z'; i++ {
		letters = append(letters, i)
	}
	lettersLen = len(letters)

	rand.Seed(time.Now().UnixNano())
}

func RandString(length int) string {
	s := make([]rune, length)
	for i := 0; i < length; i++ {
		s[i] = letters[rand.Intn(lettersLen)]
	}

	return string(s)
}

func RandIntArray(length int) []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = rand.Int()
	}

	return s
}
