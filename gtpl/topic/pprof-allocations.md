
---
***pprof: Allocations***

Profiling is a huge subject and there are many powerful ways to profile go applications. Following I will 
describe a very limited part of it.

The basis for profiling go programs is `pprof`. There are 3 different version of the `pprof` tool:
- [runtime version](https://golang.org/pkg/runtime/pprof/)
- [net http version](https://golang.org/pkg/net/http/pprof/)
- [pprof command](https://blog.golang.org/profiling-go-programs)

The pprof command version is built right into the go toolchain and can be accessed
with `go tool pprof` or just `pprof` in the terminal. We will use this now to look deeper into the 
allocations of our code.

First we need to create a memory profile. We can do this while running benchmarks:

```
go test -v --bench . --benchmem --memprofile mem.out --memprofilerate=1
```

We have a `mem.out` file now containing a memory profile. We can investigate it with the `pprof` command 
right from the terminal.

```
pprof mem.out
```

Now we are in the `pprof` console. Here we can type specific `pprof` commands. 
Best start with `help` to show a list of all available commands.
We could then have a look at what out top memory comsumers are with `top5` or `top5 -cum` to sort them be 
cumulated column.

We can also check which code exactly allocates memory and how much. For this there is a command `list` which 
takes a regular expression to search in your code. Enter for example the function name you want to investigate: 

```
list functionName
```

This will output all code where `functionName` was found and that allocates memory. Before the lines that 
allocate memory you will see how much memory was allocated. Remember that we created the profile with 
benchmarks. So don't be alarmed to see several Megabyte being allocated at a single line. The code was 
executed a few thousand times. 

Closing I want to mention how easy it is to generate nice diagrams from the top memory consumers.
`pdf` will create a diagram and save it as pdf file. `web` will open the browser and show the diagram there.

---