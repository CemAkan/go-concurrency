# Mutexes

Mutexes (mutual exclusion locks) are synchronization primitives that protect shared resources from concurrent access by multiple goroutines.

## Key Concepts

- **Exclusive Access**: Only one goroutine can hold the lock at a time
- **Blocking**: Attempts to acquire a locked mutex will block until the mutex is available
- **Data Protection**: Used to prevent race conditions on shared data

## Types of Mutexes in Go

- **sync.Mutex**: Basic mutual exclusion lock
- **sync.RWMutex**: Reader/Writer mutual exclusion lock that allows multiple readers or one writer

## Core Methods

For `sync.Mutex`:
- **Lock()**: Acquires the lock, blocking if necessary
- **Unlock()**: Releases the lock

For `sync.RWMutex`:
- **RLock()**: Acquires a read lock
- **RUnlock()**: Releases a read lock
- **Lock()**: Acquires a write lock (exclusive)
- **Unlock()**: Releases a write lock

## Common Usage Pattern

```go
var (
    mu      sync.Mutex
    counter int
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

## Best Practices

- Always use `defer` with `Unlock()` to prevent deadlocks
- Keep the critical section (locked code) as small as possible
- Use RWMutex when you have more reads than writes
- Avoid nested locks to prevent deadlocks
- Consider using atomic operations for simple cases

## Race Conditions

- Race conditions occur when multiple goroutines access shared data concurrently
- Use the `-race` flag with `go build`, `go run`, or `go test` to detect race conditions

## Exercise Ideas

1. Implement a thread-safe counter
2. Create a concurrent map with mutex protection
3. Compare performance of Mutex vs RWMutex in read-heavy scenarios

## Common Mistakes

- Forgetting to unlock a mutex
- Copying a mutex (they should not be copied after first use)
- Using locks inconsistently
- Creating deadlocks with improper lock ordering
- Unnecessarily large critical sections

## When to Use Mutexes

- When multiple goroutines need to access and modify shared data
- When atomic operations are insufficient
- When channels would be too complex for the synchronization needed

## Next Steps

After understanding mutexes, explore the Select statement for managing multiple channel operations. 