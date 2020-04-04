package main

import "fmt"

func main() {
	arr := []int{12, 4, 3, 6, 5, 1, 7, 5}
	n := len(arr)
	var index int
	var key int
	for i := 0; i < n-1; i++ {
		key = arr[i]
		for j := i + 1; j <= n-1; j++ {
			if arr[j] < key {
				key = arr[j]
				index = j
			}
		}
		if arr[i] > arr[index] {
			arr[index] = arr[i]
			arr[i] = key
		}
	}
	fmt.Println(arr)
}
