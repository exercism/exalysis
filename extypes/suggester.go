package extypes

import (
	"github.com/tehsphinx/astrav"
)

//SuggesterCreator defines a creator function for a Suggestor
type SuggesterCreator func(pkg *astrav.Package) Suggester

//Suggester defines an interface to be implemented by an exercise suggestion creator.
type Suggester interface {
	Suggest(r *Response)
}
