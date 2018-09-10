package main

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/tehsphinx/exalysis/suggestion"
	"golang.org/x/tools/go/loader"
)

func main() {
	prog := loadProg()

	local := prog.Package(".")
	sugg := suggestion.GetSuggestions(local.Pkg.Name(), local, prog)

	fmt.Println(sugg)
	clipboard.WriteAll(sugg)
}

func loadProg() *loader.Program {
	var conf loader.Config
	conf.Import(".")
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}
	return prog
}
