# Lock-Free, Wait-Free, DataRace-Free Account Transactions Using Goroutines Only

## Purpose

This project is an exercise in implementing concurrent bank account operations **without using any synchronization primitives** such as mutexes, atomic operations, channels, or wait groups. The goal was to explore how far one can go using **only goroutines** to achieve a **wait-free**, **lock-free**, and **data race-free** design.

The implementation focuses on:

- Avoiding any shared mutable state by using **immutable transaction snapshots**.
- Running all account operations asynchronously via goroutines.
- Communicating results exclusively through **callback functions**, eliminating blocking or locking.
- Ensuring safety from race conditions and deadlocks without any synchronization tools or imports.

This design challenges traditional concurrency approaches by demonstrating that **pure goroutine-based asynchronous execution**, combined with immutability and callback patterns, can achieve safe parallelism even in sensitive scenarios like financial transactions.

---

## How It Works

- Each `Transaction` represents an immutable snapshot of the account balance.
- `Deposit` and `Withdraw` operations create new transaction snapshots rather than modifying existing state.
- Account operations run inside their own goroutines, making execution fully asynchronous.
- Results (success flags and updated account states) are returned via callback functions, enabling external state management.
- The `Transfer` function coordinates asynchronous withdrawal and deposit, ensuring atomicity through callback chaining.
- No shared mutable data is accessed concurrently, so the program avoids all data races inherently.

---

## Benchmark Results

The following benchmarks were performed on a Mac Mini with an Apple M4 chipset and 24GB RAM, highlighting high throughput and low overhead:

| Benchmark            | Ops/sec   | ns/op   | Allocations/op | Bytes/op |
|----------------------|-----------|---------|---------------|----------|
| Deposit              | 4,204,900 | 287.1   | 2             | 48       |
| Withdraw             | 4,168,310 | 287.1   | 2             | 48       |
| Total (combined ops) | 4,085,821 | 292.3   | 2             | 40       |
| Parallel Mixed Ops    | 6,951,534 | 170.8   | 2             | 48       |

These results demonstrate the system's efficiency in handling millions of lock-free, concurrent transactions per second with minimal memory allocations.

---

## Summary

This exercise showcases an unconventional approach to concurrency in Go by relying exclusively on goroutines and immutable data structures, eschewing traditional synchronization primitives entirely. It highlights how asynchronous execution and callbacks can maintain data integrity and prevent race conditions, offering a lock-free and wait-free model suitable for learning and experimentation in concurrent programming.

---

## Code Location

All implementation details can be found in the `main.go` file of this repository.