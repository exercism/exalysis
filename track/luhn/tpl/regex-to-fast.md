
---
Your solution with `regexp` is very readable and totally fine. Regular expressions are very powerful, but usually slow. If things can be done manually with reasonable effort the extra work is often worth it. Here is a step-by-step guide to a faster solution:

1. First, eliminate the regular expression by moving the logic into the loop:

    - Whitespace is allowed, so we can `continue` if we encounter whitespace in the loop.

    - Other non-digits: we `return false` if we encounter them in the loop. You can do that with `unicode.IsDigit` if you are working with runes.

1. With this whitespace change, we introduced a problem: the index `i` of the loop is not counting "correctly" any more. We can create a new variable `counter` to keep track of how many digits we've seen, and increment it manually in the loop (after the `continue` and `return` cases). Don't forget to check after the loop if `counter > 1` to see if we got more than one digit.

1. Another problem we introduced with the whitespace change is that we cannot calculate the `length` before the loop anymore. But we can now iterate backwards over the string since we have an independent digit counter: `for i := len(str) - 1; 0 <= i; i-- {...}`.

    The digit rune is then `r := rune(str[i])`. You can test the value of `counter` to see whether or not the current digit should be doubled.

At this point things should work again and be many times faster!

- One last thing: We can replace `!unicode.IsDigit(r)` with `r < '0' || r > '9'` and we get another increase in speed. (However, calling `unicode.IsDigit` makes it slightly clearer what's going on. Usually, readability is more important than speed.)

  Note: with this step we can drop the conversion to `rune`. It works the same with `byte`.

Now we should be in the area of `20-100 ns/op` depending on the hardware the benchmarks are run on.

---
