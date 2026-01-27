package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	rs := []rune(s)
	rsLen := len(rs)

	if rsLen == 0 {
		return "", nil
	} else if unicode.IsDigit(rs[0]) {
		return "", ErrInvalidString
	}

	var b strings.Builder
	for i := 0; i < rsLen; i++ {
		switch {
		// string should not include numbers
		case unicode.IsDigit(rs[i]) && unicode.IsDigit(rs[i-1]):
			return "", ErrInvalidString
		case i+1 < rsLen && unicode.IsDigit(rs[i+1]):
			count, err := strconv.Atoi(string(rs[i+1]))
			if err != nil {
				return "", err
			}
			b.WriteString(strings.Repeat(string(rs[i]), count))

			// digits should be skipped in result string
			i++
		default:
			b.WriteString(string(rs[i]))
		}
	}
	return b.String(), nil
}
