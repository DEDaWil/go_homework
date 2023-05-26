package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type Symbol struct {
	value    rune
	isLetter bool
	isDigit  bool
	isSlash  bool
}

func Unpack(input string) (string, error) {
	var sb strings.Builder
	var prev Symbol
	for i, val := range input {
		current := Symbol{
			value:    val,
			isLetter: unicode.IsLetter(val),
			isDigit:  unicode.IsDigit(val),
			isSlash:  val == []rune(`\`)[0],
		}

		if prev.isSlash {
			current.isLetter = true
			current.isDigit = false
			current.isSlash = false
		}

		if current.isDigit {
			if prev.isLetter {
				sb.WriteString(strings.Repeat(string(prev.value), int(current.value-'0')))
			} else {
				return "", ErrInvalidString
			}
		} else if prev.isLetter {
			sb.WriteRune(prev.value)
		}

		if current.isLetter && i == len(input)-1 {
			sb.WriteRune(current.value)
		}

		prev = current
	}

	return sb.String(), nil
}
