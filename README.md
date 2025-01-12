# Go Slice Index out of Bounds Error

This repository contains a simple Go program that demonstrates an index out of bounds error when iterating over a slice.

## Bug Description

The program attempts to access the first element of an empty slice, resulting in a runtime panic.  This is a common error when working with slices in Go.  The solution below shows how to prevent this panic.

## Bug Reproduction

1. Clone this repository.
2. Run the `bug.go` file using `go run bug.go`.
3. Observe the runtime panic.

## Solution

The solution file, `bugSolution.go`, demonstrates a safe way to iterate over slices in Go, ensuring that index out of bounds errors are prevented.