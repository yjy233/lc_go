# sync.Cond Quick Reference

## Overview
`sync.Cond` is a condition variable that allows goroutines to wait for or announce events. It's useful for coordinating goroutines based on shared state conditions.

## Essential API
```go
cond := sync.NewCond(&mu)  // Create with mutex
cond.Wait()                // Wait for condition (releases mutex, reacquires on wake)
cond.Signal()              // Wake one waiting goroutine
cond.Broadcast()           // Wake all waiting goroutines
```

## Critical Rules
1. **Always hold the mutex** when calling Wait(), Signal(), or Broadcast()
2. **Use Wait() in a loop** to handle spurious wakeups:
   ```go
   for !condition {
       cond.Wait()
   }
   ```
3. **Use defer mu.Unlock()** to ensure mutex is always released

## Common Use Cases
- **Producer-Consumer**: Signal() when items are added to queue
- **Broadcast Events**: Broadcast() to wake all waiting workers
- **Worker Pools**: Signal() to wake one worker for new tasks
- **Barriers**: Coordinate multiple goroutines to reach same point

## Quick Pattern Examples

### Producer-Consumer
```go
// Consumer waits for items
for len(queue) == 0 {
    cond.Wait()
}

// Producer adds item and signals
queue = append(queue, item)
cond.Signal()
```

### Broadcast to All Workers
```go
// Workers wait for ready signal
for !ready {
    cond.Wait()
}

// Coordinator sets ready and broadcasts
ready = true
cond.Broadcast()
```

## Complete Examples
See `sync_cond_examples/` directory for:
