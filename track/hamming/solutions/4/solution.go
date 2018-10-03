/*
Package hamming is used for calculating difference between DNA
It has a function Distance(string, string) (int,error), that accept
2 DNA strings and return difference in integer value, or -1 and error message
 */
package hamming

import (
	"errors"
	"strings"
)

//Distance is used for calculating difference between 2 DNA, represented in strings.
func Distance(a, b string) (int, error) {
	var answer int
	first := strings.Split(a,"")
	second := strings.Split(b,"")
	if len(first) == len(second) {
		for i,v := range(first) {
			if v != second[i] {
				answer += 1
			}
		}
	}else{
		return -1, errors.New("Error, DNA strands have different length!")
	}
	return answer, nil
}
