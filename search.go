package gcache

import "fmt"

var DebugMode bool

func Debug(args ...interface{}) {
	if !DebugMode {
		return
	}
	fmt.Println(args...)
}

func Search(n int, f func(int) bool) int {
	var depth = 0
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		depth++
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	if DebugMode {
		//fmt.Println("search depth",depth)
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func SearchFrom(to int, from int, f func(int) bool) int {
	if from > to {
		panic(fmt.Sprintf("from is bigger than to from %d, to %d", from, to))
	}
	var depth = 0
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := from, to
	for i < j {
		depth++
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	if DebugMode {
		fmt.Println("search depth", depth)
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}



//this is fastest one
func SliceInsert(s []interface{}, i int, e interface{}) []interface{} {
	s = append(s, nil /* use the zero value of the element type */)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func SliceInsert2(s []interface{}, i int, e interface{}) []interface{} {
	s = append(s[:i], append([]interface{}{e}, s[i:]...)...)
	return s
}

func SliceInsert3(s []interface{}, i int, e interface{}) []interface{} {
	tmp := append([]interface{}{e}, s[i:]...)
	s = append(s[:i], tmp...)
	return s
}

func SliceInsert4(s []interface{}, i int, e interface{}) []interface{} {
	tmp := s[:i]
	tmp = append(tmp, e)
	s = append(tmp, s[i:]...)
	return s
}
