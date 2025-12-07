# CookieUtil

A minimal fork of Go's `net/http/cookiejar` with JSON serialization and change tracking.

## Features

- 💾 **JSON Export/Import** - Save and restore cookies
- 🔔 **Change Tracking** - Get notified when cookies update
- 🎯 **Minimal Changes** - Original `jar.go` barely modified
- 🔒 **Thread-Safe** - All operations are concurrent-safe

## Installation
```bash
go get github.com/roma351/cookieutil
```

## Quick Start

### Save/Load Cookies
```go
import "github.com/roma351/cookieutil"

// Create jar
jar, _ := cookieutil.NewWithPublicSuffix()

// Save to JSON
data, _ := json.Marshal(jar)
os.WriteFile("cookies.json", data, 0644)

// Load from JSON
jar, _ = cookieutil.NewFromJSON(string(data))
```

### Auto-Save on Changes
```go
jar, _ := cookieutil.NewWithPublicSuffix()

// Save cookies automatically after changes
jar.OnCookieChange(func(j *cookieutil.Jar) error {
    data, _ := j.MarshalJSON()
    return os.WriteFile("cookies.json", data, 0644)
}, 5*time.Second) // Wait 5s after last change

// Use normally - cookies auto-save
client := &http.Client{Jar: jar}
client.Get("https://example.com")
```

**Debouncing:** Multiple changes within 5 seconds = one save at the end.

## API

### `NewWithPublicSuffix() (*Jar, error)`
Creates jar with public suffix list support.

### `NewFromJSON(jsonStr string) (*Jar, error)`
Creates jar from JSON string.

### `OnCookieChange(callback func(*Jar) error, debounce time.Duration)`
Triggers callback when cookies change. Debounce delays execution until changes stop.

### `MarshalJSON() / UnmarshalJSON()`
Standard JSON serialization support.

## JSON Format
```json
{
  "nextSeqNum": 5,
  "entries": {
    "example.com": {
      "session": {
        "Name": "session",
        "Value": "abc123",
        "Domain": "example.com",
        "Path": "/",
        "Expires": "2025-12-31T23:59:59Z",
        "Persistent": true
      }
    }
  }
}
```

**Note:** Session cookies (non-persistent) are excluded by default. Add `"importSessionCookies": true` to include them.

## What Changed?

- Added `serialize.go` (new file)
- Modified `jar.go`: +1 struct field, +1 method call

Original cookiejar logic untouched.

## License

BSD-style (same as Go standard library)