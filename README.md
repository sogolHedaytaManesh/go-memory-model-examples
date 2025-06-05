<<<<<<< HEAD
# go-memory-model-examples

This repository contains practical Go code examples demonstrating key concepts from my Medium article on Go’s memory model and concurrency:

**Concurrency Without Chaos: The Untold Power of Go’s Memory Model**  
[https://medium.com/@sogol.hedayatmanesh/concurrency-without-chaos-the-untold-power-of-gos-memory-model-6a4d906ef809](https://medium.com/@sogol.hedayatmanesh/concurrency-without-chaos-the-untold-power-of-gos-memory-model-6a4d906ef809)

## Example Scenarios

- `race_conditions/` — Examples showing race conditions and fixes using mutexes.
- `happens_before/` — Demonstrates happens-before relationships using channels.
- `atomic_operations/` — Usage of atomic operations for safe concurrency.
- `visibility/` — Shows variable visibility issues and atomic fixes.

## Running the Examples

To run individual examples:

```bash
go run race_conditions/unsynchronized.go
go run happens_before/happens_before_example.go
```

## Run all examples at once
From the root directory, run:

```bash
go run main.go
```