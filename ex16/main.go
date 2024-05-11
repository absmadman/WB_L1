package main

import (
	"fmt"
)

// Рекурсивная функция быстрой сортировки
func recursiveSort(arr []int, left, right int) {
	if left < right {
		mean := partition(arr, left, right)
		// После разбиения можно запускать рекурсию для левой и правой части
		recursiveSort(arr, left, mean)
		recursiveSort(arr, mean+1, right)
	}
}

// Функция разбиения нужна чтобы разделить массив на 2 части, не обязательно равные
func partition(arr []int, left, right int) int {
	mean := arr[(left+right)/2]
	l, r := left, right
	for l <= r {
		// Ищем элемент больше mean в левой части
		for ; arr[l] < mean; l++ {
		}
		// Ищем элемент меньше mean в правой части
		for ; arr[r] > mean; r-- {
		}
		if l >= r {
			break
		}
		// Если нашли 2 элемента они гарантированно не на своем месте, их можно менять местами
		arr[l], arr[r] = arr[r], arr[l]
	}
	return r
}

// Сама функция быстрой сортировки
func quickSort(arr []int) {
	recursiveSort(arr, 0, len(arr)-1)
}

func main() {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	quickSort(a)
	fmt.Println(a)
}
