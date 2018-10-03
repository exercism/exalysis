- As you realized ranging over a `string` creates `runes` however taking an index from a string returns 
a `byte`. This is also described in more detail in the 
post [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings).

  In this exercise you could just convert the `byte` to a `rune` and then compare saving you 
one extra type conversion.