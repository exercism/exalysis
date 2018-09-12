package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/tehsphinx/exalysis/suggestion"
)

func main() {
	sugg, approval := suggestion.GetSuggestions()

	fmt.Println(sugg)
	fmt.Print(approval)
	clipboard.WriteAll(sugg)
}
