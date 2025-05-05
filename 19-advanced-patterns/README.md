# Advanced Concurrency Patterns

Advanced concurrency patterns combine multiple Go concurrency primitives to solve complex real-world problems, often addressing considerations like safety, performance, and resource management.

## Bounded Concurrency Pattern

Control the maximum number of concurrent operations while still processing all items.

```go
func boundedExecution(items []Item, concurrency int, process func(Item) Result) []Result {
    semaphore := make(chan struct{}, concurrency)
    results := make([]Result, len(items))
    
    var wg sync.WaitGroup
    wg.Add(len(items))
    
    for i, item := range items {
        i, item := i, item // Create local copies for goroutine
        
        go func() {
            defer wg.Done()
            
            semaphore <- struct{}{} // Acquire token
            defer func() { <-semaphore }() // Release token
            
            results[i] = process(item)
        }()
    }
    
    wg.Wait()
    return results
}
```

## Circuit Breaker Pattern

Prevent cascading failures by temporarily disabling operations after errors.

```go
type CircuitBreaker struct {
    mu                sync.RWMutex
    failureThreshold  uint
    resetTimeout      time.Duration
    failureCount      uint
    lastFailure       time.Time
    state             State // Closed, Open, HalfOpen
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    if !cb.AllowRequest() {
        return ErrCircuitOpen
    }
    
    err := fn()
    
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if err != nil {
        cb.failureCount++
        cb.lastFailure = time.Now()
        
        if cb.failureCount >= cb.failureThreshold {
            cb.state = Open
        }
        
        return err
    }
    
    if cb.state == HalfOpen {
        cb.state = Closed
        cb.failureCount = 0
    }
    
    return nil
}
```

## Bulkhead Pattern

Isolate components to prevent failures from cascading throughout the system.

```go
type Bulkhead struct {
    workers map[string]chan struct{}
    mu      sync.RWMutex
}

func NewBulkhead() *Bulkhead {
    return &Bulkhead{
        workers: make(map[string]chan struct{}),
    }
}

func (bh *Bulkhead) Configure(name string, capacity int) {
    bh.mu.Lock()
    defer bh.mu.Unlock()
    bh.workers[name] = make(chan struct{}, capacity)
}

func (bh *Bulkhead) Execute(name string, fn func() error) error {
    bh.mu.RLock()
    sem, exists := bh.workers[name]
    bh.mu.RUnlock()
    
    if !exists {
        return fmt.Errorf("bulkhead %s not configured", name)
    }
    
    select {
    case sem <- struct{}{}:
        defer func() { <-sem }()
        return fn()
    default:
        return ErrBulkheadFull
    }
}
```

## Timeout Context with Cleanup Pattern

Provide timeout for operations while ensuring resources are properly cleaned up.

```go
func operationWithTimeout(timeout time.Duration, operation func(context.Context) (Result, error)) (Result, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel() // Ensure resources are cleaned up
    
    resultCh := make(chan Result, 1)
    errCh := make(chan error, 1)
    
    go func() {
        result, err := operation(ctx)
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
    case <-ctx.Done():
        return Result{}, ctx.Err()
    }
}
```

## Error Group with Cancellation Pattern

Run multiple operations concurrently, but cancel all if any fail.

```go
func fetchAllWithCancellation(ctx context.Context, urls []string) ([]Response, error) {
    g, ctx := errgroup.WithContext(ctx)
    responses := make([]Response, len(urls))
    
    for i, url := range urls {
        i, url := i, url // Local copies for goroutine
        
        g.Go(func() error {
            resp, err := fetchWithContext(ctx, url)
            if err != nil {
                return err // Will cancel context for all goroutines
            }
            responses[i] = resp
            return nil
        })
    }
    
    if err := g.Wait(); err != nil {
        return nil, err
    }
    
    return responses, nil
}
```

## Debouncing Pattern

Wait until input stops changing before processing.

```go
func debounce(f func(arg interface{}), wait time.Duration) func(arg interface{}) {
    var mutex sync.Mutex
    var timer *time.Timer
    
    return func(arg interface{}) {
        mutex.Lock()
        defer mutex.Unlock()
        
        if timer != nil {
            timer.Stop()
        }
        
        timer = time.AfterFunc(wait, func() {
            f(arg)
        })
    }
}
```

## Exercise Ideas

1. Implement a service with circuit breaker and bulkhead patterns
2. Create a search aggregator using bounded concurrency
3. Build a rate-limited API client with retry and timeout patterns
4. Design a real-time event processor with debouncing

## Best Practices

- Combine patterns judiciously to address specific requirements
- Test combined patterns thoroughly under various scenarios
- Consider the maintainability of complex patterns
- Document the intent and behavior of advanced patterns
- Use well-established libraries when available rather than implementing from scratch

## Common Pitfalls

- Overcomplicating solutions with too many patterns
- Race conditions from incorrect synchronization
- Resource leaks from improper cleanup
- Deadlocks from circular dependencies
- Performance bottlenecks from excessive synchronization

## Further Learning Resources

1. "Concurrency in Go" by Katherine Cox-Buday
2. "Go Concurrency Patterns" talk by Rob Pike
3. "Advanced Go Concurrency Patterns" talk by Sameer Ajmani
4. "Go in Action" by William Kennedy
5. The Go Blog series on concurrency 