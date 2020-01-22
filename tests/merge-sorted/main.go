package main

import "fmt"

// zeros
// negatives
// empty lists
// one list is empty
// short list long list

func main() {
	fmt.Println(mergeLists([]int{3, 4, 6, 10, 11, 15}, []int{1, 5, 8, 12, 14, 19}))
	fmt.Println(mergeLists([]int{-4, -3, 1, 6, 10, 11, 15}, []int{-100, -1, 5, 8, 12, 14, 19}))
	fmt.Println(mergeLists([]int{}, []int{}))
	fmt.Println(mergeLists([]int{3, 4, 6, 10, 11, 15}, []int{}))
	fmt.Println(mergeLists([]int{2, 2, 4, 4}, []int{3, 3, 4}))
	fmt.Println(mergeLists([]int{15}, []int{1, 5, 8, 12, 14, 19}))
}

func mergeLists(a []int, b []int) []int {
	zipped := make([]int, 0, len(a)+len(b))

	aValue, bValue := 0, 0
	for {
		if len(a) == 0 {
			zipped = append(zipped, b...)
			return zipped
		}

		if len(b) == 0 {
			zipped = append(zipped, a...)
			return zipped
		}

		aValue = a[0]
		bValue = b[0]

		if aValue > bValue {
			zipped = append(zipped, bValue)
			b = b[1:]
		}
		if aValue <= bValue {
			zipped = append(zipped, aValue)
			a = a[1:]
		}

	}
}
