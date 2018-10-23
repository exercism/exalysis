- The iteration order of maps in go is not fix. This is due to the nature of hash maps: inserting items 
creates an arbitrary order by definition. To help developers realize this with their test data even if the 
test data is the same every time go's developers intentionally added some randomizition to the iteration of 
maps in go. For more information and an example read the `Iteration order` section of the 
[go maps in action](https://blog.golang.org/go-maps-in-action) article.