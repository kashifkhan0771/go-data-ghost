package graph

type DirectedGraph struct {
	// adjacencyList is a map where the key is a string representing a vertex
	// and the value is a slice of strings representing the vertices that the key vertex is connected to
	adjacencyList map[string][]string
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		adjacencyList: make(map[string][]string),
	}
}

func (dg *DirectedGraph) AddVertex(vertex string) {
	dg.adjacencyList[vertex] = []string{}
}

func (dg *DirectedGraph) AddEdge(from, to string) {
	dg.adjacencyList[from] = append(dg.adjacencyList[from], to)
}

func (dg *DirectedGraph) BreadthFirstSearch(start string) []string {
	var queue []string
	var visited = make(map[string]bool)
	var result []string

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		for _, neighbour := range dg.adjacencyList[vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	return result
}

func (dg *DirectedGraph) DepthFirstSearch(start string) []string {
	var stack []string
	var visited = make(map[string]bool)
	var result []string

	stack = append(stack, start)
	visited[start] = true

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, vertex)

		for _, neighbour := range dg.adjacencyList[vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				stack = append(stack, neighbour)
			}
		}
	}
	return result
}
