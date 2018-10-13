- To compile a static regex it is advisable to use `MustCompile`. Since a static regex is usually best 
compiled at package level and `MustCompile` panics if there is a problem with the regex the developer gets
immediate feedback when trying to execute his program.