package shortener

import (
	"crypto/rand"
)

var (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Generate(size int) (string, error) {
	chars := []rune(alphabet)
	n := len(chars)

	mask := 1
	for mask < n {
		mask <<= 1
	}
	mask--

	result := make([]rune, size)
	for i := 0; i < size; {
		b := make([]byte, 1)
		rand.Read(b)

		idx := int(b[0] & byte(mask))
		if idx < n {
			result[i] = chars[idx]
			i++
		}
	}

	return string(result), nil
}
