
---
***Welcome to the Go track!***

As this is likely your first Go exercise on Exercism, I thought I'd share a few tips that may be helpful.

1. **Run the tests** one last time before submitting. Usually running just `go test` in the exercise directory is all you need to do here, unless there are special requirements which will be mentioned in the instructions for that exercise.

    If you can't get all the tests to pass, feel free to submit the solution anyway and ask for help. Mentors will be happy to give you hints.

2. Ensure your **code is formatted** with `gofmt`. Most editors that support Go can be configured to do this automatically. If you're coming to Go from other languages, you may not be used to the idea that there's one, and only one, accepted way to format Go code, and it's the `gofmt` way. At first this may seem overly restrictive, but there are great advantages to standard formatting, not least that it avoids a lot of futile arguments about which is the best way to format Go code.

3. Make sure your **code passes `golint`**. `golint` is a tool that analyses your code for common errors and problems, and also enforces things like documentation comments for your functions and identifiers. Again, your editor can usually lint your code automatically, and it's a good idea to set this up. Run `go get -u golang.org/x/lint/golint` to install `golint`, and check your code with it before submitting.

4. **The goal of exercism is to teach fluency**. Among other things this includes guiding you to a **simple, readable and idiomatic** solution. While these are good software engineering principles in general, they're especially important in Go, where the whole language is designed for maximum simplicity and clarity.

    When you have something which works, try to simplify it as much as possible by eliminating all redundant or duplicated code, and rewriting everything in its simplest form. If there are parts of the code which seem awkward or complicated to you, trust your instincts and refactor the code until you feel good about it. In Go, 'clear is better than clever'. Keep that in mind and you won't go far wrong.

5. **Resist the temptation to optimize everything for performance**. Go programs are fast; astonishingly fast, if you're used to interpreted languages. Go also has great performance analysis tools: the benchmarker, the profiler, and so on. These are all great fun to play with, and as engineers we love trying to find the absolutely optimal way to do something. Feel free to do so but always consider simplicity and readability first. If you find a more efficient method which doesn't compromise readability, that is perfect!

6. This may be one of the few occasions in your career when you can get personal, one-to-one help and advice from an experienced software engineer. **Make the most of it, and above all, have fun!**

