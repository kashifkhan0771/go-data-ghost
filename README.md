# go-data-ghost

A library for various data structures in Go.

## Installation

```bash
go get github.com/kashifkhan0771/go-data-ghost
````

## Features

- Simple and double linked list with Append, Prepend, InsertAfter and DeleteNode methods
- Stack with Push, Pop, Peek and Size methods
- Queue with Enqueue, Dequeue, Peek and Size methods
- Tree with Insert, Search, InOrderTraversal, PreOrderTraversal, PostOrderTraversal, Delete, Minimum, Maximum, Size and Clear methods
- Graph
- Hashmap

## Usage

To use this library in your project, import the desired package and create a new instance of the data structure.

For Linkedlist:

````
import "github.com/kashifkhan0771/go-data-ghost/linkedlist"

list := linkedlist.NewLinkedList(true)
list.Append(1)
list.Append(2)
list.Prepend(0)
````

For Stack:
````
import "github.com/kashifkhan0771/go-data-ghost/stack"

stack := stack.NewStack()
stack.Push(1)
stack.Push(2)
val, _ := stack.Pop()
````

For Queue
````
import "github.com/kashifkhan0771/go-data-ghost/queue"

queue := queue.NewQueue()
queue.Enqueue(1)
queue.Enqueue(2)
val, _ := queue.Dequeue()
````

For Tree
````
package main

import (
    "fmt"
    "github.com/kashifkhan0771/go-data-ghost/tree"
)

func main() {
    t := tree.NewTree()
    t.Insert(10)
    t.Insert(5)
    t.Insert(15)
    t.Insert(2)
    t.Insert(7)
    t.Insert(12)
    t.Insert(20)
    fmt.Println(t.Minimum()) // 2
    fmt.Println(t.Maximum()) // 20
    t.Delete(5)
    fmt.Println(t.Size()) // 6
    t.Clear()
    fmt.Println(t.Size()) // 0
}
````
