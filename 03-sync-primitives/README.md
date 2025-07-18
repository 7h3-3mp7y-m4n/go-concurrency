# 03-Sync-Primitives: Mutex, RWMutex, Once, Atomic Operations

This module focuses on Go’s synchronization primitives essential for controlling shared data across goroutines in concurrent programs.

---

## Concepts Covered

| Primitive         | Purpose                                                            |
|-------------------|--------------------------------------------------------------------|
| `sync.Mutex`      | Mutual exclusion for protecting shared variables from race conditions. |
| `sync.RWMutex`    | Optimized lock allowing multiple readers or a single writer at a time. |
| `sync.Once`       | Ensures one-time initialization across goroutines (e.g., lazy init). |
| `sync/atomic`     | Lock-free atomic operations for counters, flags, and simple shared state. |

---

## Why Use Sync Primitives?

While Go encourages using channels for concurrency, there are situations where you need direct control over shared memory.

**Common Use Cases:**
- Counting concurrent events (e.g., active requests)
- Shared state like in-memory caches
- Initialization logic that should only execute once (e.g., loading configuration)
- High-performance counters without locks (using `atomic`)

---

## How to Run the Examples

Every `.go` file in this directory demonstrates one synchronization primitive in practice.

You should always run with Go’s built-in race detector:

```bash
go run -race mutex_example.go
go run -race rwmutex_example.go
go run -race once_example.go
go run -race atomic_counter.go
```
## Summary 

### sync.Mutex
- Provides exclusive locking.  
- Only one goroutine can hold the lock at a time.  
- Best for read/write access where reads and writes are equally frequent.  

### sync.RWMutex
- Allows multiple goroutines to read shared data simultaneously.  
- Only blocks readers when a writer holds the lock.  
- Useful for caches and shared lookup tables.  

### sync.Once
- Ensures a function is executed only once.  
- No matter how many goroutines call `Do()`, only the first will run the function.  
- Handy for initializing expensive resources.  

### sync/atomic
- Provides low-level, high-performance operations for numeric counters or flags.  
- No mutex required but limited to simple data types like integers.  
- Faster than locks but only use for simple state, not complex objects.  
