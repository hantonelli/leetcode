package dataStructures

import "math"

var NULL int = math.MaxInt64

func ListToTree(l []int) *TreeNode {
	if len(l) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: l[0],
	}
	nodes := []*TreeNode{root}
	i := 0
	for {
		n := nodes[0]
		nodes = nodes[1:]

		i++
		if i == len(l) {
			break
		}
		if l[i] != NULL {
			lNode := TreeNode{
				Val: l[i],
			}
			n.Left = &lNode
			nodes = append(nodes, &lNode)
		}

		i++
		if i == len(l) {
			break
		}
		if l[i] != NULL {
			rNode := TreeNode{
				Val: l[i],
			}
			n.Right = &rNode
			nodes = append(nodes, &rNode)
		}
	}

	return root
}

func TreetoList(r *TreeNode) []int {
	res := []int{r.Val}
	nodes := []*TreeNode{r}
	for 0 < len(nodes) {
		n := nodes[0]
		nodes = nodes[1:]
		if n.Left != nil {
			res = append(res, n.Left.Val)
			nodes = append(nodes, n.Left)
		} else {
			res = append(res, NULL)
		}
		if n.Right != nil {
			res = append(res, n.Right.Val)
			nodes = append(nodes, n.Right)
		} else {
			res = append(res, NULL)
		}
	}
	toDelete := 0
	for i := len(res) - 1; 0 <= i; i-- {
		if res[i] == NULL {
			toDelete++
		} else {
			break
		}
	}
	return res[:len(res)-toDelete]
}
