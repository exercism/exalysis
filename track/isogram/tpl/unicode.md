- You could look at using `unicode.%[1]s` inside the `for` loop instead of `strings.%[1]s` before the loop, to increase speed.

    (Sometimes it makes the code a little easier to read if you call `strings.%[1]s` beforehand, though. Use your best judgment, and remember readability beats performance!)
