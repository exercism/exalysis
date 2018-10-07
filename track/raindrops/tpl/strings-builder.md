- I see you are using a `strings.Builder`. In this exercise that seems a bit overkill and `+=` actually 
consumes ~2.5 times less memory and is a bit faster. You can run the `benchmarks` as described in the 
exercise instructions and try it yourself. 