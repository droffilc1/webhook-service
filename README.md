# Webhook Delivery Service

## Overview

Webhook Delivery Service is a HTTP service that allows applications to publish 
events and register webhook endpoints. When an event is received, the service 
stores it in memory/database, retrieves all registered endpoints, and delivers
the event synchronously using HTTP POST requests. The project is built entirely
with Go's standard library and demonstrates REST API design, dependency injection,
and modular backend architecture.

---

## Features

* Register webhook endpoints
* Publish webhook events
* Synchronous event delivery
* In-memory storage
* REST API using the Go standard library
* Health check endpoint

---

## Technologies Used

- Go
- In-memory data storage

---

## Run Project Locally

### Clone the repository

```sh
git clone https://github.com/droffilc1/webhook-service.git
cd webhook-service
```

### Start the server

```sh
go run ./cmd/api/main.go
```

The API will be available at:

```text
http://localhost:4000
```

---

## Example Workflow

The webhook delivery process follows this flow:

```text
Register Endpoint
        │
        ▼
Publish Event
        │
        ▼
Store Event
        │
        ▼
Retrieve Registered Endpoints
        │
        ▼
Deliver Event via HTTP POST
        │
        ▼
Record Delivery Attempt
```

Typical usage:

1. Register one or more webhook endpoints.
2. Publish an event.
3. The service delivers the event to every registered endpoint.
4. Inspect stored events and endpoints through the API.

---

# API Endpoints

## Health

### Check server status

```http
GET /health
```

Example response:
```json
{
    "status":"ok"
}
```

---

## Events

### Create an event

```sh
curl -X POST http://localhost:4000/events \
  -H "Content-Type: application/json" \
  -d '{
    "type": "user.created",
    "payload": {
      "email": "test@example.com"
    }
  }'
```

Example response:

```json
{
  "id": "32ed1be5-dad3-41a7-a1af-32ca0204f8f2",
  "type": "user.created",
  "payload": {
    "email": "test@example.com"
  },
  "created_at": "2026-08-03T17:12:41.389805212+03:00"
}
```

### List all events

```sh
curl http://localhost:4000/events
```

### Get an event

```sh
curl http://localhost:4000/events/<event-id>
```

---

## Endpoints

### Register an endpoint

```sh
curl -X POST http://localhost:4000/endpoints \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://webhook.site/your-unique-url"
  }'
```

Example response:

```json
{
  "id": "2d7b3fea-806b-41f9-a264-85cdf1d9f9f1",
  "url": "https://webhook.site/your-unique-url",
  "created_at": "2026-08-03T17:31:18.093196989+03:00"
}
```

### List all endpoints

```sh
curl http://localhost:4000/endpoints
```

Example response:

```json
[
  {
    "id": "2d7b3fea-806b-41f9-a264-85cdf1d9f9f1",
    "url": "https://webhook.site/your-unique-url",
    "created_at": "2026-08-03T17:31:18.093196989+03:00"
  }
]
```

### Get an endpoint

```sh
curl http://localhost:4000/endpoints/<endpoint-id>
```

### Update an endpoint

```sh
curl -X PUT http://localhost:4000/endpoints/<endpoint-id> \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/webhook"
  }'
```

### Delete an endpoint

```sh
curl -X DELETE http://localhost:4000/endpoints/<endpoint-id>
```

---

## Project Structure

```text
cmd/
    api/

internal/
    delivery/
    handler/
    model/
    store/
```

---

## Progress

* [x] In-memory storage
* [x] Synchronous event delivery
* [x] REST API endpoints
* [x] PostgreSQL-backed storage (`postgres.go`) implementing the `Store` interface
* [ ] Authentication middleware with API key support for publishers
* [ ] Redis-backed idempotency keys to prevent duplicate event processing
* [ ] Asynchronous event delivery using Asynq with retries, exponential backoff, and a dead-letter queue
