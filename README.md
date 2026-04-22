# Rate Limited API Service

## Overview

This project implements a simple rate-limited API service in Go.

It exposes two endpoints:

- POST /request
- GET /stats

The service enforces a limit of 5 requests per user per minute.
Concurrent requests were tested using parallel curl calls to ensure correctness under simultaneous access.

## Design Decisions

The limiter uses a sliding window approach to avoid burst allowance at window boundaries, which is common in fixed window implementations.

Mutex-based synchronization was chosen for simplicity and correctness under concurrent access.

An in-memory store was used to keep the implementation simple and focused on core logic.

---

## How it works

Each user is associated with a list of request timestamps.

For every incoming request:

1. Timestamps older than the 1 minute window are removed
2. If the number of remaining requests is greater than or equal to the limit, the request is rejected
3. Otherwise, the request is accepted and the current timestamp is recorded

This ensures that rate limiting remains accurate even under concurrent requests.

---

## Steps to run

Prerequisites:
- Go 1.22 or above

Run the server:

```bash
go run cmd/server/main.go
```
The server will start on:
http://localhost:8080

## Testing the APIs

Send request
```bash
curl -X POST http://localhost:8080/request \
-H "Content-Type: application/json" \
-d '{"user_id":"user1","payload":"test"}'
```

Get stats
```bash
curl http://localhost:8080/stats
```

Concurrency test
```bash
for i in {1..10}; do 
  curl -X POST http://localhost:8080/request \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user1","payload":"x"}' &
done
wait
```

## Limitations

* Data is stored in memory and is not persisted
* Not suitable for distributed or multi-instance environments
* Memory usage can grow with number of users
* No cleanup mechanism for inactive users


## Improvements with more time

* Replace in-memory store with Redis for distributed rate limiting
* Implement token bucket or leaky bucket algorithm for better control
* Add automatic cleanup of inactive users to manage memory
* Expose rate limit headers in API responses
* Add structured logging and metrics for observability
* Add unit and load tests
* Support graceful shutdown and request draining
