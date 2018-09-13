package defs

import (
	"golang.org/x/tools/go/loader"
	"honnef.co/go/tools/ssa"
)

//SuggesterCreator defines a creator function for a Suggestor
type SuggesterCreator func(program *loader.Program, pkg *ssa.Package) Suggester

//Suggester defines an interface to be implemented by an exercise suggestion creator.
type Suggester interface {
	Suggest() []string
}
