package main

import (
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/tehsphinx/exalysis/suggestion"
)

var (
	minConfidence = flag.Float64("min_confidence", 0.8, "golint: minimum confidence of a problem to print it")
)

func main() {
	flag.Parse()
	suggestion.LintMinConfidence = *minConfidence

	sugg, approval := suggestion.GetSuggestions()

	fmt.Println(sugg)
	fmt.Print(approval)
	clipboard.WriteAll(sugg)
}
