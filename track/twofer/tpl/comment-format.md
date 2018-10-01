---
Go has [great guidelines](https://golang.org/doc/effective_go.html)
about how to write comments.

The package comment should start with the word _Package_ followed by the package name.

```
// Package cook provides handy conversion methods for units typically used in recipes.
package cook
```

For exported functions, methods, constants, and package variables, the comment should start with the name.

```
// TbsToMl converts tablespoons to milliliters.
func TbsToMl(tbs int) int {
	// ...
}
```

The [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#comment-sentences) wiki further specifies that doc comments should be full sentences, ending with a period.

In addition to optimizing for readability, following the recommendations improves how your code
looks in [Go's documentation tools](http://whipperstacker.com/2015/09/30/go-documentation-godoc-godoc-godoc-org-and-go-doc/)
such as [godoc.org](http://godoc.org), `godoc`, and `go doc`.

Take a moment to read the section on doc comments in the official style guide,
[Effective Go](https://golang.org/doc/effective_go.html) to see the reasoning
behind these choices.
---