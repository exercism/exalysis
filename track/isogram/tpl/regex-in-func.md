- If coming from other languages you probably realized that regex in go works a bit different. Why? The go 
developers put the slow parts of parsing the regex and preparing it as much as possible for matching into a 
_compilation_ step. As a developer you should execute the `Compile`/`MustCompile` function as little as 
possible.

  Right now it is inside your function so it is executed every time the function is called. If dealing with 
a static regex you can put the compilation on package level.