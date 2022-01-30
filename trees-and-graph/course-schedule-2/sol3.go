package main

func findOrder3(numCourses int, prerequisites [][]int) []int {
	courseNexts := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, v := range prerequisites {
		courseNext, course := v[0], v[1]
		courseNexts[course] = append(courseNexts[course], courseNext)
		indegree[courseNext]++
	}

	var queue []int
	for course := 0; course < numCourses; course++ {
		if indegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	var res []int
	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]

		for _, prereq := range courseNexts[course] {
			indegree[prereq]--
			if indegree[prereq] == 0 {
				queue = append(queue, prereq)
			}
		}

		res = append(res, course)
	}

	if len(res) != numCourses {
		return []int{}
	}

	return res
}
