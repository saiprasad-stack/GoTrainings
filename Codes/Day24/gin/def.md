--
## Router

A router is the component that receives a request and decides which function should handle it.

## Routing

Routing is the process of mapping a URL and HTTP method to a handler function.

## Handler

A handler is the function that processes the request and returns a response.

## Endpoint

An endpoint is a specific URL where an API can be accessed.

Example: `/users`, `/login`

---

# Request Data Terms

## Params (Path Parameters)

Params are values inside the URL path.

Example: `/user/101` → 101 is a param

## Query Parameters

Query params are key-value pairs added after `?` in a URL.

Example: `/search?name=John`

## Request Body

The body is the data sent by the client (usually JSON).

```json
{
  "name": "John"
}
```

## Headers

Headers contain metadata about the request.

Examples:

* Authorization token
* Content-Type

---

# Middleware & Flow

## Middleware

A middleware is a function that runs before or after the main handler.

Used for:

* Authentication
* Logging
* Validation

## Context (*gin.Context)

The context holds everything about a request:

* Input (params, query, body)
* Output (response)

## Next()

Tells Gin to continue to the next middleware or handler.

## Abort()

Stops the request and prevents further processing.

---

# Response Terms

## Response

The response is what the server sends back to the client.

## JSON

A data format used to send structured data.

Example:

```json
{"name": "John"}
```

## Status Code

A number indicating the result of a request.

Common codes:

* 200 → success
* 400 → bad request
* 401 → unauthorized
* 500 → server error

---

# Server Concepts

## Server

A program that listens for requests and sends responses.

## Port

A number used to access a server.

Example: `localhost:8080`

## API (Application Programming Interface)

A set of endpoints that allow systems to communicate.

## REST API

A type of API that uses HTTP methods like GET, POST, etc.

---

# Advanced but Important

## Binding

Converting request data into a Go struct.

## Serialization

Converting a Go struct into JSON.

## Deserialization

Converting JSON into a Go struct.

## Latency

Time taken to process a request.

## Throughput

Number of requests handled per second.

---
