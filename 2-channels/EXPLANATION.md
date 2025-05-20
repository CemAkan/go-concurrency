# Bomb Passing Game Using Unbuffered Channels

## Overview

This project is a simple and educational concurrency demonstration inspired by the "ping-pong" example presented in **Google I/O 2013 - Advanced Go Concurrency Patterns**. It models a **3-player bomb passing game** where a bomb is passed asynchronously between players until it explodes.

The core purpose is to illustrate the behavior of **unbuffered channels in Go**, specifically how they **block the sender until the receiver is ready**, enforcing synchronization between goroutines.

---

## Key Concepts Demonstrated

- **Unbuffered Channels**: The game uses unbuffered channels to pass a `Bomb` struct between players. Each send operation blocks until the corresponding receive happens.
- **Goroutine Synchronization**: Each player is a goroutine waiting to receive the bomb, hold it for a random duration, and then pass it along.
- **Random Delays**: Players hold the bomb for 1 to 3 seconds, decreasing the bomb's timer.
- **Termination Condition**: The bomb explodes when its timer reaches zero or less, ending the game.
- **Circular Passing**: The bomb continuously cycles among three players until it explodes, demonstrating cyclic concurrency patterns.

---

## How It Works

- Three unbuffered channels (`a`, `b`, `c`) connect three players: Cem → Mete → Melis → Cem.
- Each player goroutine waits to receive the bomb from its input channel.
- Upon receiving, the player:
    - Sleeps for a random time between 1 and 3 seconds.
    - Decreases the bomb's remaining time accordingly.
    - Updates the bomb holder's name.
    - Prints who is holding the bomb and for how long.
- If the bomb timer reaches zero or below, the bomb explodes, printing a game over message and terminating that player's goroutine.
- Otherwise, the bomb is passed to the next player through the output channel.
- The main function seeds the random number generator, initializes channels, starts the players as goroutines, and sends the initial bomb to start the game.
- Finally, the main thread sleeps enough time to allow the game to complete.

---

## Notes

- The game logic is deterministic except for random hold durations, making each run slightly different.
- The use of unbuffered channels ensures that no player can pass the bomb until the next player is ready to receive it.
- This synchronization pattern can be extended to more complex concurrent workflows where strict handoff and sequencing are required.

---

## Code Location

The full implementation can be found in `main.go` within the relevant project directory.

---

## Summary

This bomb passing game is a fun, practical demonstration of Go's unbuffered channel semantics and goroutine synchronization. It effectively teaches the blocking behavior of unbuffered channels in a memorable, interactive way.