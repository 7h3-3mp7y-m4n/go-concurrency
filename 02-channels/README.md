# 02 – Channels in Go

This section covers Go channels in-depth. Channels allow goroutines to communicate safely without sharing memory.

---

## What You’ll Learn

* How unbuffered channels work.
* How buffered channels work.
* How to close channels and range over them.
* How to use `select` to wait on multiple channels.
* Using `select` with `default` and timeouts.

---

## Theory Overview

### What Are Channels?

Channels are Go's concurrency primitives used to communicate between goroutines. They provide a safe way to pass data without using shared memory.

### Types of Channels:

* **Unbuffered Channels:**

  * Data transfer happens when both sender and receiver are ready.
  * Useful for tightly synchronized communication.

* **Buffered Channels:**

  * Can hold a limited number of values.
  * Sender blocks only when the buffer is full.
  * Receiver blocks only when the buffer is empty.

### Closing Channels:

* Once all values have been sent, close the channel.
* Receivers detect closure using `for value := range ch`.

### Select Statement:

* Lets you wait on multiple channel operations.
* Picks whichever channel operation is ready first.
* Prevents blocking when multiple channels are involved.

### Select with Default and Timeout:

* Adding `default` allows non-blocking select.
* Using `time.After` helps in implementing timeouts when waiting for a channel value.

---

## Files in This Folder

| File                  | Description                                               |
| --------------------- | --------------------------------------------------------- |
| `channels.go`         | Demonstrates unbuffered and buffered channel basics.      |
| `closing-channels.go` | Shows how to close channels and range over them.          |
| `mini-challenge-1.go` | Practice exercise combining buffered/unbuffered channels. |
| `select.go`           | Basic `select` statement example.                         |
| `select-default.go`   | Using `select` with `default` and `time.After`.           |

---

## Key Concepts Summary

| Concept             | Explanation                                                        |
| ------------------- | ------------------------------------------------------------------ |
| Unbuffered Channel  | Blocks until both sender and receiver are ready.                   |
| Buffered Channel    | Allows sending multiple values up to buffer size without blocking. |
| Closing Channels    | Signals that no more values will be sent.                          |
| Range Over Channel  | Processes all values until channel is closed.                      |
| Select Statement    | Wait on multiple channels, whichever is ready first.               |
| Select with Default | Avoid blocking indefinitely or implement timeouts.                 |

---

## How to Run

Run each example file using:

```bash
go run channels.go
go run closing-channels.go
go run mini-challenge-1.go
go run select.go
go run select-default.go
```