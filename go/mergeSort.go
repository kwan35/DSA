package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(40)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeSort(item []int) []int {
	var length = len(item)

	if length == 1 {
		return item
	}
	middle := length / 2

	return merge(mergeSort(item[:middle]), mergeSort(item[middle:]))

}

func merge(left []int, right []int) []int {
	leftlen := len(left)
	rightlen := len(right)
	total := leftlen + rightlen
	res := []int{}
	i, j := 0, 0

	for c := 0; c < total; c++ {
		if leftlen == 0 {
			res = append(res, right[j])
			j++
			continue
		}
		if rightlen == 0 {
			res = append(res, left[i])
			i++
			continue
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
			leftlen--
		} else {
			res = append(res, right[j])
			j++
			rightlen--
		}
	}
	return res
}

/* // sample
func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	fmt.Println(middle, num, "middle")
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

var ic = 0

func merge(left, right []int) (result []int) {

	fmt.Println(ic, "merge", left, right)
	ic++
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
*/
