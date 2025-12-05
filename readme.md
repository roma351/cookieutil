# CookieUtil

A fork of Go's standard `net/http/cookiejar` package with added JSON serialization support.

## What's Different?

This library is **identical to the original `cookiejar`** with one addition:

- ✅ **Added `serialize.go`** - implements JSON import/export functionality

The original `jar.go` remains **completely unchanged** from the standard library.

## Features

- ✅ Full JSON serialization support via `MarshalJSON()` / `UnmarshalJSON()`
- ✅ Session cookie filtering (excluded by default on import)
- ✅ Compatible with standard `json.Marshal()` / `json.Unmarshal()`
- ✅ Thread-safe operations
- ✅ 100% compatible with original cookiejar API

## Installation

```bash
go get github.com/roma351/cookieutil
```

## Usage

### Export cookies to JSON

```go
import "github.com/roma351/cookieutil"

jar, _ := cookieutil.NewWithPublicSuffix()

// Add some cookies...

// Export
data, err := json.Marshal(jar)
// or
data, err := jar.MarshalJSON()
```

### Import cookies from JSON

```go
// Option 1: Create new jar from JSON
jar, err := cookieutil.NewFromJSON(jsonString)

// Option 2: Import into existing jar
jar, _ := cookieutil.NewWithPublicSuffix()
err := json.Unmarshal(data, jar)
```

### Control session cookies import

By default, session cookies (`Persistent == false`) are **not imported**. To include them:

```json
{
  "nextSeqNum": 5,
  "entries": {...},
  "importSessionCookies": true
}
```

## JSON Format

```json
{
  "nextSeqNum": 5,
  "entries": {
    "example.com": {
      "cookie-key": {
        "Name": "session",
        "Value": "abc123",
        "Domain": "example.com",
        "Path": "/",
        "Expires": "2025-12-31T23:59:59Z",
        "Persistent": true,
        "Secure": true,
        "HttpOnly": true
      }
    }
  }
}
```
