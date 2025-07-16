# 01 – Goroutines Basics

This section covers the fundamentals of using goroutines in Go for concurrent programming.

---

## What You’ll Learn

- How to start goroutines using the `go` keyword.
- Why `time.Sleep` is sometimes used temporarily.
- Correctly passing arguments to goroutines (avoiding closure pitfalls).
- Using `sync.WaitGroup` to wait for multiple goroutines.

---

## Files in This Folder

| File                     | Description                                      |
|-------------------------|--------------------------------------------------|
| `basic_goroutine.go`    | Starting a single goroutine with simple output.  |
| `waitgroup_goroutine.go`| Waiting for multiple goroutines using `sync.WaitGroup`. |
| `mini-challenge-1.go`  | Spawns 10 goroutines that each print their number using sync.WaitGroup|

---

## Key Concepts Summary

| Concept        | Explanation                                       |
|----------------|---------------------------------------------------|
| Goroutine      | Lightweight thread managed by Go runtime.         |
| Closure Trap   | Loop variables must be passed carefully to goroutines. |
| WaitGroup      | Synchronization primitive for waiting on goroutines to finish. |

---

## How to Run

Run each example file using:

```bash
go run basic_goroutine.go
go run waitgroup_goroutine.go
go run passing_arguments.go
```
