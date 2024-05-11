package main

import "fmt"

// если порядок не важен, быстро
func removeElementCutLast(arr []int, i int) []int {
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return arr[:len(arr)-1]
}

// порядок важен, долго
func removeElementTwoSlices(arr []int, i int) []int {
	newArr := []int{}
	newArr = append(newArr, arr[:i]...)
	newArr = append(newArr, arr[i+1:]...)
	return newArr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}
	arr1 = removeElementCutLast(arr1, 2)
	fmt.Println(arr1)
	arr2 = removeElementTwoSlices(arr2, 1)
	fmt.Println(arr2)
}
