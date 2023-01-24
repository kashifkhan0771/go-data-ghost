# go-data-ghost

A Go library for commonly used data structures, including linked list, stack, queue, tree, graph, hashmap, and set. This library aims to provide a simple and efficient implementation of these data structures for use in various Go projects.
## Installation

There are several ways to install the go-data-ghost library.
### Method 1: Go Get
The simplest way to install the library is to use the go get command:

```go get github.com/kashifkhan0771/go-data-ghost```
### Method 2: Download
You can also download the library from the releases page and manually install it in your Go environment.

### Method 3: Git
If you prefer to use Git, you can clone the repository and use it in your project:

```git clone https://github.com/kashifkhan0771/go-data-ghost.git```
<hr>

# Usage
## Linked List
````
package main

import (
"fmt"
"github.com/kashifkhan0771/go-data-ghost/linkedlist"
)

func main() {
list := linkedlist.NewLinkedList(false)
list.Append(1)
list.Append(2)
list.Append(3)
list.Prepend(0)

    currentNode := list.head
    for currentNode != nil {
        fmt.Println(currentNode.value)
        currentNode = currentNode.next
    }
}
````

## Stack
````
package main

import (
	"fmt"
	"github.com/kashifkhan0771/go-data-ghost/stack"
)

func main() {
    s := stack.NewStack()
    s.Push(1)
    s.Push(2)
    s.Push(3)

    for !s.IsEmpty() {
        fmt.Println(s.Pop())
    }
}
````

## Queue
````
package main

import (
	"fmt"
	"github.com/kashifkhan0771/go-data-ghost/queue"
)

func main() {
    q := queue.NewQueue()
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    for !q.IsEmpty() {
        fmt.Println(q.Dequeue())
    }
}
````

## Tree
````
package main

import (
"fmt"
"github.com/kashifkhan0771/go-data-ghost/tree"
)

func main() {
t := tree.NewTree()
t.Insert(5)
t.Insert(3)
t.Insert(8)
t.Insert(1)
t.Insert(4)
t.Insert(7)
t.Insert(10)
t.InOrder()
t.Delete(5)
t.InOrder()
fmt.Println(t.Size())
}
````

## Graph
````
package main

import (
"fmt"
"github.com/kashifkhan0771/go-data-ghost/graph"
)

func main() {
g := graph.NewGraph(true)
g.AddVertex("A")
g.AddVertex("B")
g.AddVertex("C")
g.AddVertex("D")
g.AddEdge("A", "B", 1)
g.AddEdge("A", "C", 2)
g.AddEdge("B", "D", 3)
g.AddEdge("C", "D", 4)
g.Print()
g.BFS("A")
g.DFS("A")
}
````

## Hashmap
````
package main

import (
"fmt"
"github.com/kashifkhan0771/go-data-ghost/hashmap"
)

func main() {
h := hashmap.NewHashMap(10, 0.8)
h.Set("name", "John Doe")
h.Set("age", 30)
h.Set("gender", "male")
fmt.Println("Size:", h.Len())
fmt.Println("Load Factor:", h.LoadFactor())
fmt.Println("Keys:", h.Keys())
fmt.Println("Values:", h.Values())
}
````

## Set
````
package main

import (
	"fmt"
	"go-data-ghost/set"
)

func main() {
	s := set.NewSet()
	s.Add("apple")
	s.Add("banana")
	s.Add("orange")
	s.Add("apple") // duplicate, will not be added
	
	fmt.Println("Size of set: ", s.Size())
	// Output: Size of set:  3
	
	fmt.Println("Set contains apple: ", s.Contains("apple"))
	// Output: Set contains apple:  true
	
	s.Remove("banana")
	
	fmt.Println("Size of set after removing banana: ", s.Size())
	// Output: Size of set after removing banana:  2
	
	s.Clear()
	fmt.Println("Size of set after clear: ", s.Size())
	// Output: Size of set after clear:  0
}
````

# Contribution
All contributions, bug reports, bug fixes, documentation improvements, enhancements and ideas are welcome
