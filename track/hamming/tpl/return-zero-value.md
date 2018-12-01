- If returning an `error` it is custom to return the [zero value](https://golang.org/ref/spec#The_zero_value) on all other values.
Returning `-1` for the hamming distance is often done in other programming languages to indicate an error.
In Go this is uncalled-for.
