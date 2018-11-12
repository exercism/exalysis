- As you know, using `range` over a string gives you a rune, but indexing into a string returns a byte value. The difference is described in more detail here: [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings).

  In this exercise you could just convert the byte to a rune, saving you one extra type conversion.
