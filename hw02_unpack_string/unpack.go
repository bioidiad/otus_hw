package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

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

	s = Reverse(s)
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

	return Reverse(b.String()), nil
}
