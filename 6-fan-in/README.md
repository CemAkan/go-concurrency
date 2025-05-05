# Fan-In Pattern

The Fan-In pattern is a concurrency design pattern where multiple goroutines send data to a single channel, combining multiple data streams into one.

## Key Concepts

- **Multiplexing**: Combining multiple input channels into a single output channel
- **Consolidation**: Gathering results from multiple concurrent operations
- **Throughput**: Potentially improving throughput by parallelizing operations

## Basic Pattern

```go
func fanIn(channels ...<-chan T) <-chan T {
    var wg sync.WaitGroup
    multiplexed := make(chan T)
    
    // Function to forward values from input channels to output channel
    multiplex := func(c <-chan T) {
        defer wg.Done()
        for value := range c {
            multiplexed <- value
        }
    }
    
    // Start a goroutine for each input channel
    wg.Add(len(channels))
    for _, c := range channels {
        go multiplex(c)
    }
    
    // Close the multiplexed channel when all input channels are done
    go func() {
        wg.Wait()
        close(multiplexed)
    }()
    
    return multiplexed
}
```

## Use Cases

- **Gathering results**: Collecting results from multiple workers
- **Service aggregation**: Combining responses from multiple services
- **Event handling**: Multiplexing events from different sources
- **Data processing**: Merging streams from different data sources

## Variations

- **Ordered Fan-In**: Preserving the order of results
- **Batched Fan-In**: Collecting results in batches instead of one by one
- **Priority Fan-In**: Prioritizing certain input channels over others
- **Timed Fan-In**: Collecting results within a time limit

## Challenges

- **Managing resource cleanup**: Properly closing channels and goroutines
- **Handling errors**: Propagating errors from worker goroutines
- **Cancellation**: Stopping all goroutines when needed
- **Backpressure**: Dealing with slow consumers

## Exercise Ideas

1. Implement a basic fan-in function
2. Create a web scraper that fans in results from multiple URLs
3. Build a system that processes data from multiple sources in parallel

## Best Practices

- Use a WaitGroup to track when all input channels are done
- Consider buffered channels for better performance in some cases
- Propagate cancellation to all worker goroutines
- Handle errors explicitly

## Common Mistakes

- Not closing the output channel, leading to deadlocks
- Forgetting to handle the case when input channels close
- Creating too many goroutines for the available resources
- Not considering channel buffer sizes for performance

## Next Steps

After understanding fan-in, explore the complementary fan-out pattern for distributing work across multiple goroutines. 