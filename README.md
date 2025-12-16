# go-webkit

**Opinionated toolkit for building Go Web APIs, fast and consciously.**

`go-webkit` is a small, composable toolkit for building Go Web APIs with a clear structure and minimal ceremony.

It is designed for developers who want to:

-   understand Go’s HTTP stack deeply
-   avoid framework lock-in
-   start API projects _fast_, with a consistent foundation

This is **not** a “one-size-fits-all” framework.  
It is an opinionated starting point.

---

## Philosophy

### 1. Core over framework

The core of `go-webkit` is built on `net/http`.

Frameworks (Gin, Echo, Chi, etc.) are supported via **thin adapters**, never baked into the core.  
You should be able to understand what happens on the wire.

### 2. Explicit over magic

-   Errors are explicit
-   HTTP status codes are explicit
-   JSON writing is explicit

Nothing happens “automatically behind the scenes”.

### 3. Opinionated, but escapable

`go-webkit` has opinions — intentionally.

If you disagree, you should be able to:

-   replace a part
-   bypass it
-   or stop using it entirely

No global state. No hidden hooks.

---

## What this library is (and is not)

### ✅ This library is

-   A **toolkit**, not a framework
-   A **foundation** for API projects
-   A place to encode _your_ defaults once and reuse them

### ❌ This library is not

-   A full-stack framework
-   A router replacement
-   A dependency injector
-   A silver bullet

---

## Packages overview

| Package        | Description                                             |
| -------------- | ------------------------------------------------------- |
| `errors`       | RFC7807-style Problem Details and error mapping         |
| `respond`      | Unified JSON / error HTTP responses                     |
| `middleware`   | HTTP middleware (request-id, recover, access log, etc.) |
| `adapters/gin` | Thin adapter for Gin                                    |
| `examples/gin` | Reference implementation                                |

---

## Quick Start (Gin)

This example demonstrates the **recommended way** to build a Go API using `go-webkit`.

```bash
go run ./examples/gin
```

Then call the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{ "status": "ok" }
```

---

## Minimal usage example

```go
r := gin.New()

r.Use(ginadapter.Use( // go-webkit gin adapter
    middleware.RequestID(), // go-webkit middleware
    middleware.Recover(),  // go-webkit middleware
))

r.GET("/users/:id", func(c *gin.Context) {
    if c.Param("id") == "0" {
        respond.Error(c.Writer, errors.BadRequest("invalid id")) // go-webkit error response
        return
    }
    respond.JSON(c.Writer, 200, gin.H{"id": c.Param("id")}) // go-webkit JSON response
})
```

---

## Design goals

-   Small surface area
-   Predictable behavior
-   Easy to remove or replace
-   Friendly to reading the source

If reading the code feels harder than using it, that’s a bug.

---

## Versioning

go-webkit follows semantic versioning.

-   v0.x: APIs may change
-   v1.0: API stability guaranteed
-   Breaking changes will only happen in major versions

---

## Who is this for?

This library is for developers who:

-   build many Go API projects
-   want a reusable starting point
-   care about understanding Go deeply

If you prefer maximal abstraction or batteries-included frameworks, this library may not be for you — and that’s okay.

---

## License

MIT
