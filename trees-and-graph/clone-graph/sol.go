package clone_graph

import "container/list"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

type ClonePair struct {
	original *Node
	copied   *Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	rootCopy := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0, len(node.Neighbors)),
	}

	clonedNodes := map[int]*Node{} // map[id]copied
	clonedNodes[node.Val] = rootCopy

	l := list.New()
	l.PushFront(&ClonePair{original: node, copied: rootCopy})

	for l.Len() != 0 {
		e := l.Front()
		currentPair := e.Value.(*ClonePair)
		l.Remove(e)

		for _, originalNe := range currentPair.original.Neighbors {
			clonedNe, ok := clonedNodes[originalNe.Val]
			if !ok {
				clonedNe = &Node{
					Val:       originalNe.Val,
					Neighbors: make([]*Node, 0, len(originalNe.Neighbors)),
				}
				clonedNodes[originalNe.Val] = clonedNe
				l.PushFront(&ClonePair{original: originalNe, copied: clonedNe})
			}
			currentPair.copied.Neighbors = append(currentPair.copied.Neighbors, clonedNe)
		}
	}
	return rootCopy
}
