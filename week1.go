// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func bai1(x, y int) (int, int) {
	return x * y, 2*x + 2*y
}
func bai2(s string) bool {
	if len(s)%2 == 0 {
		return true
	}
	return false
}
func bai3(s []int) (int, int, int, float64, []int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	mn, mx := s[0], s[len(s)-1]
	sum := 0
	for _, v := range s {
		sum += v
	}
	var tbc float64
	tbc = float64(sum) / float64(len(s))
	return mn, mx, sum, tbc, s
}
func bai4(s []int, sum int) (int, int) {
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	for i := 0; i < len(s); i++ {
		elem, ok := m[sum-s[i]]
		if ok == true && elem != i {
			return i, elem
		}
	}
	return -1, -1
}
func main() {
	a, b := bai1(1, 2)
	fmt.Println(a, b)
	fmt.Println(bai2("hellos"))
	fmt.Println(bai3([]int{5, 4, 3, 2, 1}))
	fmt.Println(bai4([]int{4, 1, 2, 6}, 8))
}
