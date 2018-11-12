
---
***Benchmarking***

I have made some suggestions that involve benchmarking (testing the performance and efficiency of your code). If you're already familiar with benchmarking you can skip this.

Benchmarks already exist for this exercise. You can execute them with:
```
go test -v --bench . --benchmem
```

This will first run the tests, and then the benchmarks, producing something like this:

```
goos: darwin
goarch: amd64
pkg: path/to/exercise
BenchmarkName-8  100000  1859 ns/op  57 B/op  7 allocs/op
PASS
ok  path/to/exercise   2.042s
```

For each benchmark there will be a line starting with the `BenchmarkName` followed by the number of cores (`8`) available. The next number (`100000`) indicates how many times the benchmark was run. Go will automatically work out how many times to run the benchmark to get a statistically useful result.

The next three numbers indicate:

1. the time it took to execute the benchmark, in `ns/op` (nanoseconds per operation)

1. the memory usage, in `B/op` (bytes of memory allocated per operation)

1. the number of memory allocations, in `alloc/op`

For all these numbers, lower is better.

The last number `2.042s` indicates the `total execution time` for all tests and benchmarks. Ignore this. This is *not* significant for measuring speed! The benchmarking tool in Go executes faster benchmarks more often to produce more reliable results, possibly increasing the total execution time.

Dave Cheney has written a good blog post on Go benchmarking which you may find interesting: [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

---