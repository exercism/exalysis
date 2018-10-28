// Package raindrops ... .
package raindrops

import (
	"strconv"
	"strings"
)

var m = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert ... .
func Convert(num int) string {
	var out strings.Builder

	for k, v := range m {
		if num%k == 0 {
			out.WriteString(v)
		}
	}

	if out.Len() == 0 {
		out.WriteString(strconv.Itoa(num))
	}

	return out.String()
}
