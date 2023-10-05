package construct_quad_tree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return dfs(grid, len(grid), 0, 0)
}

func dfs(grid [][]int, n, r, c int) *Node {
	allSame := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[r+i][c+j] != grid[r][c] {
				allSame = false
				break
			}
		}
	}
	if allSame {
		return &Node{
			Val:    grid[r][c] == 1,
			IsLeaf: true,
		}
	}
	n /= 2
	topLeft := dfs(grid, n, r, c)
	topRight := dfs(grid, n, r, c+n)
	bottomLeft := dfs(grid, n, r+n, c)
	bottomRight := dfs(grid, n, r+n, c+n)

	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}
