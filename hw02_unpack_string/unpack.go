package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type Symbol struct {
	Value    rune
	IsLetter bool
	IsDigit  bool
}

func Unpack(input string) (string, error) {
	var sb strings.Builder
	var prev Symbol
	for i, val := range input {
		current := Symbol{
			Value:    val,
			IsLetter: unicode.IsLetter(val),
			IsDigit:  unicode.IsDigit(val),
		}

		if current.IsDigit {
			if prev.IsLetter {
				for i := 0; i < int(current.Value-'0'); i++ {
					sb.WriteRune(prev.Value)
				}
			} else {
				return "", ErrInvalidString
			}
		} else if prev.IsLetter {
			sb.WriteRune(prev.Value)
		}
		if i == len(input)-1 {
			sb.WriteRune(current.Value)
		}

		prev = current
	}

	return sb.String(), nil
}
