# Load Balancer in Go

A load balancer built from scratch in Go as part of a transition from web development into systems programming.

---

## What It Does

Distributes incoming requests across multiple backend servers using a Round Robin algorithm. Automatically skips servers that are marked as unhealthy.

---

## Project Structure

load-balancer/
├── go.mod
├── main.go
└── balancer/
└── balancer.go

---

## What Has Been Built

### balancer/balancer.go

Defines the core data structures and algorithm.

**Server struct**
Represents a single backend server with three fields — Address, Health and Name.

**LoadBalancer struct**
Holds the list of servers, the current index for rotation, and a mutex for safe concurrent access.

**NewLoadBalancer()**
Constructor function that takes a slice of servers and returns a ready LoadBalancer.

**NextServer()**
The round robin algorithm. Picks the next healthy server in rotation, moves the index forward using modulo wrapping, skips unhealthy servers, and returns an error if no healthy servers are available. Protected by a mutex to prevent race conditions.

---

### main.go

Entry point for the program. Creates three test servers with Server2 marked as unhealthy. Initializes the LoadBalancer using NewLoadBalancer. Simulates 7 incoming tasks and prints which server each one is assigned to.

---

## How Round Robin Works

- Request 1 → Server1
- Request 2 → Server3  ← Server2 skipped, unhealthy
- Request 3 → Server1
- Request 4 → Server3

The current index moves forward after each pick and wraps back to zero after the last server using modulo:

````go
next = (current + 1) % total servers
````
---

## Running It

```bash
go run main.go
```

Expected output:

- Task 1 -> assigned to Server1
- Server2 is not healthy, skipping
- Task 2 -> assigned to Server3
- Task 3 -> assigned to Server1
- Server2 is not healthy, skipping
- Task 4 -> assigned to Server3
- Task 5 -> assigned to Server1
- Server2 is not healthy, skipping
- Task 6 -> assigned to Server3
- Task 7 -> assigned to Server1

---

## Requirements

Go 1.21 or higher. No external dependencies — standard library only.