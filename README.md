# Linked List Library

This library provides an implementation of a simple and double linked list in Go.

## Features
- Append a value to the end of the list
- Prepend a value to the beginning of the list
- Insert a value after a specific node
- Delete a specific node
- Traverse the list and return all the values in the list
- Return the size of the list
- Find a specific node in the list based on its value
- Reverse the list
- Sort the list using quicksort or merge sort
- Merge two lists

## Usage
To use this library in your project, import the `linkedlist` package and create a new list with `NewLinkedList(isDouble bool)` function.
```go
import "linkedlist"

list := linkedlist.NewLinkedList(true)
