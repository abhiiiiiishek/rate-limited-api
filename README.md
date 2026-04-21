# Rate Limited API Service

## Overview

This project implements a simple rate-limited API service in Go.

It exposes two endpoints:
POST /request
GET /stats

The service enforces a limit of 5 requests per user per minute.

## Run

go run cmd/server/main.go

## Test

curl -X POST http://localhost:8080/request \
-H "Content-Type: application/json" \
-d '{"user_id":"user1","payload":"test"}'

curl http://localhost:8080/stats
