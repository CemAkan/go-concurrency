# Time Management in Concurrency

Time management is a critical aspect of concurrent programming in Go, enabling operations like timeouts, delays, and rate limiting.

## Key Concepts

- **Timeouts**: Preventing operations from blocking indefinitely
- **Rate limiting**: Controlling the frequency of operations
- **Delays**: Pausing execution for a specified duration
- **Tickers**: Performing operations at regular intervals
- **Deadlines**: Executing operations before a specific time

## Core Time Utilities

- **time.After**: Returns a channel that receives a value after a specified duration
- **time.Tick**: Returns a channel that delivers tick events at regular intervals
- **time.Timer**: Represents a single event in the future
- **time.Ticker**: Represents a repeated event at regular intervals
- **time.Sleep**: Pauses execution for a specified duration

## Timeout Pattern

```go
func timeoutOperation(timeout time.Duration) (Result, error) {
    resultCh := make(chan Result, 1)
    errCh := make(chan error, 1)
    
    go func() {
        result, err := performOperation()
        if err != nil {
            errCh <- err
            return
        }
        resultCh <- result
    }()
    
    select {
    case result := <-resultCh:
        return result, nil
    case err := <-errCh:
        return Result{}, err
    case <-time.After(timeout):
        return Result{}, errors.New("operation timed out")
    }
}
```

## Rate Limiting Pattern

```go
func rateLimitedOperations() {
    limiter := time.Tick(200 * time.Millisecond) // 5 operations per second
    
    for request := range requests {
        <-limiter // Wait for next tick
        go process(request)
    }
}
```

## Scheduled Operations

```go
func scheduledTask() {
    ticker := time.NewTicker(1 * time.Hour)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            performPeriodicTask()
        case <-quit:
            return
        }
    }
}
```

## Combining Time with Context

```go
func timeConstrainedOperation(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    return performOperationWithContext(ctx)
}
```

## Common Time Management Patterns

- **Heartbeats**: Regular signals indicating a goroutine is still alive
- **Debouncing**: Waiting until a quiet period before acting on rapid events
- **Throttling**: Limiting the rate of operations
- **Timeouts**: Cancelling operations that take too long
- **Retries with backoff**: Increasing delay between retry attempts

## Exercise Ideas

1. Implement a rate limiter for API requests
2. Create a retry mechanism with exponential backoff
3. Build a service with configurable timeouts
4. Implement a job scheduler with time-based triggers

## Best Practices

- Always use timeouts for I/O operations
- Consider the precision needed for timing operations
- Use context for propagating deadlines through call chains
- Clean up timers and tickers when they're no longer needed
- Be aware of timer/ticker memory usage in long-running programs

## Common Mistakes

- Forgetting to stop Tickers (memory leak)
- Using time.Sleep in goroutines without cancellation
- Incorrect handling of timer events
- Not considering timer drift for long-running operations
- Using the zero Time value incorrectly

## Next Steps

After mastering time management, explore error groups for managing concurrent errors in groups of goroutines. 