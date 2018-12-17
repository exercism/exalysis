package raindrops

import (
	"fmt"
	"sort"
)

var factorMap = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converting the input nr to rain drops
func Convert(input int) string {
	factors := extractFactors(input)
	var msg string
	for _, factor := range factors {
		if m, ok := factorMap[factor]; ok {
			msg += m
		}
	}
	if msg == "" {
		msg = fmt.Sprintf("%d", input)
	}
	return msg
}

func extractFactors(number int) []int {
	if number == 1 {
		return []int{1}
	}
	var factors = []int{}
	var factorsMap = map[int]bool{}
	var factor = 1
	var factorExist = false
	for !factorExist {
		if number%factor == 0 {
			factors = append(factors, factor)
			factorsMap[factor] = true
			var result = number / factor
			if result != factor {
				factors = append(factors, result)
				factorsMap[number/factor] = true
			}
		}
		factor++
		_, factorExist = factorsMap[factor]
	}
	sort.Ints(factors)
	return factors
}
