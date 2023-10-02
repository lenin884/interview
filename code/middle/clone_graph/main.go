package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if _, ok := visited[node]; ok {
		return visited[node]
	}

	visited[node] = &Node{Val: node.Val}
	for _, neighbor := range node.Neighbors {
		visited[node].Neighbors = append(visited[node].Neighbors, dfs(neighbor, visited))
	}

	return visited[node]
}
