
---
***Regular Expressions***

According to the [documentation of the regexp package](https://golang.org/pkg/regexp/) go's implementation is 
guaranteed to "run in time linear in the size of the input".

To accomplish this go's regexp package has one big difference to using regular expressions in many other 
programming languages: the separation of `compilation` and `usage`.

**Compilation**

When compiling a `static` regex it is advicable to use `regexp.MustCompile` and move the compilation to 
`package level`. Since `MustCompile` panics if the regex is invalid this has another advantage: A developer 
will get feedback on an invalid regex as soon as he starts his application. This should reduce 
the risk of a static but invalid regular expressions making it into production to a minimum.

```
var (
    someRegex = regexp.MustCompile(`someregex`)
)

func SomeFunc(s string) {
    result := someRegex.FindAllString(s, -1)
    ...
}
```

If the regex is not static it still makes sense to reduce the number of recompiling to a minimum. In this 
case `rexexp.Compile` is better suited as it does not panic on runtime but returns an error if the regex 
is invalid.

```
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