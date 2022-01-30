package main

func findKthLargest2(nums []int, k int) int {
	s := Solution{nums}
	ln := len(nums)
	s.quickSelect(0, ln-1, ln-k)

	return s.nums[ln-k]
}

type Solution struct {
	nums []int
}

func (s *Solution) quickSelect(left, right, k int) {

	if left == right {
		return
	}

	pivot := (left + right) / 2
	pivot = s.partition(left, right, pivot)

	if k == pivot {
		return
	}
	if k < pivot {
		s.quickSelect(left, pivot-1, k)
		return
	}
	s.quickSelect(pivot+1, right, k)
}

func (s *Solution) swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
}

func (s *Solution) partition(left, right, pivot int) int {
	pivotVal := s.nums[pivot]
	s.swap(right, pivot)

	nextPivot := left

	for i := left; i < right; i++ {
		if s.nums[i] < pivotVal {
			s.swap(i, nextPivot)
			nextPivot++
		}
	}

	s.swap(nextPivot, right)

	return nextPivot
}
