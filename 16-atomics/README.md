# Atomic Operations

Atomic operations provide low-level, lock-free primitives for concurrent programming in Go, allowing thread-safe access to shared variables without the overhead of mutexes.

## Key Concepts

- **Indivisible operations**: Execute completely or not at all, with no partial execution
- **Lock-free**: Do not require explicit locks like mutexes
- **Hardware-supported**: Implemented using CPU atomic instructions
- **Performance**: Typically faster than mutex-based synchronization for simple operations
- **Limited scope**: Best suited for simple data types and operations

## Types Supported in sync/atomic

- **integers**: `int32`, `int64`, `uint32`, `uint64`, `uintptr`
- **pointers**: Unsafe pointers via `unsafe.Pointer`
- **value**: Generic value type via `atomic.Value` (Go 1.4+)

## Core Operations

For integer types:
- **Load**: Read a value atomically
- **Store**: Write a value atomically
- **Add**: Add a value and return the new value
- **Swap**: Replace a value and return the old value
- **CompareAndSwap**: Replace a value only if it matches an expected value

For pointers:
- **LoadPointer**: Read a pointer atomically
- **StorePointer**: Write a pointer atomically
- **SwapPointer**: Replace a pointer and return the old pointer
- **CompareAndSwapPointer**: Replace a pointer only if it matches an expected pointer

For `atomic.Value`:
- **Load**: Read a value atomically
- **Store**: Write a value atomically

## Basic Usage Examples

### Atomic Counter

```go
var counter int64

// Increment counter atomically
atomic.AddInt64(&counter, 1)

// Read counter atomically
value := atomic.LoadInt64(&counter)
```

### Atomic Flag (Similar to sync.Once)

```go
var initialized uint32

func initialize() {
    // Only execute initialization once
    if atomic.CompareAndSwapUint32(&initialized, 0, 1) {
        // Do initialization work
        performInitialization()
    }
}
```

### Atomic Value

```go
var config atomic.Value

// Update configuration atomically
func updateConfig(newConfig *Config) {
    config.Store(newConfig)
}

// Read configuration atomically
func getConfig() *Config {
    return config.Load().(*Config)
}
```

## When to Use Atomics

- **Simple shared counters**: Incrementing/decrementing values
- **Flags and state variables**: On/off switches, status indicators
- **Performance-critical code**: When mutex overhead is a concern
- **Wait-free algorithms**: Building algorithms that don't block
- **Low-contention scenarios**: When conflicts are rare

## When NOT to Use Atomics

- **Complex data structures**: Use mutexes or channels instead
- **Multiple related variables**: Atomics can't guarantee consistency across variables
- **High-level synchronization**: Use channels for communication-based concurrency
- **When code clarity matters more than performance**: Atomics can be harder to reason about

## Exercise Ideas

1. Implement a thread-safe counter using atomics
2. Create a wait-free data structure (e.g., a concurrent stack)
3. Build a lightweight rate limiter using atomic operations
4. Implement a simple spin lock using CompareAndSwap

## Best Practices

- Use atomic operations for individual variables, not for coordinating across variables
- Prefer channels or mutexes when operations need to affect multiple variables atomically
- Remember that atomic operations only guarantee atomicity, not ordering
- Use memory barriers if operation ordering is important
- Consider using the higher-level sync types unless you need the performance

## Common Mistakes

- Mixing atomic and non-atomic operations on the same variable
- Assuming atomics enforce operation ordering (they don't without memory barriers)
- Using atomic operations for related variables that need consistent updates
- Not understanding the memory model implications
- Overusing atomics when mutexes would be clearer

## Next Steps

After understanding atomic operations, explore the Once type for one-time initialization in concurrent programs. 