package util

import (
	"fmt"
	"strings"
)

func QuoteSpaces(datum interface{}) string {
	str := fmt.Sprintf("%+v", datum)

	if len(strings.Split(str, " ")) > 1 {
		return `"` + str + `"`
	}

	return str
}

func StringSliceContains(haystack []string, needles ...string) bool {
	for _, val := range haystack {
		for _, needle := range needles {
			if val == needle {
				return true
			}
		}
	}

	return false
}
