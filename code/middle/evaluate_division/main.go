package evaluate_division

type Vertex struct {
	Name  string
	Value float64
}

/*
Обход графа в глубину (DFS) - это алгоритм, используемый для обхода или поиска дерева или графа.
Идея состоит в том, чтобы начать с узла и идти вглубь, пока не будет найден узел, удовлетворяющий условию поиска.

По данной задаче:
1. Делаем мапу и записываем в нее значение пары и обратное значение обратной пары
2. Проходимся по запросам и ищем в мапе начальную вершину
3. Если начальная вершина не найдена, то записываем в результат -1
4. Если начальная вершина найдена, то запускаем функцию dfs и ищем конечную вершину попутно умножая результаты выхода
*/
// calcEquation is the main application function.
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	vertexMap := make(map[string][]Vertex)
	for idx, equation := range equations {
		v1 := equation[0]
		v2 := equation[1]

		vertexMap[v1] = append(vertexMap[v1], Vertex{Name: v2, Value: values[idx]})
		vertexMap[v2] = append(vertexMap[v2], Vertex{Name: v1, Value: 1 / values[idx]})
	}

	result := make([]float64, len(queries))

	for idx, query := range queries {
		start, end := query[0], query[1]
		visited := make(map[string]bool)

		if _, ok := vertexMap[start]; !ok {
			result[idx] = -1
		} else {
			result[idx] = dfs(start, end, vertexMap, visited)
		}
	}

	return result
}

func dfs(start, end string, vertexMap map[string][]Vertex, visited map[string]bool) float64 {
	if start == end {
		return 1.0
	}

	for _, vertex := range vertexMap[start] {
		if visited[vertex.Name] {
			continue
		}

		visited[vertex.Name] = true

		if result := dfs(vertex.Name, end, vertexMap, visited); result != -1.0 {
			return result * vertex.Value
		}
	}

	return -1.0
}
