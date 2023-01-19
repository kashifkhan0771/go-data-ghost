# Data Structure Library

This library provides an implementation of various data structures in Go, including a simple and double linked list,
stack, queue, tree, graph and hashmap. These data structures can be used to perform various operations such as
insertion, deletion, searching and sorting efficiently.

## Features

- Simple and double linked list with Append, Prepend, InsertAfter and DeleteNode methods
- Stack with Push, Pop, Peek and Size methods
- Queue
- Tree
- Graph
- Hashmap

## Usage

To use this library in your project, import the desired package and create a new instance of the data structure.

For Linkedlist:

````
import "github.com/username/data-structure-lib/linkedlist"

list := linkedlist.NewLinkedList(true)
list.Append(1)
list.Append(2)
list.Prepend(0)
````

For Stack:

````
import "github.com/username/data-structure-lib/stack"

stack := stack.NewStack()
stack.Push(1)
stack.Push(2)
val, _ := stack.Pop()
````
