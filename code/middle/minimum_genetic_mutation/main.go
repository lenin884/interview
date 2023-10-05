package minimum_genetic_mutation

// minMutation use BFS to find the shortest path
func minMutation(startGene string, endGene string, bank []string) int {
	queue := []string{startGene}

	visited := make(map[string]bool)

	mutations := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			gene := queue[0]
			queue = queue[1:]

			if gene == endGene {
				return mutations
			}

			if visited[gene] {
				continue
			}

			visited[gene] = true

			for _, next := range bank {
				if !visited[next] && isMutation(gene, next) {
					queue = append(queue, next)
				}
			}
		}

		mutations++
	}

	return -1
}

func isMutation(gene1 string, gene2 string) bool {
	diff := 0

	for i := 0; i < len(gene1); i++ {
		if gene1[i] != gene2[i] {
			diff++
		}

		if diff > 1 {
			return false
		}
	}

	return true
}
