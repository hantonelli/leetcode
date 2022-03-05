package clone_graph

import (
	"fmt"
	"testing"
)

func Test_Clone_Graph(t *testing.T) {
	var n1 = Node{Val: 1}
	var n2 = Node{Val: 2}
	var n3 = Node{Val: 3}
	var n4 = Node{Val: 4}

	n1.Neighbors = []*Node{&n2, &n4}
	n2.Neighbors = []*Node{&n1, &n3}
	n3.Neighbors = []*Node{&n2, &n4}
	n4.Neighbors = []*Node{&n1, &n3}

	r := cloneGraph(&n1)
	fmt.Printf("%v", r)
}
