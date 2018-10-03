- You are currently converting a `rune` to a `byte`. The DNA strands in this exercise do not contain 
unicode characters so this is fine. Just remember: if you have to deal with special characters this 
conversion is lossy! A rune is `int32` a byte only `uint8`. A rune can consist of multiple bytes.

    To get a deeper understanding you could 
read the post on [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings).