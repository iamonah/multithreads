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
