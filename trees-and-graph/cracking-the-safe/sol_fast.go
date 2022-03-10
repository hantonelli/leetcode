package crackingTheSafe

import "strings"

type CrackList struct {
	Digit byte
	Next  *CrackList
}

type CrackStackRec struct {
	Vertex  int
	ListPtr *CrackList
}

func crackSafeFast(n int, k int) string {
	kN := 1
	for i := 0; i < n; i++ {
		kN *= k
	}
	kN1 := kN / k
	visitedE := make([]bool, kN)
	path := new(CrackList)
	stack := make([]CrackStackRec, 1)
	stack[0] = CrackStackRec{Vertex: 0, ListPtr: path}
	for len(stack) > 0 {
		rec := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ptr := rec.ListPtr
		nextV := -1
		for i := 0; i < k; i++ {
			if !visitedE[rec.Vertex*k+i] {
				nextV = rec.Vertex*k + i
				break
			}
		}
		if nextV == -1 {
			continue
		}
		for nextV != -1 {
			visitedE[nextV] = true
			tmp := new(CrackList)
			tmp.Next = ptr.Next
			tmp.Digit = byte(nextV%k) + '0'
			ptr.Next = tmp
			ptr = tmp
			rec2 := CrackStackRec{Vertex: nextV % kN1, ListPtr: ptr}
			stack = append(stack, rec2)
			nextV = -1
			for i := 0; i < k; i++ {
				if !visitedE[rec2.Vertex*k+i] {
					nextV = rec2.Vertex*k + i
					break
				}
			}
		}
		//fmt.Println(sprintfPath(path))
	}
	//return sprintfPath(path)
	var sb strings.Builder
	ptr := path.Next
	for ptr != nil {
		sb.WriteByte(ptr.Digit)
		ptr = ptr.Next
	}
	ptr = path.Next
	for i := 0; i < n-1; i++ {
		sb.WriteByte(ptr.Digit)
		ptr = ptr.Next
		if ptr == nil {
			ptr = path.Next
		}
	}
	return sb.String()
}
