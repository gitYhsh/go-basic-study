package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	s := arr1[2:5]
	s1 := arr1[:]
	fmt.Println(s, s1)
	fmt.Println("golang传值是copy 传值")

	zhicopy(&arr1)

	fmt.Println(arr1)
}

func zhicopy(arr *[6]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Print(v)
	}
	arrs := make([]int, 0)
	dd := append(arrs, 10)
	fmt.Println("数组格式", dd)
}
