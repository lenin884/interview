package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
		1 -> 0 -> 2

	*/

	// 1. Create a graph
	graph := make(map[int][]int)

	for _, edge := range prerequisites {
		// edge[0] -> edge[1]
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	// 2. Create a visited map
	visited := make(map[int]bool)

	// 3. Create a visiting map
	visiting := make(map[int]bool)

	// 4. Iterate through the graph
	for node := range graph {
		if !dfs(node, graph, visited, visiting) {
			return false
		}
	}

	return true
}

func dfs(node int, graph map[int][]int, visited map[int]bool, visiting map[int]bool) bool {
	// 1. Check if node is in visiting map
	if visiting[node] {
		return false
	}

	// 2. Check if node is in visited map
	if visited[node] {
		return true
	}

	// 3. Add node to visiting map
	visiting[node] = true

	// 4. Iterate through the neighbours
	for _, neighbour := range graph[node] {
		if !dfs(neighbour, graph, visited, visiting) {
			return false
		}
	}

	// 5. Remove node from visiting map
	visiting[node] = false

	// 6. Add node to visited map
	visited[node] = true

	return true
}
