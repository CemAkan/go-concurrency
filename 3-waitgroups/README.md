# WaitGroups

WaitGroups are a synchronization primitive from the `sync` package used to wait for a collection of goroutines to finish their execution.

## Key Concepts

- **Counter-based**: WaitGroups use a counter to track active goroutines
- **Blocking**: `.Wait()` blocks until the counter reaches zero
- **Thread-safe**: WaitGroups can be safely accessed from multiple goroutines

## Core Functions

- **Add(delta int)**: Increases the counter by the specified amount
- **Done()**: Decreases the counter by one
- **Wait()**: Blocks until the counter becomes zero

## Common Usage Pattern

```go
var wg sync.WaitGroup

// For each goroutine
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()

// Wait for all goroutines to complete
wg.Wait()
```

## Best Practices

- Call `Add()` before starting goroutines
- Use `defer wg.Done()` at the start of goroutine functions to ensure it gets called
- Ensure `Add()` and `Done()` are balanced
- Pass WaitGroup to functions by pointer, but be careful about variable capture in loops

## Challenges

- Forgetting to call `Done()`
- Calling `Wait()` before all goroutines have started
- Calling `Done()` more than `Add()`
- Incorrect handling of WaitGroups in loops

## Exercise Ideas

1. Create a program that spawns multiple goroutines and waits for all to complete
2. Implement a parallel file processor that uses WaitGroups to track completion
3. Combine WaitGroups with channels for collecting results from concurrent operations

## Common Mistakes

- Copying a WaitGroup (should be passed by pointer)
- Calling `Add()` after goroutines have already started
- Not using `defer` for `Done()` calls
- Using the wrong initial value in `Add()`

## When to Use WaitGroups

- When you need to wait for a group of goroutines to complete
- When you don't need communication between goroutines
- When you're implementing a "fork-join" pattern

## Next Steps

After mastering WaitGroups, explore Mutexes for safe access to shared data between goroutines. 