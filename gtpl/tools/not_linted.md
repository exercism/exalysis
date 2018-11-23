- There's a great tool called [`golint`](https://github.com/golang/lint) which will examine your code for common problems and style issues. Try running `golint` on your code; it will make some useful suggestions for you.

    Try to get into the habit of always checking your code with `golint` (or configuring your editor to do it for you). When you're writing Go software for production, many build pipelines will fail commits automatically if they don't pass `golint` (and `gofmt`). Avoid this by linting code yourself!
