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
	} else if err := validateString(rs); err != nil {
		return "", err
	}

	var b strings.Builder
	for i := 0; i < rsLen; i++ {
		if i+1 < rsLen && unicode.IsDigit(rs[i+1]) {
			count, err := strconv.Atoi(string(rs[i+1]))
			if err != nil {
				panic(err)
			}
			b.WriteString(strings.Repeat(string(rs[i]), count))

			// digits should be skipped in result string
			i++
		} else {
			b.WriteString(string(rs[i]))
		}
	}
	return b.String(), nil
}

func validateString(rs []rune) error {
	rsLen := len(rs)

	// string should not begin with digit
	if unicode.IsDigit(rs[0]) {
		return ErrInvalidString
	}

	// string should not include numbers
	for i := 1; i < rsLen; i++ {
		if unicode.IsDigit(rs[i]) && unicode.IsDigit(rs[i-1]) {
			return ErrInvalidString
		}
	}
	return nil
}
