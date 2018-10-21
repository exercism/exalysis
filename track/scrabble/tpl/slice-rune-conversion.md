- You are converting a string into a `[]rune` to then loop over it. This is not necessary as looping over a 
string will automatically create runes: 
  ```
  for i, r := range word {
    // r is of type rune
  }
  ```