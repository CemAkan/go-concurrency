# Pipeline Patterns

Pipeline patterns in Go enable efficient data processing by breaking complex operations into stages connected by channels.

## Key Concepts

- **Stages**: Independent processing units connected by channels
- **Composition**: Building complex pipelines from simple, reusable stages
- **Throughput**: Optimizing overall processing speed
- **Backpressure**: Natural rate limiting as slower stages block faster ones
- **Parallel processing**: Running multiple instances of compute-intensive stages

## Basic Pipeline Structure

A pipeline typically consists of:
1. **Source/Generator**: Produces data for processing
2. **Intermediate stages**: Transform the data
3. **Sink**: Consumes the final results

```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func sink(in <-chan int) []int {
    var result []int
    for n := range in {
        result = append(result, n)
    }
    return result
}

// Usage
func main() {
    // Build pipeline
    source := generator(1, 2, 3, 4, 5)
    squared := square(source)
    result := sink(squared)
    
    fmt.Println(result) // [1, 4, 9, 16, 25]
}
```

## Fan-Out, Fan-In Pipeline

```go
func fanOut(in <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        outputs[i] = square(in) // Each worker reads from the same input
    }
    return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(len(inputs))
    
    for _, ch := range inputs {
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch {
                out <- n
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// Usage
source := generator(1, 2, 3, 4, 5, 6, 7, 8)
workerOutputs := fanOut(source, 3) // 3 parallel square workers
merged := fanIn(workerOutputs...)  // Combine results
result := sink(merged)
```

## Pipeline with Cancellation

```go
func generator(ctx context.Context, nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            select {
            case out <- n:
            case <-ctx.Done():
                return
            }
        }
    }()
    return out
}

// Other stages follow the same pattern
```

## Common Pipeline Patterns

- **Sequential processing**: Simple linear stages
- **Parallel processing**: Multiple instances of compute-intensive stages
- **Filter stages**: Discard some items based on criteria
- **Batch processing**: Group items into batches
- **Dynamic pipelines**: Changing the pipeline structure at runtime

## Exercise Ideas

1. Build a pipeline for processing log files
2. Create an image processing pipeline with multiple transformation stages
3. Implement a data enrichment pipeline that fetches additional information
4. Build a pipeline with dynamic scaling based on workload

## Best Practices

- Design each stage to be independent and reusable
- Use buffered channels where appropriate to smooth processing
- Implement proper error handling at each stage
- Ensure all goroutines can exit when the pipeline is cancelled
- Consider memory usage and backpressure

## Common Mistakes

- Forgetting to close channels
- Not handling cancellation properly
- Creating too many or too few parallel workers
- Ignoring error propagation
- Not considering the performance implications of channel operations

## Next Steps

After understanding pipeline patterns, explore rate limiting to control the flow of operations in concurrent systems. 