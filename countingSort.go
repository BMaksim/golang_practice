package main

import "fmt"

func main() {
	arr := []int{1, 4, 3, 5, 6, 2, 6, 1, 6, 8, 12, 15, 64, 87, 1, 7, 87, 12, 98, 4, 5, 6, 1, 2}
	mass := make([]int, 100)
	for _, i := range arr {
		mass[i] += 1
	}
	out := make([]int, 0, len(arr))
	for i, num := range mass {
		for j := num; j > 0; j-- {
			out = append(out, i)
		}
	}
	fmt.Println(out)
}
