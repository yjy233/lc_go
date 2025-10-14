# How to Use sync.Cond

## What is sync.Cond?

`sync.Cond` is a condition variable that allows goroutines to wait for or signal the occurrence of an event. It's used when you need goroutines to wait for a specific condition to become true.

## Basic Usage

### 1. Create a Cond
```go
var mu sync.Mutex
cond := sync.NewCond(&mu)
```

### 2. Three Main Methods
- `cond.Wait()` - Wait for a condition to be true
- `cond.Signal()` - Wake up one waiting goroutine
- `cond.Broadcast()` - Wake up all waiting goroutines

## Essential Pattern

**Always follow this pattern:**

```go
mu.Lock()
defer mu.Unlock()

// Use a loop to check condition
for !condition {
    cond.Wait()  // Releases mutex, waits, then reacquires mutex
}
// Now condition is true, do work
```

## Simple Example: Producer-Consumer

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    items := []string{}

    // Consumer goroutine
    go func() {
        mu.Lock()
        defer mu.Unlock()
        
        // Wait until there are items
        for len(items) == 0 {
            fmt.Println("Consumer waiting...")
            cond.Wait()
        }
        
        // Consume item
        item := items[0]
        items = items[1:]
        fmt.Printf("Consumed: %s\n", item)
    }()

    // Producer
    time.Sleep(2 * time.Second)
    mu.Lock()
    items = append(items, "apple")
    fmt.Println("Produced: apple")
    cond.Signal() // Wake up consumer
    mu.Unlock()
    
    time.Sleep(1 * time.Second)
}
```

## Key Rules

1. **Always hold the mutex** when calling Wait(), Signal(), or Broadcast()
2. **Use Wait() in a loop**, not an if statement
3. **Use defer mu.Unlock()** to ensure mutex is released

## When to Use Signal vs Broadcast

- **Signal()**: Wake up ONE waiting goroutine
  ```go
  cond.Signal() // Wake up one worker
  ```

- **Broadcast()**: Wake up ALL waiting goroutines
  ```go
  cond.Broadcast() // Wake up all workers
  ```

## Common Use Cases

### 1. Producer-Consumer Queue
```go
// Consumer waits for items
for len(queue) == 0 {
    cond.Wait()
}

// Producer adds item
queue = append(queue, item)
cond.Signal()
```

### 2. Waiting for Ready State
```go
// Workers wait for ready signal
for !ready {
    cond.Wait()
}

// Coordinator sets ready
ready = true
cond.Broadcast() // Wake all workers
```

### 3. Worker Pool
```go
// Worker waits for tasks
for len(tasks) == 0 && !shutdown {
    cond.Wait()
}

// Add task to pool
tasks = append(tasks, newTask)
cond.Signal() // Wake one worker
```

## Complete Working Example

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    ready := false

    // Start 3 workers
    for i := 0; i < 3; i++ {
        go func(id int) {
            mu.Lock()
            defer mu.Unlock()
            
            // Wait for ready signal
            for !ready {
                fmt.Printf("Worker %d waiting...\n", id)
                cond.Wait()
            }
            
            fmt.Printf("Worker %d started!\n", id)
        }(i)
    }

    // Wait a bit, then signal all workers
    time.Sleep(2 * time.Second)
    
    mu.Lock()
    ready = true
    fmt.Println("Signaling all workers...")
    cond.Broadcast() // Wake up all workers
    mu.Unlock()
    
    time.Sleep(1 * time.Second)
}
```

## Common Mistakes to Avoid

1. **Don't use if instead of for:**
   ```go
   // WRONG
   if !condition {
       cond.Wait()
   }
   
   // CORRECT
   for !condition {
       cond.Wait()
   }
   ```

2. **Don't call methods without holding mutex:**
   ```go
   // WRONG
   cond.Signal()
   
   // CORRECT
   mu.Lock()
   cond.Signal()
   mu.Unlock()
   ```

3. **Don't forget to unlock:**
   ```go
   // RISKY
   mu.Lock()
   // ... might return early
   mu.Unlock()
   
   // SAFE
   mu.Lock()
   defer mu.Unlock()
   ```

## Summary

`sync.Cond` is perfect when you need goroutines to wait for conditions and coordinate based on shared state. Remember the three key rules: hold the mutex, use loops with Wait(), and choose Signal vs Broadcast appropriately.
