# 🚀 EasyCache – Lightweight In-Memory Cache for Go

**EasyCache** is a simple, efficient, and thread-safe in-memory cache library for Go. It supports both **synchronous** and **asynchronous** caching with automatic expiration and cleanup.

## ✨ Features

- 🔥 **Fast in-memory caching** for quick lookups
- ⏳ **Automatic expiration** with configurable TTL
- ⚡ **Sync & Async modes** for different performance needs
- 🧹 **Auto cleanup** of expired items
- 🎯 **Generic support** to store any data type

---

## 📦 Installation

`go get github.com/mmdali-dev/easycache`

---

## 🛠 Usage

### **Synchronous Cache**

```go
package main

import (
    "fmt"
    "time"
    "github.com/mmdali-dev/easycache/sync"
)

func main() {
    cache := sync.NewSyncCache[string](time.Second * 10)

    cache.Set("user", "JohnDoe", time.Second*5)

    value, found := cache.GetWithCheck("user")
    if found {
        fmt.Println("User:", value) // Output: User: JohnDoe
    }
}

```

### **Asynchronous Cache**

```go
package main

import (
    "fmt"
    "time"
    "github.com/mmdali-dev/easycache/async"
)

func main() {
    cache := async.NewAsyncCache[string](time.Second * 10)

    cache.Set("session", "xyz123", time.Second*5)

    value, found := cache.GetWithCheck("session")
    if found {
        fmt.Println("Session:", value) // Output: Session: xyz123
    }
}

```

---

## 📜 API Reference

### **SyncCache Methods**

| Method                 | Description                              |
| ---------------------- | ---------------------------------------- |
| `Set(key, value, ttl)` | Stores a value with an optional TTL      |
| `GetWithCheck(key)`    | Retrieves value, checks expiration       |
| `GetWithoutCheck(key)` | Retrieves value without expiration check |
| `Delete(key)`          | Removes a specific key                   |
| `Clear()`              | Clears all cache entries                 |

### **AsyncCache Methods**

(Same as above but uses `sync.Map` for thread safety)

---

## 📂 Examples

For a more **detailed example**, check out the [`examples`](./examples) directory:

```go
cd examples/async_session

go run async_session.go
```

or

```go
cd examples/sync_session

go run sync_session.go
```

This example demonstrates how to **cache session tokens**

## 🔥 Why Use EasyCache?

- ✅ **Simple API** – Minimal setup, easy integration
- 🚀 **High Performance** – Optimized for speed
- 🔧 **Generic Support** – Store any data type
- 💡 **Use Cases** – Session caching, API rate limiting, temporary data storage

---

## 📄 License

MIT License © 2025

---

### 🔗 Contribute

Want to improve **EasyCache**? PRs are welcome! 🎉

---

This README gives **a quick overview** without being too overwhelming. Let me know if you need refinements! 🚀
