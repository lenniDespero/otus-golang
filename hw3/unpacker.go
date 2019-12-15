package hw3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Unpack(s string) (string, error) {
	var unpack []string
	var flagEscape bool
	for i, value := range strings.Split(s, "") {
		switch {
		case value == `\` && !flagEscape:
			flagEscape = true
		case flagEscape:
			unpack, flagEscape = append(unpack, value), false
		case isNumber(value) && len(unpack) > 0:
			number, _ := strconv.Atoi(value)
			if number == 0 {
				unpack = unpack[:len(unpack)-1]
			} else {
				w := unpack[len(unpack)-1]
				for a := 0; a < number-1; a++ {
					unpack = append(unpack, w)
				}
			}
		case !isNumber(value):
			unpack = append(unpack, value)
		default:
			return "", errors.New(fmt.Sprintf("Something strange with unpacking string '%s' on character - %s, position - %d.", s, value, i))
		}
	}
	return strings.Join(unpack, ""), nil
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

