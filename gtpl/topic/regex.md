
---
***Regular Expressions***

According to the [documentation of the regexp package](https://golang.org/pkg/regexp/), Go's implementation is guaranteed to run in _linear time_; that is, it doesn't get slower as the size of the input gets bigger. Because of this, regular expressions in Go have two separate stages: `compilation` and `usage`.

**Compilation**

When compiling a `static` regex it is advisable to use `regexp.MustCompile()`, and move the compilation to `package` level. If the regex is invalid, the program will panic on startup, making it very obvious to the developer that there's a problem.

```go
var (
	someRegex = regexp.MustCompile(`someregex`)
)

func SomeFunc(s string) {
	result := someRegex.FindAllString(s, -1)
	//...
}
```

If the regex is not static, it still makes sense to compile it only when necessary. In this case `rexexp.Compile()` is better, as it will return an error for an invalid regex, rather than panicking.

```go
func SomeFunc(s, param string) error {
    regexStr := fmt.Sprintf("regexWithParam%s", param)
    someRegex, err := regexp.Compile(regexStr)
    if err != nil {
        return err
    }

    result := someRegex.FindAllString(s, -1)
    ...
}
```

---