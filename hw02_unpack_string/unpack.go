package hw02unpackstring

import (
	"errors"
	"golang.org/x/example/stringutil"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}

	var (
		err error
		b   strings.Builder
	)

	s = stringutil.Reverse(s)
	tmp := rune(s[0])
	counter := 1

	for _, r := range s {
		if unicode.IsDigit(r) {
			if unicode.IsDigit(tmp) {
				return "", ErrInvalidString
			}
			counter, err = strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
		} else {
			b.WriteString(strings.Repeat(string(r), counter))
			counter = 1
		}
		tmp = r
	}

	return stringutil.Reverse(b.String()), nil
}
