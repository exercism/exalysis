package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/tehsphinx/exalysis/suggestion"
)

func main() {
	sugg := suggestion.GetSuggestions()

	fmt.Println(sugg)
	clipboard.WriteAll(sugg)
}
