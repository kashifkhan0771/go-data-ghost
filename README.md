# go-data-ghost
A collection of popular data structures implemented in Go for use as a library.
## Installation

```sh
go get github.com/kashifkhan0771/go-data-ghost
```
## Usage
````
package main

import (
"fmt"
"github.com/kashifkhan0771/go-data-ghost/linkedlist"
"github.com/kashifkhan0771/go-data-ghost/stack"
"github.com/kashifkhan0771/go-data-ghost/queue"
"github.com/kashifkhan0771/go-data-ghost/tree"
"github.com/kashifkhan0771/go-data-ghost/graph"
)

func main() {
// Linked List usage
list := linkedlist.NewLinkedList(true)
list.Append(1)
list.Append(2)
list.Append(3)
fmt.Println(list)

    // Stack usage
    s := stack.NewStack()
    s.Push(1)
    s.Push(2)
    s.Push(3)
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())

    // Queue usage
    q := queue.NewQueue()
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    fmt.Println(q.Dequeue())
    fmt.Println(q.Dequeue())
    fmt.Println(q.Dequeue())

    // Tree usage
    t := tree.NewTree()
    t.Insert(10)
    t.Insert(5)
    t.Insert(15)
    fmt.Println(t.Minimum())
    fmt.Println(t.Maximum())
    t.Delete(5)
    fmt.Println(t.Size())
    t.Clear()
    fmt.Println(t.Size())

    // Graph usage - Undirected
    g := graph.NewGraph()
    g.AddVertex("A")
    g.AddVertex("B")
    g.AddVertex("C")
    g.AddEdge("A", "B")
    g.AddEdge("B", "C")
    fmt.Println(g.BreadthFirstSearch("A"))
    fmt.Println(g.DepthFirstSearch("A"))

    // Directed Graph usage
    dg := graph.NewDirectedGraph()
    dg.AddVertex("A")
    dg.AddVertex("B")
    dg.AddVertex("C")
    dg.AddEdge("A", "B")
    dg.AddEdge("B", "C")
    fmt.Println(dg.BreadthFirstSearch("A"))
    fmt.Println(dg.DepthFirstSearch("A"))
}
````