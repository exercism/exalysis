package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/tehsphinx/exalysis"
)

var (
	minConfidence = flag.Float64("min_confidence", 0.8, "golint: minimum confidence of a problem to print it")
)

func main() {
	flag.Parse()
	exalysis.LintMinConfidence = *minConfidence

	sugg, approval := exalysis.GetSuggestions()

	fmt.Println("\n" + sugg)
	fmt.Print("\n\n" + approval)
	if err := clipboard.WriteAll(sugg); err != nil {
		log.Println(err)
	}
}
