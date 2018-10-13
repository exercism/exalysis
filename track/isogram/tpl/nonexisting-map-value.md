- Did you know that a map in go never panics if you ask for a nonexistent key? Instead it returns the 
[zero value](https://golang.org/ref/spec#The_zero_value) every type in go has.
  If you have a `map[sometype]bool` it will return `false` if the entry does not exist. You can turn that to
your advantage in this exercise and query the `map` without using the `value, ok` syntax: `if somemap[somekey] {...}`.