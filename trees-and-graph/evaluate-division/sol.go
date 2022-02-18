package evaluateDivision

type Node struct {
	id    string
	value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	valuesMap := map[string][]Node{}
	for i := range equations {
		curr := valuesMap[equations[i][0]]
		curr = append(curr, Node{id: equations[i][1], value: values[i]})
		valuesMap[equations[i][0]] = curr

		curr2 := valuesMap[equations[i][1]]
		curr2 = append(curr2, Node{id: equations[i][0], value: 1 / values[i]})
		valuesMap[equations[i][1]] = curr2
	}

	res := make([]float64, 0, len(queries))
	for _, q := range queries {
		visited := map[string]bool{}
		res = append(res, resolve(q[0], q[1], valuesMap, visited))
	}
	return res
}

func resolve(start, to string, valuesMap map[string][]Node, visited map[string]bool) float64 {
	visited[start] = true
	for _, n := range valuesMap[start] {
		if n.id == to {
			return n.value
		} else {
			if !visited[n.id] {
				res := resolve(n.id, to, valuesMap, visited)
				if res != -1 {
					return res * n.value
				}
			}
		}
	}
	return -1
}
