
---
***Concurrency: What is the point?***

The point of this exercise is that using concurrency does not always help in regards to speed. You can get the concurrent function as fast as the other one -- maybe slightly faster. But the cost of concurrency is too high - in comparison to the actual work that needs to be done - to be effective. Parallelizing a single task is often not worth it or even counter productive.

Still Go is fast if it does many small tasks concurrently like e.g. in a web server. The single task might not gain from Go's concurrency but you can still get more throughput as the web server as a whole can work on multiple requests at once.

---
