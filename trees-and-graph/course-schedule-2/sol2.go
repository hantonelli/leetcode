package main

func findOrder2(numCourses int, prerequisites [][]int) []int {
	nexts := make([][]int, numCourses)
	visited := make([]VisitedState, numCourses)
	stack := make([]int, 0, numCourses)
	recursive := false

	for _, v := range prerequisites {
		nexts[v[0]] = append(nexts[v[0]], v[1])
	}

	for i := 0; i < numCourses; i++ {
		dfs2(i, nexts, visited, &stack, &recursive)
	}

	if recursive {
		return []int{}
	}
	return stack
}

func dfs2(i int, nexts [][]int, visited []VisitedState, stack *[]int, recursive *bool) {
	// If it had already been flag as recursive, just return
	if *recursive {
		return
	}
	// If the current node is already in a Visiting state, then this graph is recursive,
	// so flag it as so and return
	if visited[i] == Visiting {
		*recursive = true
		return
	}
	// If we already visited the node, skip it
	if visited[i] == Visited {
		return
	}
	// If this node have child
	if len(nexts[i]) != 0 {
		// Mark that we are visiting this node
		visited[i] = Visiting
		// Process each child before adding this to the stack
		for _, next := range nexts[i] {
			dfs2(next, nexts, visited, stack, recursive)
		}
	}
	// Add this node to the stack and flag it as Visited
	*stack = append(*stack, i)
	visited[i] = Visited
}
