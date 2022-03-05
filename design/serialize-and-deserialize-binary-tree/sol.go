package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"

	"github.com/hantonelli/leetcode/ds"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(r *ds.TreeNode) string {
	res := []int{r.Val}
	nodes := []*ds.TreeNode{r}
	for 0 < len(nodes) {
		n := nodes[0]
		nodes = nodes[1:]
		if n.Left != nil {
			res = append(res, n.Left.Val)
			nodes = append(nodes, n.Left)
		} else {
			res = append(res, ds.NULL)
		}
		if n.Right != nil {
			res = append(res, n.Right.Val)
			nodes = append(nodes, n.Right)
		} else {
			res = append(res, ds.NULL)
		}
	}
	toDelete := 0
	for i := len(res) - 1; 0 <= i; i-- {
		if res[i] == ds.NULL {
			toDelete++
		} else {
			break
		}
	}
	values := res[:len(res)-toDelete]

	valuesText := []string{}
	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	// Join our string slice.
	return strings.Join(valuesText, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *ds.TreeNode {
	if len(data) == 0 {
		return nil
	}
	// var l []int
	// if err := json.Unmarshal([]byte(data), &l); err != nil {
	// 	panic(err)
	// }
	strings := strings.Split(data, ", ")
	l := make([]int, len(strings))
	for i, s := range strings {
		l[i], _ = strconv.Atoi(s)
	}

	root := &ds.TreeNode{
		Val: l[0],
	}
	nodes := []*ds.TreeNode{root}
	i := 0
	for {
		n := nodes[0]
		nodes = nodes[1:]

		i++
		if i == len(l) {
			break
		}
		if l[i] != ds.NULL {
			lNode := ds.TreeNode{
				Val: l[i],
			}
			n.Left = &lNode
			nodes = append(nodes, &lNode)
		}

		i++
		if i == len(l) {
			break
		}
		if l[i] != ds.NULL {
			rNode := ds.TreeNode{
				Val: l[i],
			}
			n.Right = &rNode
			nodes = append(nodes, &rNode)
		}
	}

	return root
}
