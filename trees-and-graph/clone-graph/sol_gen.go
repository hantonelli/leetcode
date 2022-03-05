package clone_graph

import "github.com/hantonelli/leetcode/ds"

func cloneGraphG(node *Node) *Node {
	if node == nil {
		return nil
	}

	rootCopy := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0, len(node.Neighbors)),
	}

	clonedNodes := map[int]*Node{ // map[id]copied
		node.Val: rootCopy,
	}

	l := ds.NewStack3[ClonePair]()
	l.Push(ClonePair{original: node, copied: rootCopy})

	for !l.IsEmpty() {
		currentPair := l.Pop()
		for _, originalNe := range currentPair.original.Neighbors {
			clonedNe, ok := clonedNodes[originalNe.Val]
			if !ok {
				clonedNe = &Node{
					Val:       originalNe.Val,
					Neighbors: make([]*Node, 0, len(originalNe.Neighbors)),
				}
				clonedNodes[originalNe.Val] = clonedNe
				l.Push(ClonePair{original: originalNe, copied: clonedNe})
			}
			currentPair.copied.Neighbors = append(currentPair.copied.Neighbors, clonedNe)
		}
	}
	return rootCopy
}
