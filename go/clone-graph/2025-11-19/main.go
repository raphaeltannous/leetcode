/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	graph := make(map[int]*Node)
	seen := make(map[int]bool)
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if seen[node.Val] {
			return
		}

		seen[node.Val] = true
		newNode, ok := graph[node.Val]
		if !ok {
			newNode = &Node{node.Val, make([]*Node, 0)}
			graph[node.Val] = newNode
		}

		for _, neighbor := range node.Neighbors {
			newNeighbor, ok := graph[neighbor.Val]
			if !ok {
				newNeighbor = &Node{neighbor.Val, make([]*Node, 0)}
				graph[neighbor.Val] = newNeighbor
			}

			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
			dfs(neighbor)
		}
	}
	dfs(node)

	return graph[1]
}
