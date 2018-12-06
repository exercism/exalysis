- I spotted that you're using `strings.Join` to build the string. That works, but you might like to look at the `fmt.Sprintf` function instead; it's a very powerful way of building up strings from multiple pieces.

    (When you only need to concatenate strings, you can use the `+` operator to do this instead of calling `fmt`.)
