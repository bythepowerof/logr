package utils

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
