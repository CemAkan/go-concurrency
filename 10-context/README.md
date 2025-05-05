# Context Package

The Context package provides a standardized way to carry deadlines, cancellation signals, and request-scoped values across API boundaries and between processes.

## Key Concepts

- **Cancellation propagation**: Efficiently communicates cancellation through call chains
- **Deadline management**: Associates deadlines with operations
- **Value propagation**: Carries request-scoped values across API boundaries
- **Thread safety**: Safe for concurrent use by multiple goroutines

## Core Context Types

- **context.Background()**: Root context, never cancelled
- **context.TODO()**: Placeholder when you're not sure which context to use
- **context.WithCancel()**: Context with explicit cancellation
- **context.WithDeadline()**: Context that cancels at a specific time
- **context.WithTimeout()**: Context that cancels after a duration
- **context.WithValue()**: Context carrying a key-value pair

## Basic Usage

```go
func main() {
    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Always call cancel to avoid leaks
    
    // Pass the context to a function
    result, err := doSomething(ctx)
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    fmt.Println("Result:", result)
}

func doSomething(ctx context.Context) (string, error) {
    // Periodically check if context is done
    select {
    case <-ctx.Done():
        return "", ctx.Err() // ctx.Err() returns the cancellation reason
    case <-time.After(1 * time.Second):
        return "Operation completed", nil
    }
}
```

## Context in HTTP Servers

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Get the request context
    ctx := r.Context()
    
    // Use the context for downstream operations
    result, err := doOperationWithContext(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    fmt.Fprintf(w, "Result: %s", result)
}
```

## Context Values

```go
type key int
const userIDKey key = 0

// Add a value to context
ctx := context.WithValue(parentCtx, userIDKey, "user-123")

// Retrieve value from context
if userID, ok := ctx.Value(userIDKey).(string); ok {
    fmt.Println("User ID:", userID)
}
```

## Best Practices

- Always pass context as the first parameter to functions
- Always call cancel when creating cancellable contexts
- Don't store Contexts in structs; pass them explicitly
- Use context Values only for request-scoped data, not for optional parameters
- Use custom types for context keys to avoid key collisions
- Don't use context.TODO() in production code

## Common Mistakes

- Forgetting to call cancel functions
- Using background context when a more specific context is available
- Storing sensitive information in context Values
- Using basic types like strings as context keys
- Creating deeply nested context chains

## Exercise Ideas

1. Implement a service with timeout-based operations
2. Create an HTTP middleware that adds values to context
3. Build a client that respects server-side cancellation
4. Design a pipeline that propagates cancellation through stages

## Next Steps

After mastering the Context package, explore time management patterns for controlling the timing of concurrent operations. 