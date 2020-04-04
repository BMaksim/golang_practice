package main

import "fmt"

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Print("\n")
}

func main() {
	arr := []int{1, 4, 3, 5, 6, 2}
	n := len(arr)
	for j := 1; j <= n-1; j++ {
		key := arr[j]
		for i := j; i > 0; i-- {
			if arr[i-1] > arr[i] {
				arr[i] = arr[i-1]
				arr[i-1] = key
			} else {
				break
			}
		}
		printArr(arr)
	}
}
