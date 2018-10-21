/*
Package scrabble used to calculate score for word
*/
package scrabble

import (
	"regexp"
	"strings"
)

//Score used to calculate score for input word
func Score(input string) int {
	input = strings.ToLower(input)
	oneReg := regexp.MustCompile("[aeioulnrst]")
	twoReg := regexp.MustCompile("[dg]")
	threeReg := regexp.MustCompile("[bcmp]")
	thourReg := regexp.MustCompile("[fhvwy]")
	fiveReg := regexp.MustCompile("[k]")
	eightReg := regexp.MustCompile("[jx]")
	tenReg := regexp.MustCompile("[qz]")
	return len(oneReg.FindAllString(input, -1)) +
		len(twoReg.FindAllString(input, -1))*2 +
		len(threeReg.FindAllString(input, -1))*3 +
		len(thourReg.FindAllString(input, -1))*4 +
		len(fiveReg.FindAllString(input, -1))*5 +
		len(eightReg.FindAllString(input, -1))*8 +
		len(tenReg.FindAllString(input, -1))*10
}
