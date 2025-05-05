# sync.Once

The `sync.Once` type in Go provides a safe way to perform one-time initialization in a concurrent environment, ensuring that an initialization function is executed exactly once, even when called from multiple goroutines.

## Key Concepts

- **One-time execution**: Guarantees that a function will be executed exactly once
- **Thread safety**: Safe for concurrent use from multiple goroutines
- **Laziness**: Initialization happens on demand, not at program start
- **Blocking**: All callers wait for the first execution to complete

## Basic Usage

```go
import (
    "fmt"
    "sync"
)

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        fmt.Println("Creating singleton instance (happens only once)")
        instance = &Singleton{}
        // Perform initialization of the instance
        instance.Init()
    })
    return instance
}
```

## How sync.Once Works

- Internally uses mutual exclusion (mutex) and a "done" flag
- First caller executes the initialization function
- Subsequent callers check the flag and skip execution
- All callers are blocked until the first execution completes

## Common Use Cases

- **Singleton pattern**: Ensuring only one instance of a type is created
- **Lazy initialization**: Creating expensive resources only when needed
- **Package initialization**: One-time setup for packages
- **Connection pooling**: Initializing connection pools
- **Configuration loading**: Loading configuration files once

## Comparison to Other Techniques

- **init() functions**: Run at program startup, not lazily
- **Global variables with initialization**: Not concurrency-safe
- **Mutexes with flags**: Requires more boilerplate code
- **Atomic operations**: Lower-level, more error-prone

## Exercise Ideas

1. Implement a lazy-loaded configuration manager
2. Create a singleton database connection pool
3. Build a cache that initializes on first access
4. Design a plugin system with one-time initialization

## Best Practices

- Use `sync.Once` for expensive initializations that should happen only once
- Keep the initialization function focused on one responsibility
- Don't nest `sync.Once` calls
- Don't reuse `sync.Once` instances for different initializations
- Remember that panics in the initialization function are propagated to all callers

## Patterns with sync.Once

### Reset Pattern (Workaround)

```go
type ResettableOnce struct {
    m     sync.Mutex
    done  uint32
}

func (o *ResettableOnce) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    
    o.m.Lock()
    defer o.m.Unlock()
    
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}

func (o *ResettableOnce) Reset() {
    o.m.Lock()
    defer o.m.Unlock()
    atomic.StoreUint32(&o.done, 0)
}
```

### Lazy Singleton

```go
type Singleton struct {
    // fields
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{
            // Initialize fields
        }
    })
    return instance
}
```

## Common Mistakes

- Calling `Do()` with different functions, expecting them all to execute
- Trying to reset `sync.Once` (it's not designed to be reset)
- Forgetting that `Do()` propagates panics
- Using a global `sync.Once` for unrelated initializations
- Creating complex dependencies in initialization functions

## Next Steps

After mastering `sync.Once`, explore advanced concurrency patterns that combine multiple synchronization primitives for complex scenarios. 