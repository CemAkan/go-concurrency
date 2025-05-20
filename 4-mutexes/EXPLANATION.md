# Project Explanation: Mutex-Protected Click Counter with Environment-Based Configuration in a Minimal URL Shortener

## Overview

This project implements a **minimal URL shortener** service in Go, focusing on:

- **Thread-safe in-memory click counting** using **mutex** for concurrency safety,
- **PostgreSQL-backed URL storage** with GORM ORM,
- **Fiber web framework** for lightweight and performant HTTP API,
- **Configuration management via environment variables** to enable flexible deployment across environments.

The primary goal is to provide a clean, efficient, and concurrency-safe solution for counting clicks on shortened URLs without incurring database overhead on every click, while keeping the system simple and maintainable.

---

## Key Components

### 1. In-Memory Click Counter

- Stored as a Go map of `map[string]int64` mapping short codes to their click counts.
- Protected by a `sync.Mutex` to avoid data races in concurrent increment and read operations.
- Provides two core methods:
    - `Increment(code string)` safely increments the click count for a given short code.
    - `Get(code string)` safely retrieves the current count for a short code.
- Designed to reduce database writes by aggregating clicks in-memory, improving performance under high load.

### 2. Environment-Based Configuration

- Uses environment variables (`DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT`, `APP_PORT`) for all sensitive and environment-dependent settings.
- Enables easy deployment and configuration across local, staging, and production environments without code changes.
- Configuration is loaded via a dedicated `config` package (e.g., using `github.com/caarlos0/env`).

### 3. PostgreSQL and GORM Integration

- URL records (`OriginalURL`, `ShortCode`, `ClickCount`) are stored persistently in PostgreSQL.
- GORM ORM provides seamless mapping and migration support.
- Database schema migrations are handled automatically on application start.
- Click count in the DB is updated asynchronously or periodically from the in-memory counter, minimizing write contention.

### 4. HTTP API with Fiber

- Endpoints:
    - `POST /shorten`: Accepts a long URL, generates a unique short code, and stores the mapping.
    - `GET /:code`: Redirects to the original URL and increments the click count both in-memory and the DB.
- Fiber framework chosen for performance and developer ergonomics.

---

## Concurrency and Safety Considerations

- **Why Mutex?**  
  Go maps are not safe for concurrent write/read access. Using a mutex ensures no race conditions when multiple goroutines increment click counts simultaneously.

- **Avoiding DB Overhead**  
  Writing to the database on every click can cause bottlenecks. The in-memory counter aggregates clicks safely and can batch updates to the DB asynchronously (not implemented here but recommended).

- **Clean Shutdown & Resource Management**  
  The counter exposes `Stop()` to allow graceful shutdown of any cleanup goroutines, preventing memory leaks and ensuring data consistency.

---

## Environment Variables Example

```env
DB_HOST=localhost
DB_USER=cemakan
DB_PASSWORD=123456789cem
DB_NAME=goroutines
DB_PORT=5432
APP_PORT=383