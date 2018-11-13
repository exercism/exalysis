- I noticed you're converting a string into a `[]rune` to then loop over it. Actually, using `range` with a string can return runes directly:
  ```go
  for i, r := range word {
    // r is a rune
  }
  ```
