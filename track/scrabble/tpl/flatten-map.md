- The way the map is created you currently need two loops. You could have a loop at the benchmarks,
  rearrange the map for direct lookup and check the speed again. Best use a `map[rune]int`.