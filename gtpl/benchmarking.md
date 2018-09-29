
---
***Benchmarking***

I have made some suggestions that involve benchmarking. In case you already know about benchmarking you 
can skip this.

Benchmarks already exist for this exercise. You just need to execute them with:
```
go test -v --bench . --benchmem
```

This will first run the tests and then the benchmarks outputting something like this:

```
goos: darwin
goarch: amd64
pkg: path/to/exercise
BenchmarkName-8  100000  1859 ns/op  57 B/op  7 allocs/op
PASS
ok  path/to/exercise   2.042s
```

For each benchmark there will be a line starting with the `BenchmarkName` followed by the number of 
cores `8` available to the benchmark. The next number `100000` indicates how many times (`operations`) 
the benchmark was executed.

The next three numbers indicate `speed` in `ns/op` (nanoseconds per operation), `memory` in `B/op` 
(bytes of memory allocated per operation) and `allocations` in `alloc/op` (amount of memory allocations per 
operation). *"Lower is better"* applies to these three numbers.

The last number `2.042s` indicates the `total execution time` for all tests and benchmarks. It is *not* 
significant for measuring speed! The benchmarking tool in go executes faster benchmarks more often to produce 
more reliable results -- possibly increasing the total execution time.

---