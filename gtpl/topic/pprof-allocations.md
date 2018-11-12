
---
***pprof: Allocations***

The basic tool for profiling Go programs is `pprof`. There are three different version of the `pprof` tool:

- [runtime version](https://golang.org/pkg/runtime/pprof/)
- [net/http version](https://golang.org/pkg/net/http/pprof/)
- [command-line version](https://blog.golang.org/profiling-go-programs)

The command-line version of `pprof` is part of the Go toolchain, and you can run it with `go tool pprof` or just `pprof` in the terminal. We will use this now to look deeper into the allocations of our code.

First we need to create a memory profile. We can do this while running benchmarks:

```
go test -v --bench . --benchmem --memprofile mem.out --memprofilerate=1
```

We now have a `mem.out` file containing a memory profile. We can investigate it with the `pprof` command from the terminal.

```
pprof mem.out
```

Now we are in the `pprof` console. Here we can type specific `pprof` commands. Start with `help` to show a list of all available commands. We can then see what our top memory consumers are with `top5`, or `top5 -cum` to sort them by the cumulative column.

We can also check which code allocates memory, and how much. The command `list` takes a regular expression to search in your code. Try giving it a function name to investigate:

```
list functionName
```

This will show all code where `functionName` was found that involves memory allocation. Before the lines that allocate memory you will see how much memory was allocated.

Remember that we created the profile by running a benchmark, so don't be alarmed to see several megabytes being allocated at once; the code was executed a few thousand times.

Finally, it's easy to generate a nice visualization of the top memory consumers. The `pdf` command will create a diagram and save it as a PDF file. Alternatively, the `web` command will open your browser and show the diagram there.

---