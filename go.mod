module github.com/exercism/exalysis

require (
	github.com/atotto/clipboard v0.1.1 // indirect
	github.com/kevinburke/go-bindata v3.13.0+incompatible
	github.com/logrusorgru/aurora v0.0.0-20181002194514-a7b3b318ed4e
	github.com/pmezard/go-difflib v1.0.0
	github.com/stretchr/testify v1.2.2
	github.com/tehsphinx/astrav v0.1.0
	github.com/tehsphinx/clipboard v0.1.3
	golang.org/x/lint v0.0.0-20181212231659-93c0bb5c8393
	golang.org/x/tools v0.0.0-20181214171254-3c39ce7b6105 // indirect
)

exclude (
	github.com/tehsphinx/astrav v0.2.0
	github.com/tehsphinx/astrav v0.2.1
	github.com/tehsphinx/astrav v0.2.2
	github.com/tehsphinx/astrav v0.2.3
)

replace golang.org/x/lint => github.com/golang/lint v0.0.0-20180702182130-06c8688daad7
