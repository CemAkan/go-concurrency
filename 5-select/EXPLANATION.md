# TCP-Based Local Network Terminal Bomb Game Explanation

## Overview

This project implements a **multiplayer bomb-passing game** playable over a **local network** using TCP connections in Go. The game runs entirely in the terminal and uses **random timers** with interactive controls to simulate a bomb being passed between players.

The core gameplay involves:

- Players passing a virtual bomb to each other over TCP.
- The bomb carries a countdown timer with a randomized duration.
- Players can **hold down the space bar** to reduce the bomb timer faster.
- When the timer reaches zero, the bomb explodes on the current holder, ending the game.
- The game leverages Go's `select` statement to concurrently handle TCP input/output and user keyboard input efficiently.

---

## Key Concepts

### 1. **TCP Networking for Multiplayer**

- Players connect to a host server using TCP sockets over the local network.
- The bomb state and player actions are communicated over these TCP connections.
- This architecture enables multiplayer gameplay beyond a single machine.

### 2. **Terminal-Based User Interaction**

- The game is played entirely in the terminal.
- Real-time user input (e.g., pressing and holding space) is captured without blocking the game loop.
- Terminal rendering shows the bomb timer, current holder, and game status.

### 3. **Randomized Bomb Timer**

- Each bomb round has a randomly assigned countdown timer.
- The timer decrements naturally over time but can be accelerated by player actions (holding space bar).

### 4. **Concurrent Event Handling via `select`**

- Go's `select` statement is used to listen to multiple event sources concurrently:
   - Incoming TCP messages from other players.
   - Local keyboard input events.
   - Timer ticks to update bomb countdown.
- This allows the game to react instantly to network and user events without blocking.

---

## How The Game Works

1. **Initialization**  
   The host starts a TCP server, and players connect as clients over the local network.

2. **Game Loop**  
   Inside each player's process, a goroutine runs the main game loop using a `select` statement to handle:
   - Receiving bomb state updates over TCP.
   - Processing local keyboard events to detect space bar presses.
   - Updating the bomb timer periodically via ticker channels.

3. **Bomb Passing**  
   Players pass the bomb by sending messages over TCP sockets, transferring bomb ownership.

4. **Timer Reduction**  
   When a player holds the space bar, the bomb timer reduces faster locally and is synchronized over TCP.

5. **Bomb Explosion**  
   When the timer reaches zero, a bomb explosion event is triggered on the current holder, ending the game with a message.

6. **Game Termination**  
   The game ends gracefully, closing all TCP connections.

---

## File Structure (Example)

- `main.go` — Entry point, game setup, and main loop.
- `host.go` — TCP server logic for hosting the game.
- `client.go` — TCP client logic to connect to the host.
- `bomb.go` — Bomb struct and timer handling.
- `play.go` — Game logic and event processing.
- `menu.go` — Terminal menus and input handling.
- `result.go` — Game results and termination messages.
- `start.go` — Initialization routines and environment setup.

---

## Summary

This TCP-connected terminal bomb game is a practical exploration of real-time multiplayer gaming over local networks in Go. It highlights the effective use of Go's concurrency model with `select`, terminal event handling, and TCP socket programming to create an engaging interactive experience.

---

## Notes

- Requires local network connectivity.
- Designed for terminal/console play.
- Can be extended with more players, richer UI, or network reliability features.