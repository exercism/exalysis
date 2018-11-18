
---
Your solution with multiple loops is readable and totally fine. If you want to avoid looping over the string multiple times you can find a step by step instruction here:

- First thing we need to do is incorporate the non-digit handling into the main loop:

  Whitespace: whitespace is allowed so we can `continue` on whitespace in the loop.

  Other non-digit: we `return false` if we encounter them in the loop. You can do that with `unicode.IsDigit` if you are working with runes.

- With the WhiteSpace change we introduced a problem: The index `i` of the loop is not counting "correctly" any more. We can create a new, independant `counter` and increment manually in the loop (after the `continue` and `return` cases). Don't forget to check after the loop if `counter > 1`.

- Another Problem we introduced with the WhiteSpace change is that we cannot calculate the `length` before the loop any more. But we can now iterate backwards over the string since we have our independant counter counting up: `for i := len(str) - 1; 0 <= i; i-- {...}`. Our digit variable is then `r := rune(str[i])`. We use `counter` to check for `modulo 2` now.

At this point things should work again and be many times faster.

- One last thing: We can replace `!unicode.IsDigit(r)` with `r < '0' || '9' < r` and we get another doubling of speed.

  Note: with this step we can drop the conversion to `rune`. It works the same with `byte`.

Now we should be in the area of `20-100 ns/op` depending on the hardware the benchmarks are run on.

---
