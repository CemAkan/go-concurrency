# Channels

Channels are the primary mechanism for communication between goroutines in Go. They provide a way for goroutines to synchronize execution and exchange data safely.

## Key Concepts

- **Communication**: Channels enable goroutines to send and receive values
- **Synchronization**: Channel operations block until the other side is ready
- **Typed**: Each channel can only transport values of a specific type

## Types of Channels

- **Unbuffered channels**: Synchronous, blocking until both sender and receiver are ready
- **Buffered channels**: Asynchronous up to the buffer size
- **Closed channels**: Can no longer accept data but can still be read from

## Common Operations

- **Creation**: `ch := make(chan Type, [capacity])`
- **Sending**: `ch <- value`
- **Receiving**: `value := <-ch` or `value, ok := <-ch`
- **Closing**: `close(ch)`
- **Range**: `for value := range ch { ... }`

## Common Patterns

- **Pipeline**: Connect stages of processing with channels
- **Fan-out/Fan-in**: Distribute work and collect results
- **Signaling**: Use channels as semaphores to coordinate goroutine execution
- **Done channels**: Signal completion or cancellation

## Challenges

- Deadlocks when channels are blocked
- Panics when sending on a closed channel
- Understanding channel directionality

## Exercise Ideas

1. Create a basic producer-consumer pattern
2. Implement a timeout using channels and select
3. Build a pipeline processing data through multiple stages

## Best Practices

- Use channel direction constraints in function signatures (`chan<-` for send-only, `<-chan` for receive-only)
- Close channels from the sender side, never from the receiver
- Use `range` to receive values until a channel is closed
- Prefer unbuffered channels for synchronization guarantees

## Common Mistakes

- Forgetting to close channels
- Writing to a closed channel (causes panic)
- Creating deadlocks with improper channel usage
- Not handling channel emptiness check with the comma-ok idiom

## Next Steps

After understanding channels, explore WaitGroups for goroutine synchronization when communication isn't needed. 