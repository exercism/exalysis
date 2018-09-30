package extypes

import (
	"github.com/tehsphinx/astrav"
)

//SuggestionFunc defines a function checking a solution for a specific problem
type SuggestionFunc func(pkg *astrav.Package, r *Response)
