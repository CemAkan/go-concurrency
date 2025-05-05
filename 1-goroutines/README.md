# Goroutines

Goroutines are the fundamental building block of concurrency in Go. They are lightweight threads managed by the Go runtime.

## Key Concepts

- **Lightweight**: Goroutines have a small memory footprint (starting at around 2KB of stack size)
- **Multiplexing**: Many goroutines are multiplexed onto a smaller number of OS threads
- **Simple syntax**: Just add the `go` keyword before a function call

## How Goroutines Work

- Goroutines run concurrently with other goroutines in the same address space
- They start with a small stack that can grow and shrink as needed
- The Go scheduler handles the distribution of goroutines across available CPU cores

## Common Patterns

- Starting multiple goroutines to perform concurrent tasks
- Understanding that the main function runs in its own goroutine
- Knowing that a program exits when the main goroutine completes, regardless of other goroutines

## Challenges

- Race conditions when accessing shared data
- Program termination before goroutines complete
- Lack of communication between goroutines

## Exercise Ideas

1. Create a program that starts multiple goroutines to print numbers
2. Observe the interleaving of output from different goroutines
3. Explore what happens when the main function completes before goroutines finish

## Best Practices

- Use synchronization mechanisms to coordinate goroutines
- Be cautious about spawning too many goroutines
- Consider using worker pools for controlled parallelism

## Common Mistakes

- Forgetting that goroutines need coordination to work properly
- Not handling potential race conditions
- Assuming goroutine execution order

## Next Steps

After mastering goroutines, learn about channels to enable communication between goroutines. 