module github.com/exercism/exalysis

go 1.14

require (
	github.com/atotto/clipboard v0.1.1 // indirect
	github.com/exercism/go-analyzer v0.2.1
	github.com/kevinburke/go-bindata v3.13.0+incompatible
	github.com/logrusorgru/aurora v0.0.0-20181002194514-a7b3b318ed4e
	github.com/pmezard/go-difflib v1.0.0
	github.com/stretchr/testify v1.3.0
	github.com/tehsphinx/astrav v0.4.1
	github.com/tehsphinx/clipboard v0.1.3
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/tools v0.0.0-20200925191224-5d1fdd8fa346 // indirect
)

//replace golang.org/x/lint => github.com/golang/lint v0.0.0-20180702182130-06c8688daad7

// replace github.com/exercism/go-analyzer => ../go-analyzer

// replace github.com/tehsphinx/astrav => ../astrav
