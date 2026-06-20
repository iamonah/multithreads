# Go Multithreading

This repository contains Go programs centered on multithreading and concurrency.

The code focuses on practical use of Go's core concurrency primitives, including:

- goroutines
- `sync.WaitGroup`
- `sync.Mutex`
- `sync.RWMutex`
- shared state coordination

## Projects

### `boidsumulation`

A boids simulation built with Ebiten where many boids move concurrently while sharing position and velocity data.

Concepts used:

- goroutines for boid movement
- `sync.RWMutex` for coordinating reads and writes to shared simulation state
- concurrent updates to a shared grid of boid positions

Run it with:

```bash
go run ./boidsumulation
```

### `filesearch`

A concurrent file search example that walks directories and looks for matching file names.

Concepts used:

- goroutines for recursive directory traversal
- `sync.WaitGroup` for tracking active work
- `sync.Mutex` for safely appending to shared results

Run it with:

```bash
go run ./filesearch
```

Note: `filesearch/main.go` currently uses a hardcoded search path. Update that path before running it on your machine if needed.

## Focus

This repository highlights concurrent program structure, coordination of shared state, and synchronization patterns in Go.

## Conditional Variables

The matrix multiplication example also introduces `sync.Cond`, Go's conditional variable type. A conditional variable lets goroutines sleep until some shared state changes, instead of spinning in a loop and repeatedly checking the same condition.

In this repository's matrix example, the intended flow is:

- worker goroutines lock the shared state and call `cond.Wait()`
- `Wait()` releases the lock and parks the goroutine
- the main goroutine updates the matrices for the next round
- the main goroutine calls `cond.Broadcast()` to wake every waiting worker
- each worker re-acquires the lock before `Wait()` returns and continues processing

This is useful when work happens in phases and goroutines should stay idle between phases. In practice, `sync.Cond` is a good fit when:

- multiple goroutines depend on the same state transition
- a channel does not naturally model the wake-up pattern
- you need explicit control over the lock that protects the shared data

The important rule is that the condition variable and the shared state must always be reasoned about together: change the state while holding the lock, then signal with `Signal()` or `Broadcast()` so waiting goroutines wake up against a consistent view of memory.
