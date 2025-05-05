# Cancellation Patterns

Cancellation patterns in Go provide mechanisms to stop goroutines gracefully, preventing resource leaks and ensuring proper cleanup.

## Key Concepts

- **Cooperative cancellation**: Goroutines check for cancellation signals and stop themselves
- **Propagation**: Cancellation signals are typically propagated through a call tree
- **Resource cleanup**: Proper cancellation ensures resources are freed
- **Early termination**: Avoiding unnecessary work when results are no longer needed

## Common Cancellation Patterns

### Done Channel Pattern

```go
func worker(done <-chan struct{}, tasks <-chan Task) {
    for {
        select {
        case <-done:
            return // Exit when done signal received
        case task, ok := <-tasks:
            if !ok {
                return // Exit when tasks channel is closed
            }
            // Process task
            process(task)
        }
    }
}

// Usage
done := make(chan struct{})
go worker(done, tasks)

// Later, to cancel:
close(done)
```

### Context-Based Cancellation

```go
func worker(ctx context.Context, tasks <-chan Task) {
    for {
        select {
        case <-ctx.Done():
            return // Exit when context is cancelled
        case task, ok := <-tasks:
            if !ok {
                return
            }
            // Process task
            process(task)
        }
    }
}

// Usage
ctx, cancel := context.WithCancel(context.Background())
go worker(ctx, tasks)

// Later, to cancel:
cancel()
```

## Cancellation Reasons

- **Timeout**: Operation taking too long
- **User request**: User explicitly cancels an operation
- **Error**: An error occurs that makes continuing pointless
- **Parent cancellation**: Parent context is cancelled
- **Resource constraints**: System running low on resources

## Propagation Patterns

- **Explicit passing**: Passing cancellation channels or contexts to child goroutines
- **Context chains**: Creating child contexts that inherit cancellation
- **Middleware**: Wrapping operations with cancellation-aware middleware

## Cleanup Considerations

- **Releasing resources**: Closing files, network connections, etc.
- **Stopping other goroutines**: Propagating cancellation to children
- **State consistency**: Ensuring system state remains consistent after partial completion

## Exercise Ideas

1. Implement a timeout for a long-running operation
2. Create a service with graceful shutdown
3. Build a pipeline with cancellation propagation
4. Implement a web server with request cancellation

## Best Practices

- Always provide a way to cancel long-running operations
- Respect cancellation promptly to free resources
- Check for cancellation at appropriate intervals
- Combine cancellation with timeouts where appropriate
- Propagate cancellation to child operations

## Common Mistakes

- Ignoring cancellation signals
- Not propagating cancellation to sub-operations
- Resource leaks due to lack of cleanup
- Using channels incorrectly for cancellation
- Not implementing timeouts for operations that could block indefinitely

## Next Steps

After understanding cancellation, learn about the Context package, which provides a standardized way to handle cancellation, deadlines, and request-scoped values. 