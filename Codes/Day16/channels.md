## Go Channels

Channels are one of Go’s most distinctive features, enabling safe and elegant concurrent communication. 

---

## Table of Contents (Part 1)

1. [Introduction to Channels](#introduction)
2. [Why Channels? The CSP Model](#why-channels)
3. [Channels vs Mutexes: When to Use What](#channels-vs-mutexes)
4. [Creating a Channel](#creating-a-channel)
5. [Channel Types: Directional Channels](#directional-channels)
6. [Channel Capacity: Buffered vs Unbuffered](#channel-capacity)
7. [Sending and Receiving: The Basics](#sending-and-receiving)
8. [Unbuffered Channels: Synchronous Communication](#unbuffered-channels)
9. [Buffered Channels: Asynchronous Communication](#buffered-channels)
10. [Closing Channels](#closing-channels)
11. [Ranging Over a Channel](#ranging-over-a-channel)
12. [The `select` Statement](#select-statement)
13. [Channel Length and Capacity](#len-and-cap)
14. [Nil Channels](#nil-channels)
15. [Common Patterns (Intro)](#common-patterns)

---

## 1. Introduction to Channels <a name="introduction"></a>

In Go, **channels** are typed conduits through which you can send and receive values with the channel operator, `<-`. They are the primary way to synchronize goroutines and pass data between them safely without explicit locks or condition variables.

Think of a channel as a pipe: one goroutine puts data in one end, and another goroutine takes data out the other end. The language runtime ensures that this transfer is synchronized and race-free.

---

## 2. Why Channels? The CSP Model <a name="why-channels"></a>

Go’s concurrency model is inspired by **Communicating Sequential Processes (CSP)**, a formal language for describing patterns of interaction in concurrent systems. The core idea:

> **Don’t communicate by sharing memory; instead, share memory by communicating.**

In traditional threading models, you protect shared data with mutexes. In Go, you often send data over a channel to another goroutine, so only one goroutine has access at a time. This approach can make concurrent programs easier to reason about and less error-prone.

---

## 3. Channels vs Mutexes: <a name="channels-vs-mutexes"></a>

Both channels and mutexes solve concurrency problems, but they are suited for different situations.

| **Scenario** | **Prefer Channels** | **Prefer Mutexes** |
|--------------|----------------------|---------------------|
| Data ownership transfer | ✅ Channel hands over data to another goroutine | ❌ |
| Streaming data / pipeline | ✅ Natural fit for passing data through stages | ❌ |
| Fan-out, fan-in, worker pools | ✅ Channels orchestrate many workers | ❌ |
| Select over multiple conditions | ✅ `select` with channels handles multiple cases | ❌ |
| Simple counter / flag protection | ❌ Overkill | ✅ Mutex or atomic |
| Protecting a cache or struct field | ❌ Awkward | ✅ Mutex protects critical sections |
| Performance‑sensitive low‑level code | ❌ Channels have overhead | ✅ Mutex may be faster |

**Rule of thumb:** Use channels when you need to coordinate goroutines or pass ownership of data. Use mutexes when you need to protect a small amount of shared state.

---

## 4. Creating a Channel <a name="creating-a-channel"></a>

A channel is created with the built-in `make` function, specifying the type of values it will carry.

```go
// Unbuffered channel of integers
ch := make(chan int)

// Buffered channel of strings with capacity 5
chStr := make(chan string, 5)
```

You can also declare a channel variable without initializing it. Such a channel is `nil` and must be initialized before use.

```go
var ch chan int // ch == nil
ch = make(chan int)
```

---

## 5. Channel Types: Directional Channels <a name="directional-channels"></a>

Channels can be constrained to only send or only receive. This is enforced at compile time and helps express intent and prevent misuse.

- `chan T` – bidirectional (send and receive)
- `chan<- T` – send‑only channel
- `<-chan T` – receive‑only channel

```go
func sendData(ch chan<- int) {
    ch <- 42 // OK
    // x := <-ch // compile error: receive from send-only channel
}

func receiveData(ch <-chan int) {
    x := <-ch // OK
    // ch <- 100 // compile error: send to receive-only channel
}
```

You can convert a bidirectional channel to directional ones implicitly when passing to functions.

---

## 6. Channel Capacity: Buffered vs Unbuffered <a name="channel-capacity"></a>

The capacity of a channel is the number of elements that can be held in the channel’s buffer.

- **Unbuffered channel** – capacity 0 (created with `make(chan T)`). Sends and receives happen synchronously.
- **Buffered channel** – capacity > 0 (created with `make(chan T, capacity)`). Sends block only when the buffer is full; receives block only when the buffer is empty.

Buffered channels decouple the sender and receiver to a degree, but they are not asynchronous queues for unbounded data – the buffer size must be chosen wisely.

---

## 7. Sending and Receiving: The Basics <a name="sending-and-receiving"></a>

**Send** a value onto a channel:

```go
ch <- v   // v is sent to channel ch
```

**Receive** a value from a channel:

```go
v := <-ch // receive and assign
<-ch      // receive and discard
```

Receives also return an optional second boolean indicating whether the value was sent before the channel was closed:

```go
v, ok := <-ch
if ok {
    // v was successfully received
} else {
    // channel is closed and empty
}
```

Sends and receives are blocking operations by default. The exact behavior depends on whether the channel is buffered or unbuffered.

---

## 8. Unbuffered Channels: Synchronous Communication <a name="unbuffered-channels"></a>

An unbuffered channel has no capacity. For a send to complete, there must be a receiver ready to take the value. Both the sender and the receiver block until the other side is ready. This creates a **synchronisation point** – data is passed directly from sender to receiver.

### Example: Unbuffered Channel

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)

    // Start a receiver goroutine
    go func() {
        value := <-ch
        fmt.Println("Received:", value)
    }()

    // Sender (main goroutine)
    ch <- 42
    fmt.Println("Sent 42")

    // Give receiver time to print (in real code use sync.WaitGroup)
    time.Sleep(100 * time.Millisecond)
}
```

**Output:**
```
Received: 42
Sent 42
```

Note: The main goroutine blocks on `ch <- 42` until the receiver is ready. Without the receiver goroutine, this would deadlock.

### Why use unbuffered channels?
- Guarantee that a send and receive happen at the same moment.
- Implement handshakes or confirmations between goroutines.
- Ensure that data is processed as soon as it is sent.

---

## 9. Buffered Channels: Asynchronous Communication <a name="buffered-channels"></a>

A buffered channel has a fixed capacity. Sends add an element to the buffer and proceed immediately **unless** the buffer is full. Receives remove an element from the buffer and proceed immediately **unless** the buffer is empty.

### Example: Buffered Channel

```go
package main

import (
    "fmt"
)

func main() {
    ch := make(chan string, 2)

    // Send two strings without a receiver
    ch <- "hello"
    ch <- "world"
    // ch <- "third" // would block because buffer is full

    fmt.Println(<-ch) // hello
    fmt.Println(<-ch) // world
}
```

**Output:**
```
hello
world
```

### Buffer behaviour
- **Send** to a buffered channel only blocks when the buffer is full.
- **Receive** from a buffered channel only blocks when the buffer is empty.
- The channel can hold up to `cap(ch)` elements without blocking.

### Important: Buffered channels are not queues for unlimited data
Choosing the buffer size is a design decision. Too small may cause unnecessary blocking; too large may hide backpressure and lead to high memory usage or stale data.

---

## 10. Closing Channels <a name="closing-channels"></a>

A channel can be closed to indicate that no more values will be sent. This is done with the `close()` built-in function.

```go
close(ch)
```

### Rules and consequences
- Only the sender should close a channel; never the receiver. Closing a channel is a signal that no more values are coming.
- Sending on a closed channel causes a **panic**.
- Receiving from a closed channel:
  - If there are still buffered values, they can be received in order.
  - After all buffered values are drained, receives return the zero value of the channel’s type and the second boolean `ok` becomes `false`.
- Closing a closed channel also panics.
- Channels don’t need to be closed unless you need to signal completion to a receiver (e.g., with a `range` loop).

### Example: Receiving after close

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    close(ch)

    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
    fmt.Println(<-ch) // 0 (zero value), ok would be false
    v, ok := <-ch
    fmt.Println(v, ok) // 0 false
}
```

---

## 11. Ranging Over a Channel <a name="ranging-over-a-channel"></a>

The `for range` loop can be used to receive values from a channel until it is closed and drained.

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 2)
    ch <- "apple"
    ch <- "banana"
    close(ch)

    for fruit := range ch {
        fmt.Println(fruit)
    }
}
```

**Output:**
```
apple
banana
```

The loop automatically stops after the last value is received and the channel is closed. If you forget to close the channel, `range` will block forever, causing a deadlock.

---

## 12. The `select` Statement <a name="select-statement"></a>

`select` lets a goroutine wait on multiple channel operations simultaneously. It blocks until one of its cases can proceed, then executes that case. If multiple are ready, it picks one pseudo-randomly.

### Basic syntax

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case ch3 <- 42:
    fmt.Println("Sent 42 to ch3")
default:
    fmt.Println("No communication ready")
}
```

- The `default` case runs immediately if no other case is ready, making `select` non‑blocking.
- `select` with no cases (i.e., `select {}`) blocks forever.

### Example: Multiplexing two channels

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received", msg2)
        }
    }
}
```

**Output (order may vary due to sleep):**
```
Received one
Received two
```

### Timeouts using `select`

A common pattern is to combine `select` with `time.After` to implement timeouts.

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout, no message received")
}
```

---

## 13. Channel Length and Capacity <a name="len-and-cap"></a>

- `len(ch)` returns the number of elements currently queued in the channel buffer.
- `cap(ch)` returns the buffer capacity.

For an unbuffered channel, `len` is always 0 and `cap` is 0.

These functions are rarely used in production because inspecting the buffer length introduces a race condition – by the time you act on the length, it may have changed. They are more useful for debugging or in special cases (e.g., implementing a custom backpressure strategy).

```go
ch := make(chan int, 5)
ch <- 1
ch <- 2
fmt.Println(len(ch)) // 2
fmt.Println(cap(ch)) // 5
```

---

## 14. Nil Channels <a name="nil-channels"></a>

A nil channel is a channel variable that hasn’t been initialised (or was set to `nil`). Operations on nil channels behave in a special way:

- **Sending** to a nil channel blocks forever.
- **Receiving** from a nil channel blocks forever.
- **Closing** a nil channel panics.

Nil channels can be useful in `select` statements to disable cases dynamically.

### Example: Disabling a case with nil

```go
package main

import "fmt"

func main() {
    var ch chan int // nil

    select {
    case <-ch:
        // never reached because ch is nil
    default:
        fmt.Println("Default case runs")
    }
}
```

---
