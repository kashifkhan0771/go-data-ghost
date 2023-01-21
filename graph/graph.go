package graph

type Graph struct {
	// adjacencyList is a map where the key is a string representing a vertex
	// and the value is a slice of strings representing the vertices that the key vertex is connected to
	adjacencyList map[string][]string
}

// NewGraph creates and returns a new graph
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
	}
}

// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(vertex string) {
	g.adjacencyList[vertex] = []string{}
}

// AddEdge adds a new edge to the graph connecting two vertices
func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
}

// RemoveVertex removes a vertex from the graph and any edges connected to it
func (g *Graph) RemoveVertex(vertex string) {
	delete(g.adjacencyList, vertex)
	for _, v := range g.adjacencyList {
		for i, connectedVertex := range v {
			if connectedVertex == vertex {
				g.adjacencyList[connectedVertex] = append(v[:i], v[i+1:]...)
			}
		}
	}
}

func (g *Graph) BreadthFirstSearch(start string) []string {
	var queue []string
	var visited = make(map[string]bool)
	var result []string

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		for _, neighbour := range g.adjacencyList[vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	return result
}

func (g *Graph) DepthFirstSearch(start string) []string {
	var stack []string
	var visited = make(map[string]bool)
	var result []string

	stack = append(stack, start)
	visited[start] = true

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, vertex)

		for _, neighbour := range g.adjacencyList[vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				stack = append(stack, neighbour)
			}
		}
	}
	return result
}
