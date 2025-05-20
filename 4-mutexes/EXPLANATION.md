# Mutex-Protected Click Counter in a Minimal URL Shortener — Guide & Explanation

## Overview

This project implements a minimal URL shortening service in Go that focuses on:

- **Thread-safe in-memory click counting** with `sync.RWMutex` to avoid data races,
- **PostgreSQL database storage** for URLs using GORM ORM,
- **HTTP API** built with Fiber web framework,
- **Environment-based configuration** for flexible deployment,
- **Clean and efficient concurrency management** for counting URL clicks without excessive DB writes.

---

## Components & Architecture

### 1. ClickCounter: Thread-Safe In-Memory Counter

- Uses a Go map `map[string]int64` to store click counts by short URL code.
- Protects map access using `sync.RWMutex` for safe concurrent increments and reads.
- Supports:
  - `Increment(code string)` — safely increments the click count,
  - `Get(code string)` — safely retrieves the current count.
- Maintains an internal cleanup goroutine that periodically trims the map to a max size (`maxEntries`) to avoid unbounded memory growth.
- Cleanup interval and max entries are configurable via environment variables (`CLICK_COUNTER_CLEANUP_INTERVAL`, `CLICK_COUNTER_MAX_ENTRIES`).

### 2. Database Layer with GORM and PostgreSQL

- Defines a `ShortURL` struct with fields:
  - `OriginalURL string` — the full URL,
  - `ShortCode string` — unique short code,
  - `ClickCount int64` — persisted click count.
- Uses GORM migrations to create and maintain the database schema.
- Generates random short codes for new URLs, checking uniqueness before creation.
- Retrieves URLs by short code on requests.

### 3. HTTP API with Fiber

- Provides endpoints:
  - `POST /shorten` — accepts JSON with a long URL, creates and returns a unique short code,
  - `GET /:code` — redirects to the original URL, increments click counters in-memory and asynchronously updates DB,
  - **`GET /stats/:code`** — returns JSON with statistics for a short code, including:
    - The original URL,
    - The total clicks persisted in the database,
    - The current in-memory clicks that are not yet persisted.

---

## Usage Guide

### Environment Variables Example

```env
DB_HOST=localhost
DB_USER=cemakan
DB_PASSWORD=123456789cem
DB_NAME=goroutines
DB_PORT=5432

APP_PORT=4000

CLICK_COUNTER_CLEANUP_INTERVAL=5m     # Duration format: "5m", "30s" (supports time.ParseDuration syntax)
CLICK_COUNTER_MAX_ENTRIES=10000        # Maximum entries in the in-memory click counter map
```