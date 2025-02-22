# EETF
> **E**asy **e**BPF **T**racepoint **F**inder

EETF is a straightforward tool to help you quickly locate and inspect eBPF(extended Berkeley Packet Filter) tracepoints. Built in Go, it leverages [`spf13/cobra`](https://github.com/spf13/cobra) for a TUI framework and [`koki-develop/fzf-go`](https://github.com/koki-develop/go-fzf) for interactive fuzzy searching.

## Preview

![image](https://github.com/user-attachments/assets/a5eadce2-d674-4462-80f3-edd1fc66a09b)
![image](https://github.com/user-attachments/assets/b3e9c27f-62f5-4e62-85b4-143dd138c0c5)

## Features

- **Rapid Tracepoint Discovery:** Scans `/sys/kernel/debug/tracing/events/` to list all available tracepoints.
- **Flexible Data Output:** Fetches and reformats tracepoint format data (raw, C struct, or table) from `/sys/kernel/debug/tracing/events/*/*/format` for clear console display.
  
  (Every search is conducted via `go-fzf` to fast, easy, and fuzzy-based searching experience.)

## How to Use

1. Clone the repository:
   ```bash
   git pull https://github.com/KnightChaser/EETF.git
   ```
2. Build the binary:
   ```bash
   go build .
   ```
3. Run the binary as **root** (required to access the `/sys` directory):
   ```bash
   sudo ./eetf
   ```
