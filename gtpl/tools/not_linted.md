- There's a great tool called [`golint`](https://github.com/golang/lint) which will examine your code for common problems and style issues. Try running `golint` on your code, e.g.: `golint two_fer.go`; it will make some useful suggestions for you.

    It's a good idea to always check your code with `golint` (or configure your editor to do it for you). When you're writing Go software for production, many build pipelines will automatically fail if the code doesn't pass `golint` (and `gofmt`). You can avoid this by linting code yourself prior to submitting it.
