package raindrops

import (
	"strconv"
	"strings"
)

func Convert(num int) string {
	result := strings.Builder{}
	if num%3 == 0 {
		result.WriteString("Pling")
	}
	if num%5 == 0 {
		result.WriteString("Plang")
	}
	if num%7 == 0 {
		result.WriteString("Plong")
	}
	if result.Len() == 0 {
		result.WriteString(strconv.Itoa(num))
	}
	return result.String()
}
