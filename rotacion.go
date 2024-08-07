package main

import (
	"fmt"
)

func rotateRight(arr []int, m int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	m = m % n
	return append(arr[n-m:], arr[:n-m]...)
}

func printArray(arr []int) {
	fmt.Print("arr[")
	for i, v := range arr {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func main() {
	testCases := []struct {
		arr []int
		m   int
	}{
		{arr: []int{1, 2, 3, 4}, m: 1},
		{arr: []int{1, 2, 3, 4}, m: 2},
		{arr: []int{1, 2, 3, 4}, m: 3},
		{arr: []int{1, 2, 3, 4}, m: 4},
	}

	for _, tc := range testCases {
		rotated := rotateRight(tc.arr, tc.m)
		fmt.Printf("Ejemplo\tValor de m\tResultado\n")
		fmt.Printf("arr%v\tm = %d\t", tc.arr, tc.m)
		printArray(rotated)
	}
}
