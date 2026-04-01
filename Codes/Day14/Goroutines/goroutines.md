# Go Concurrency

## 1. Concurrency vs. Parallelism
* **Concurrency:** Managing multiple tasks at once. (e.g., One chef rapidly switching between cooking four different dishes).
* **Parallelism:** Executing multiple tasks at the exact same time. (e.g., Four chefs cooking four dishes simultaneously).
* **In Go:** * `runtime.GOMAXPROCS(1)`: Forces goroutines to share a single CPU core. They take turns rapidly (Concurrent, not parallel).
  * `runtime.GOMAXPROCS(runtime.NumCPU())`: Allows goroutines to run simultaneously across multiple CPU cores (Parallel).

## 2. The Go Scheduler
* Standard Operating System (OS) threads are "heavy" and consume a lot of memory.
* **M:N Scheduling:** Go's built-in scheduler maps **M** lightweight *goroutines* onto **N** heavier *OS threads*.
* **How it works:** Think of the scheduler as an efficient manager. If a goroutine gets blocked (like waiting for a network request), the manager instantly swaps it off the OS thread and replaces it with a ready-to-run goroutine. The CPU stays constantly busy.

## 3. Philosophy
* "Do not communicate by sharing memory; instead, share memory by communicating."

This phrase represents a massive shift in how we think about managing data between concurrent tasks.

## 1. The Old Way: "Communicate by sharing memory"
In traditional programming languages (like C++, Java, or Python), threads talk to each other by reading and writing to the exact same variables in memory. 

* **The Analogy:** Imagine two chefs trying to chop different vegetables on the exact same cutting board at the exact same time. 
* **The Problem:** They will inevitably bump knives or ruin each other's work. In programming, this is called a **race condition**. 
* **The Traditional Fix:** To prevent this, you have to use a **Mutex** (Mutual Exclusion lock). One chef locks the door to the kitchen, chops their veggies, unlocks the door, and then the next chef goes in. It is slow, prone to deadlocks (everyone waiting for a key that got lost), and notoriously hard to debug.

## 2. The Go Way: "Share memory by communicating"
Instead of letting multiple goroutines fight over the same piece of memory, Go encourages you to pass the actual data (the memory) back and forth between goroutines using **Channels**. 

* **The Analogy:** Instead of sharing one cutting board, Chef A chops the carrots and puts them on a conveyor belt (the channel). Chef B waits at the end of the conveyor belt, picks up the carrots, and puts them in the soup. 
* **The Solution:** Ownership of the data is passed from one goroutine to another. At any given instant, **only one goroutine has access to the data**. 
* **Why it's better:** Because goroutines never access the same memory at the exact same time, you don't need to manually lock and unlock mutexes. The synchronization happens naturally as a byproduct of sending and receiving data on the channel.

## 4. Goroutine Internals 
* An OS thread typically starts with a large, fixed memory stack (usually 1-2 Megabytes). A goroutine starts with a tiny stack of just **2 Kilobytes**. This is why you can run hundreds of thousands of them at once.
* **Dynamic Stack Growth:** If a goroutine needs more memory, the Go runtime automatically allocates more space, growing and shrinking the stack dynamically as needed.
* **Ultra-Fast Context Switching:** Switching between OS threads requires asking the operating system kernel for permission, which is slow. Go switches goroutines entirely in "user space" (handled by the Go runtime), making the switch significantly faster and cheaper.
* **The GMP Architecture:** The Go scheduler relies on three core entities to make this work:
  * **G (Goroutine):** The actual Go code/task you want to execute.
  * **M (Machine):** The underlying OS thread executing the code.
  * **P (Processor):** A logical processor that holds a queue of runnable 'G's. The number of 'P's is determined by `GOMAXPROCS`. An 'M' must acquire a 'P' to execute Go code.