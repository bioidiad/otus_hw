package hw02unpackstring

import (
	"errors"
	"fmt"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sentence string) (string, error) {
	for _, char := range sentence {
		fmt.Println(char)
	}
	return "", nil
}
