# Semaphores

Semaphores are synchronization primitives used to control access to a common resource in a concurrent environment. While Go doesn't have a built-in semaphore type, they can be implemented using channels or the `golang.org/x/sync/semaphore` package.

## Key Concepts

- **Limited concurrent access**: Restricting the number of goroutines that can access a resource
- **Counting semaphore**: Tracks a count of available resources
- **Binary semaphore**: Special case with only one available resource (similar to a mutex)
- **Acquire/Release**: Core operations to obtain and return access to the resource

## Channel-Based Semaphore Implementation

```go
// Simple semaphore implementation using a buffered channel
type Semaphore chan struct{}

// Acquire n resources
func (s Semaphore) Acquire(n int) {
    for i := 0; i < n; i++ {
        s <- struct{}{}
    }
}

// Release n resources
func (s Semaphore) Release(n int) {
    for i := 0; i < n; i++ {
        <-s
    }
}

// Create a new semaphore with a maximum count
func NewSemaphore(maxCount int) Semaphore {
    return make(Semaphore, maxCount)
}

// Usage
func main() {
    // Create a semaphore allowing 3 concurrent operations
    sem := NewSemaphore(3)
    
    // Use in a worker pool
    for i := 0; i < 10; i++ {
        go func(id int) {
            sem.Acquire(1)
            defer sem.Release(1)
            
            // Do work with limited concurrency
            fmt.Printf("Worker %d is processing\n", id)
            time.Sleep(1 * time.Second)
        }(i)
    }
}
```

## Using golang.org/x/sync/semaphore

```go
import (
    "context"
    "fmt"
    "golang.org/x/sync/semaphore"
    "time"
)

func main() {
    // Create a weighted semaphore with a maximum count of 3
    sem := semaphore.NewWeighted(3)
    ctx := context.Background()
    
    for i := 0; i < 10; i++ {
        // Acquire 1 resource
        if err := sem.Acquire(ctx, 1); err != nil {
            fmt.Printf("Failed to acquire semaphore: %v\n", err)
            break
        }
        
        go func(id int) {
            defer sem.Release(1)
            
            // Do work with limited concurrency
            fmt.Printf("Worker %d is processing\n", id)
            time.Sleep(1 * time.Second)
        }(i)
    }
}
```

## Common Use Cases

- **Connection pooling**: Limiting the number of concurrent connections
- **Database access**: Restricting concurrent database operations
- **API rate limiting**: Limiting concurrent API calls
- **Resource-intensive operations**: Controlling CPU or memory usage
- **File descriptor limits**: Managing limited OS resources

## Variations

- **Weighted semaphore**: Different operations consume different amounts of the resource
- **Timed semaphore**: Automatic release after a timeout
- **Try-acquire**: Non-blocking attempt to acquire the semaphore
- **Fair semaphore**: Ensures FIFO ordering of waiting goroutines

## Exercise Ideas

1. Implement a connection pool using semaphores
2. Create a resource-limited web crawler
3. Build a task scheduler with resource constraints
4. Implement a fair semaphore with FIFO ordering

## Best Practices

- Choose an appropriate maximum count based on system resources
- Always release acquired semaphores, typically using `defer`
- Consider using contexts for cancellation
- Use semaphores for resource constraints, not for mutual exclusion (use mutex for that)
- Be careful with deadlocks when acquiring multiple semaphores

## Common Mistakes

- Forgetting to release semaphores
- Using semaphores when mutexes would be more appropriate
- Setting semaphore limits too high or too low
- Creating deadlocks by acquiring multiple semaphores in different orders
- Not handling the context cancellation error when acquiring

## Next Steps

After learning about semaphores, explore atomic operations for lock-free synchronization of simple values. 