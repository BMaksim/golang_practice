package main

import "fmt"

func main() {
	arr := []int{12, 4, 3, 6, 5, 1, 7, 5}
	n := len(arr)
	var key int
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				key = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = key
			}
		}
	}
	fmt.Println(arr)
}
