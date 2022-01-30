package main

import "container/list"

type VisitedState int

const (
	_ VisitedState = iota
	Visiting
	Visited
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	var course, prereq int
	coursePrereq := make(map[int][]int, numCourses)
	visited := make(map[int]VisitedState, numCourses)
	stack := list.New()
	recursive := false

	for _, v := range prerequisites {
		course, prereq = v[0], v[1]
		coursePrereq[course] = append(coursePrereq[course], prereq)
	}

	for course := 0; course < numCourses; course++ {
		dfs(course, coursePrereq, visited, stack, &recursive)
	}

	if recursive {
		return []int{}
	}

	res := make([]int, 0, numCourses)
	var el *list.Element
	for stack.Len() != 0 {
		el = stack.Front()
		res = append(res, el.Value.(int))
		stack.Remove(el)
	}
	return res
}

func dfs(course int, coursePrereq map[int][]int, visited map[int]VisitedState, stack *list.List, recursive *bool) {
	// If it had already been flag as recursive, just return
	if *recursive {
		return
	}
	// If the current node is already in a Visiting state, then this graph is recursive,
	// so flag it as so and return
	if visited[course] == Visiting {
		*recursive = true
		return
	}
	// If we already visited the node, skip it
	if visited[course] == Visited {
		return
	}
	// If this node have child
	if len(coursePrereq[course]) != 0 {
		// Mark that we are visiting this node
		visited[course] = Visiting
		// Process each child before adding this to the stack
		for _, next := range coursePrereq[course] {
			dfs(next, coursePrereq, visited, stack, recursive)
		}
	}
	// Add this node to the stack and flag it as Visited
	stack.PushBack(course)
	visited[course] = Visited
}
