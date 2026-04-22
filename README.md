# Rate Limited API Service

## Overview

This project implements a simple rate-limited API service in Go.

It exposes two endpoints:
POST /request
GET /stats

The service enforces a limit of 5 requests per user per minute.
Concurrent requests were tested using parallel curl calls to ensure correctness under simultaneous access.

## Design Decisions

The limiter uses a sliding window approach to avoid burst allowance at window boundaries, which is common in fixed window implementations.
Mutex-based synchronization was chosen for simplicity and correctness under concurrent access.

## Run

go run cmd/server/main.go

## Test

curl -X POST http://localhost:8080/request \
-H "Content-Type: application/json" \
-d '{"user_id":"user1","payload":"test"}'

curl http://localhost:8080/stats
