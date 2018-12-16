
---
***Concurrency: What's the point?***

If you've benchmarked your solution to this exercise, you might be mildly surprised that the concurrent version isn't much faster than the sequential one. Why not?

The point of this exercise is that using concurrency does not always improve performance. Sometimes it actually makes things worse! In this exercise, you can get the concurrent version to run about as fast as the sequential one; maybe slightly faster. But concurrency isn't free; we have to do a fair bit of work to organise the concurrent goroutines and pass the data and results around between them. Also, counting letter frequency is a relatively fast task; merging maps is slow. So we are actually parallelising only a small part of the work here.

Using concurrency is most beneficial when you can break up a task into multiple chunks that can be processed independently, and when the process of merging the chunked results isn't too expensive. Keep this in mind when you're thinking about whether a concurrent approach to solving a given problem makes sense or not.

But concurrency isn't just about speed. It's also about dealing with multiple things at once; for example, handling requests to a web server. Processing requests concurrently means that you can get very high throughput (requests per second). Whenever you're dealing with multiple simultaneous requests, a concurrent solution is probably the way to go.

To learn more about concurrency in Go and how to use it effectively, watch Rob Pike's talk on [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs).

---
