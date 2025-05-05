# Rate Limiting

Rate limiting is a technique used to control the amount of requests or operations that can be performed within a specific timeframe, preventing resource exhaustion and ensuring fair usage.

## Key Concepts

- **Throughput control**: Limiting the number of operations per time unit
- **Resource protection**: Preventing overload of downstream services or resources
- **Fairness**: Ensuring no single client consumes all available resources
- **Backpressure**: Signaling to producers when they're generating too much load

## Common Rate Limiting Algorithms

- **Token Bucket**: Tokens are added to a bucket at a fixed rate; operations consume tokens
- **Leaky Bucket**: Operations are processed at a fixed rate, with a queue for waiting operations
- **Fixed Window**: Count operations in fixed time windows (e.g., per minute)
- **Sliding Window**: Count operations over a rolling time period
- **Sliding Window Log**: Track timestamps of operations and count those within the time window

## Basic Time-Based Rate Limiter

```go
func basic() {
    // Allow one operation every 200ms (5 per second)
    limiter := time.Tick(200 * time.Millisecond)
    
    for request := range requests {
        <-limiter // Wait for a token
        go process(request)
    }
}
```

## Bursty Rate Limiter

```go
func bursty() {
    // Allow bursts of up to 3 operations, then 1 per 200ms
    burstyLimiter := make(chan time.Time, 3)
    
    // Fill the burst capacity
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }
    
    // Refill at a rate of 5 per second
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()
    
    for request := range requests {
        <-burstyLimiter // Wait for a token
        go process(request)
    }
}
```

## Using golang.org/x/time/rate Package

```go
import "golang.org/x/time/rate"

func main() {
    // Create a limiter with rate of 10 per second and burst of 3
    limiter := rate.NewLimiter(10, 3)
    
    // Blocking approach
    for request := range requests {
        limiter.Wait(context.Background())
        go process(request)
    }
    
    // Non-blocking approach
    for request := range requests {
        if limiter.Allow() {
            go process(request)
        } else {
            // Handle rate limiting (e.g., return 429 Too Many Requests)
            rejectRequest(request)
        }
    }
}
```

## Distributed Rate Limiting

- Rate limiting across multiple instances requires coordination
- Common approaches include:
  - Redis-based rate limiters
  - Centralized token bucket services
  - Consistent hashing for sharded rate limiting

## Exercise Ideas

1. Implement a token bucket rate limiter
2. Create a rate limiter for a REST API
3. Build a service with different rate limits for different clients
4. Implement a distributed rate limiter using Redis

## Best Practices

- Match rate limits to available resources and expected load
- Communicate rate limit information to clients (headers, documentation)
- Consider separate rate limits for different operations or users
- Implement graceful degradation when rate limits are exceeded
- Log and alert on sustained high rate limit hits

## Common Mistakes

- Rate limiting at the wrong level (too coarse or too fine-grained)
- Not considering burst requirements
- Not adapting rate limits to changing conditions
- Implementing rate limiters that aren't thread-safe
- Not communicating rate limits to clients

## Next Steps

After understanding rate limiting, explore semaphores for controlling access to limited resources. 