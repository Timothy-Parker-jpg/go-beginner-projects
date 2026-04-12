# Port Scanner

A CLI tool that scans a host for open TCP ports concurrently using goroutines and channels.

---

## What You'll Learn

- Launching goroutines with `go func()`
- Using a `sync.WaitGroup` to wait for all goroutines to finish
- Using a buffered channel as a worker pool (semaphore pattern)
- Sorting results (ports come back out of order from concurrent scans)
- TCP connection attempts with `net.DialTimeout`
- Working with `time.Duration`

---

## Project Structure

```
port-scanner/
├── main.go
└── scanner.go
```

### `main.go`
Parses flags:
- `-host` — target hostname or IP (default: `"localhost"`)
- `-start` — start of port range (default: `1`)
- `-end` — end of port range (default: `1024`)
- `-timeout` — connection timeout in milliseconds (default: `200`)
- `-workers` — max concurrent goroutines (default: `100`)

Calls `Scan()` and prints open ports sorted in ascending order.

### `scanner.go`
The `Scan(host string, startPort, endPort, workers int, timeout time.Duration) []int` function:

```
1. Create a results channel: make(chan int, endPort-startPort+1)
2. Create a semaphore channel: make(chan struct{}, workers)
3. Create a sync.WaitGroup
4. For each port in range:
   a. wg.Add(1)
   b. Launch a goroutine:
      - Send to semaphore channel (blocks if workers is full)
      - Try net.DialTimeout("tcp", host:port, timeout)
      - If no error: send port to results channel; close the connection
      - Read from semaphore channel (releases one slot)
      - wg.Done()
5. After loop: go func() { wg.Wait(); close(results) }()
6. Collect all values from results into a []int
7. Sort the slice with sort.Ints()
8. Return
```

---

## How It Works

```
$ go run . -host scanme.nmap.org -start 1 -end 1000 -workers 200

Scanning scanme.nmap.org (ports 1–1000)...

Open ports:
  22  (SSH)
  80  (HTTP)

Scan complete in 3.2s
```

---

## The Semaphore Pattern (Key Concept)

```go
sem := make(chan struct{}, maxWorkers) // buffered channel acts as semaphore

go func() {
    sem <- struct{}{} // acquire slot (blocks when full)
    defer func() { <-sem }() // release slot when done
    // ... do work
}()
```

This pattern limits how many goroutines run simultaneously — critical when scanning thousands of ports.

---

## ⚠️ Legal & Ethical Notice

Only scan hosts you own or have explicit permission to scan. Unauthorized port scanning may be illegal in your jurisdiction.

---

## Suggested Extensions

- Map common port numbers to service names (22 → SSH, 80 → HTTP, 443 → HTTPS, etc.)
- Add UDP scanning (note: UDP is connectionless and harder to probe)
- Add a progress bar showing how many ports have been checked
- Time each scan and report total duration
