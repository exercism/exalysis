
---
***Welcome to the Go track!***

As this is likely your first Go exercise on Exercism, I thought I'd share a few tips that may be helpful. Following these steps will make sure that your solutions are approved quickly, and will save time for both you and the volunteer mentors who will be reviewing your code.

1. Always run the tests. Usually running just `go test` in the exercise directory is all you need to do here, but if there are any special requirements for an exercise, you'll find them in the instructions for that exercise.

    Any solution which does not pass the tests will be bounced straight back to you by the mentor, so to avoid this, do one final `go test` before submitting, just to make sure.

    If you can't get all the tests to pass, and you've made a genuine attempt to fix the problem yourself, feel free to submit the solution anyway and ask for help. Mentors will be happy to give you hints if you can show them that you've put some effort into your solution.

2. Always ensure your code is formatted with `gofmt`. Most editors that support Go can be configured to do this automatically. If you're coming to Go from other languages, you may not be used to the idea that there's one, and only one, accepted way to format Go code, and it's the `gofmt` way. At first this may seem overly restrictive, but there are great advantages to standard formatting, not least that it avoids a lot of futile arguments about which is the best way to format code!

    Again, mentors will expect you to `gofmt` your code, and anything which is not correctly formatted will be bounced back to you. Don't waste the mentors' time and patience by arguing with them about this. If you want to join the worldwide community of Gophers, that means accepting community standards, many of which may seem silly or pointless to you at first. Stick with it; you may feel differently about things further down the line.

3. Always make sure your code passes `golint`. `golint` is a tool that analyses your code for common errors and problems, and also enforces things like documentation comments for your functions and identifiers. Again, your editor can usually lint your code automatically, and it's a good idea to set this up. Run `go get -u golang.org/x/lint/golint` to install `golint`, and check your code with it before submitting.

    Like all automated tools, `golint` is not perfect, and there are occasions where you can safely ignore its advice, but you won't encounter any of them in Exercism exercises. Therefore, please follow all `golint`'s suggestions for now. Some of them may seem silly or excessive, but again, stick with it for now. You are building habits which will stand you in good stead later.

    One of the first things `golint` is likely to complain about is missing documentation comments on your functions. Every Go function should have a comment above it explaining what the function does, what arguments it takes, and what it returns. Don't skip this or write silly comments to silence `golint`, like "Square yada yada yada." Sure, it's just an exercise, and you don't really need documentation, but you're learning the skills and habits to be a successful Go programmer, and good documentation is one of them.

4. Assuming your code is correctly formatted and linted, and passes tests, will it be automatically approved? Not necessarily. Mentors are looking for a number of things apart from merely correct code. One of them is 'idiomatic' code; that is, code which uses the language and library in a simple, natural way, rather than working against it. The more idiomatic your code, the easier and more straightforward it will be for someone to read and maintain it.

    Mentors are also looking for code which is simple, clear, and straightforward. While these are good software engineering principles in general, they're especially important in Go, where the whole language is designed for maximum simplicity and clarity. Avoid the temptation to 'over-solve' problems by building complicated structures and abstractions which are not absolutely necessary for this specific case.

    When you have something which works, try to simplify it as much as possible by eliminating all redundant or duplicated code, and rewriting everything in its simplest form. If there are parts of the code which seem awkward or complicated to you, trust your instincts and refactor the code until you feel good about it. 'Ugly' code is usually bad code. Good code tends to be simple, straightforward, readable, minimal, and even elegant.

    In Go, 'clear is better than clever'. Keep that in mind and you won't go far wrong.

5. Resist the temptation to optimize everything for performance. Go programs are fast; astonishingly fast, if you're used to interpreted languages. Go also has great performance analysis tools: the benchmarker, the profiler, and so on. These are all great fun to play with, and as engineers we love trying to find the absolutely optimal way to do something.

    But bear in mind that performance simply isn't an issue for most real programs. Most programs spend nearly all their time waiting for user input, waiting for network or disk I/O, and so on. Raw speed is occasionally important for very small and special sections of code (the 'hot path'), and you can deal with those when you come to them. There are no performance requirements for any Exercism exercises, so concentrate instead on readability and simplicity. If there is a natural, straightforward solution to a problem, use it. Mentors are looking for good code, not fast code.

    Of course, there are pathologically bad ways to solve a problem, and if there is a very much more efficient method which doesn't compromise readability, mentors will let you know about it. Don't take this as meaning that you should spend a great deal of time investigating and tuning the performance of your code. By all means learn to use the benchmarking tools, and learn about the way your choices affect performance. But remember, developer time is a lot more expensive than CPU time, and it's no good optimizing for the latter at the expense of the former.

6. Mentoring, or any kind of code review, can occasionally be a frustrating process. It may seem to you that the mentor is being unnecessarily picky, or forcing you to go in a direction that doesn't seem natural to you. When you've submitted five or ten iterations of a solution, and the mentor is still asking for changes, you may feel that you just want to get on and do a different exercise.

    It's worth remembering that all the mentors on Exercism are volunteers, working unpaid and in their spare time, purely out of their desire to help people like you. It would be much easier from their point of view to simply approve every solution that passes tests, and move you on to the next exercise. But that would miss the point. The real value of mentoring is that you can get advice and suggestions from someone who's been where you are, knows many of the pitfalls, and has gained a lot of useful experience that you don't yet have. If you had to pay for this advice, it would probably cost you a great deal of money (think about how much it costs to get flying lessons, or even driving lessons). The fact that you're getting it for free shouldn't make you undervalue it.

    That said, mentors are human beings too, and they can get bored, tired, stressed, or frustrated, just like you. They do their best to be friendly, supportive, patient, and helpful. On the rare occasions when they fall below these standards, be forgiving. If you're feeling stuck on a particular exercise, tell the mentor so, and ask them what you need to do to get your solution up to the level where they'll approve it. Don't pester them to approve your solution. If they have asked for changes, don't ignore them or argue with them, just make the changes. You may feel you know better, but put that feeling aside for now. If you don't understand why a particular change makes the code better, ask. Some things are just "Go style", and this will become more natural to you as you gain experience with writing Go. Instead of insisting on your own personal style, listen to the mentor and trust in their experience.

    This may be one of the few occasions in your career when you can get personal, one-to-one help and advice from an experienced software engineer. Make the most of it, and above all, remember to have fun!

