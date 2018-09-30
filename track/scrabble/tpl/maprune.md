- you could use a `map[rune]int` instead of `%s` for direct lookup of the score value for a `rune`: 
no type conversion needed. A `rune` is created with single quotes e.g. 'A' just like a `byte`.