# Error Groups

Error Groups (errgroups) provide a way to manage a group of goroutines and collect the first error that occurs among them. They are implemented in the `golang.org/x/sync/errgroup` package.

## Key Concepts

- **Synchronized goroutine management**: Similar to WaitGroups but with error propagation
- **First error cancellation**: Cancels all goroutines when any one returns an error
- **Context integration**: Can be used with contexts for broader cancellation flows
- **Concurrency limiting**: Can limit the number of concurrent goroutines (in newer versions)

## Basic Usage

```go
import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
)

func main() {
    g, ctx := errgroup.WithContext(context.Background())
    
    // Launch multiple tasks
    for i := 0; i < 10; i++ {
        i := i // Capture loop variable
        g.Go(func() error {
            return processItem(ctx, i)
        })
    }
    
    // Wait for all tasks to complete or an error to occur
    if err := g.Wait(); err != nil {
        fmt.Printf("Error occurred: %v\n", err)
        return
    }
    
    fmt.Println("All tasks completed successfully")
}

func processItem(ctx context.Context, id int) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // Do work
        if id == 5 {
            return fmt.Errorf("processing failed for item %d", id)
        }
        return nil
    }
}
```

## Advantages Over Manual Error Handling

- **Simplified coordination**: No need to manually track errors from goroutines
- **Automatic propagation**: Context cancellation happens automatically on error
- **Clean API**: Cleaner code compared to manual channel-based error handling
- **Consistent pattern**: Encourages a standard way to handle errors in concurrent code

## Error Handling Patterns

- **Early return**: The first error causes all goroutines to be cancelled
- **Error collection**: Collect all errors using a custom error type (requires custom implementation)
- **Error transformation**: Process or wrap errors before returning them

## Exercise Ideas

1. Implement a parallel file processor with error handling
2. Create a service that makes multiple API calls concurrently
3. Build a concurrent data validator with early failure
4. Implement a custom errgroup that collects all errors

## Best Practices

- Use errgroup when you need to run multiple operations concurrently and fail fast
- Pair errgroup with context for proper cancellation
- Handle context cancellation in worker functions
- Design worker functions to check for cancellation at appropriate intervals
- Consider using a custom implementation if you need to collect all errors

## Limitations

- Only returns the first error encountered
- No built-in way to collect all errors (needs custom implementation)
- Requires Go modules or vendoring (part of x/sync, not standard library)
- No priority ordering of tasks (just concurrent execution)

## Next Steps

After understanding errgroups, explore pipeline patterns for building efficient data processing workflows. 